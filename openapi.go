// Copyright 2015 Chen Xianren. All rights reserved.

// Package openapi implements a library of Aliyun OpenAPI.
package openapi

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

// UserAgent is the default user agent and is used by GetRequest.
var UserAgent = "aliyun-openapi-go-sdk"

// ReadBody reads the response body and stores the result in the value pointed to by v,
//
// When the status code is 4xx or 5xx return the Error.
//
// If v is nil, the response body is discarded or, if v's type is not
// *[]byte, *os.File and *bytes.Buffer,
// then decode XML or JSON to it.
//
// Content-Encoding deflate and gzip are supported.
func ReadBody(res *http.Response, v interface{}) error {
	defer res.Body.Close()

	var rc io.ReadCloser
	var err error

	switch res.Header.Get("Content-Encoding") {
	case "deflate":
		rc = flate.NewReader(res.Body)
		defer rc.Close()
	case "gzip":
		rc, err = gzip.NewReader(res.Body)
		if err != nil {
			return err
		}
		defer rc.Close()
	default:
		rc = res.Body
	}

	isJSON := strings.Contains(strings.ToUpper(res.Header.Get("Content-Type")), "JSON") ||
		res.Request.URL.Query().Get("Format") == "JSON"

	switch res.StatusCode / 100 {
	case 4, 5: // 7
		b, err := ioutil.ReadAll(rc)
		if err != nil {
			return err
		}
		var e Error
		if len(b) > 0 { // no content
			if isJSON {
				err = json.Unmarshal(b, &e)
			} else {
				err = xml.Unmarshal(b, &e)
			}
			if err != nil {
				return err
			}
		} else {
			e.Code = res.Status
		}
		return e
	}

	switch i := v.(type) {
	case nil:
	case *[]byte:
		*i, err = ioutil.ReadAll(rc)
	case *os.File:
		_, err = io.Copy(i, rc)
	case *bytes.Buffer:
		_, err = i.ReadFrom(rc)
	default:
		if isJSON {
			err = json.NewDecoder(rc).Decode(v)
		} else {
			err = xml.NewDecoder(rc).Decode(v)
		}
	}

	return err
}

// Error represents the error response when the status code is 4xx or 5xx.
//
// Relevant documentation:
//
// https://docs.aliyun.com/#/pub/ecs/open-api/requestmethod&commonresponse
type Error struct {
	Code      string
	Message   string
	RequestId string
	HostId    string
}

func (e Error) Error() string {
	return fmt.Sprintf("Code: %s, RequestId: %s, HostId: %s, Message: %s", e.Code, e.RequestId, e.HostId, e.Message)
}

// Response represents a response only returns RequestId.
type Response struct {
	XMLName   xml.Name `json:"-"`
	RequestId string
}

// M represents optional parameters.
type M map[string]interface{}

// A Params represents the request parameters.
type Params struct {
	Product     string
	Style       string
	Version     string
	Method      string
	Scheme      string
	Region      string
	Pattern     string
	Path        map[string]string
	Query, Body url.Values
	Header      http.Header
}

// NewParams returns a initialized Params.
func NewParams() Params {
	return Params{
		Path:   make(map[string]string),
		Query:  make(url.Values),
		Body:   make(url.Values),
		Header: make(http.Header),
	}
}

var squareBracketRegexp = regexp.MustCompile(`(\[.*?\])`)

// Expand returns the string of pattern with all the text between square brackets
// are replaced by the corresponding path value.
func (args Params) Expand() string {
	return squareBracketRegexp.ReplaceAllStringFunc(args.Pattern, func(s string) string {
		return args.Path[s[1:len(s)-1]]
	})
}
