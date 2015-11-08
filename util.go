// Copyright 2015 Chen Xianren. All rights reserved.

package openapi

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"hash"
	"strings"
)

// IsIn returns true if s in the vertical bar separated v.
func IsIn(s, v string) bool {
	for _, i := range strings.Split(v, "|") {
		if i == s {
			return true
		}
	}
	return false
}

// Md5sum returns the base64 MD5 checksum of the data.
func Md5sum(data []byte) string {
	a := md5.Sum(data)
	return base64.StdEncoding.EncodeToString(a[:])
}

// HmacSha1 returns the base64 HMAC-SHA1 hash of the data using the given secret.
func HmacSha1(secret, data string) string {
	return HmacX(secret, data, nil)
}

// HmacSha256 returns the base64 HMAC-SHA256 hash of the data using the given secret.
func HmacSha256(secret, data string) string {
	return HmacX(secret, data, sha256.New)
}

// HmacX returns the base64 HMAC-X hash of the data using the given secret.
//
// The SHA1 is used if the given hash.Hash type is nil.
func HmacX(secret, data string, h func() hash.Hash) string {
	if h == nil {
		h = sha1.New
	}
	x := hmac.New(h, []byte(secret))
	x.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(x.Sum(nil))
}
