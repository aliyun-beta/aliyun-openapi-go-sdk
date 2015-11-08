// Copyright 2015 Chen Xianren. All rights reserved.

package openapi

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
)

var r1, r2 *rand.Rand

func init() {
	t := time.Now()
	r1 = rand.New(rand.NewSource(t.UnixNano()))
	r2 = rand.New(rand.NewSource(int64(os.Getppid() + os.Getpid() + t.Nanosecond())))
}

func format(u uint32) string {
	return fmt.Sprintf("%08x", u)
}

func random() string {
	t := time.Now()
	return format(r1.Uint32()) + format(r2.Uint32()) + format(uint32(t.Unix())) + format(uint32(t.Nanosecond()))
}

type dict [][2]string

func (a dict) Len() int {
	return len(a)
}
func (a dict) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}
func (a dict) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a dict) Sort() {
	sort.Sort(a)
}

// SignatureRPC signature the RPC api style and set the Signature query.
func SignatureRPC(accessKeyId, accessKeySecret, method string, query url.Values) {
	query.Set("AccessKeyId", accessKeyId)
	sm := query.Get("SignatureMethod")
	if sm != "HMAC-SHA256" {
		sm = "HMAC-SHA1"
	}
	query.Set("SignatureMethod", sm)
	query.Set("SignatureVersion", "1.0")
	query.Set("SignatureNonce", random())
	query.Set("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	delete(query, "Signature")

	s := method + "&%2F&" + url.QueryEscape(CanonicalizedQuery(query, true))
	if sm == "HMAC-SHA256" {
		s = HmacSha256(accessKeySecret+"&", s)
	} else {
		s = HmacSha1(accessKeySecret+"&", s)
	}
	query.Set("Signature", s)
}

// CanonicalizedHeader canonicalized the prefix headers.
func CanonicalizedHeader(header http.Header, prefix string) string {
	if len(header) > 0 {
		var a dict
		for k, v := range header {
			if l := strings.ToLower(k); strings.HasPrefix(l, prefix) {
				x := ""
				if len(v) > 0 {
					x = v[0]
				}
				a = append(a, [2]string{l, x})
			}
		}
		if n := len(a); n > 0 {
			a.Sort()
			b := make([]string, n)
			for k, v := range a {
				b[k] = v[0] + ":" + v[1]
			}
			return strings.Join(b, "\n")
		}
	}
	return ""
}

// CanonicalizedQuery canonicalized the query,
// if the escape is true the key and value are escaped.
func CanonicalizedQuery(q url.Values, escape bool) string {
	var a dict
	for k, v := range q {
		x := ""
		if len(v) > 0 {
			x = v[0]
		}
		a = append(a, [2]string{k, x})
	}
	a.Sort()
	b := make([]string, len(a))
	for k, v := range a {
		if escape {
			b[k] = url.QueryEscape(v[0]) + "=" + url.QueryEscape(v[1])
		} else {
			b[k] = v[0]
			if v[1] != "" {
				b[k] += "=" + v[1]
			}
		}
	}
	return strings.Join(b, "&")
}

// SignatureROA signature the ROA api style and set the Authorization header.
func SignatureROA(accessKeyId, accessKeySecret string, method string, path string, query url.Values, header http.Header) {
	a := []string{method}

	if header.Get("Date") == "" {
		header.Set("Date", time.Now().UTC().Format(http.TimeFormat))
	}

	for _, k := range []string{"Accept", "Content-Md5", "Content-Type", "Date"} {
		a = append(a, header.Get(k))
	}

	sm := header.Get("x-acs-signature-method")
	if sm != "HMAC-SHA256" {
		sm = "HMAC-SHA1"
	}
	header.Set("x-acs-signature-method", sm)
	header.Set("x-acs-signature-version", "1.0")

	cq := CanonicalizedQuery(query, false)
	if cq == "" {
		cq = path
	} else {
		cq = path + "?" + cq
	}
	a = append(a, CanonicalizedHeader(header, "x-acs-"), cq)

	s := strings.Join(a, "\n")
	if sm == "HMAC-SHA256" {
		s = HmacSha256(accessKeySecret, s)
	} else {
		s = HmacSha1(accessKeySecret, s)
	}
	header.Set("Authorization", "acs "+accessKeyId+":"+s)
}
