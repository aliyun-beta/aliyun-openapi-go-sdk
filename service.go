// Copyright 2015 Chen Xianren. All rights reserved.

package openapi

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Service represents Aliyun Open API Service,
// the AccessKeyId and the AccessKeySecret are required.
type Service struct {
	Method          string
	Unsafe          bool
	Domain          string
	AccessKeyId     string
	AccessKeySecret string
}

// NewService returns a new Service given a accessKeyId and accessKeySecret.
func NewService(accessKeyId, accessKeySecret string) Service {
	return Service{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
}

// Scheme returns http if the Unsafe is ture otherwise returns https.
func (s Service) Scheme() string {
	if s.Unsafe {
		return "http"
	}
	return "https"
}

// Do sends an HTTP request and read the HTTP response to v.
//
// See the method GetResponse and the function ReadBody to get more.
func (s Service) Do(v interface{}, args Params) error {
	res, err := s.GetResponse(args)
	if err != nil {
		return err
	}
	return ReadBody(res, v)
}

// just for test
var pause = 0 // seconds

// GetResponse sends an HTTP request and returns the HTTP response.
//
// See the method GetRequest to get more.
func (s Service) GetResponse(args Params) (*http.Response, error) {
	if pause > 0 {
		time.Sleep(time.Duration(pause) * time.Second)
	}
	req, err := s.GetRequest(args)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

// GetRequest returns a new Request given a Params.
func (s Service) GetRequest(args Params) (*http.Request, error) {
	if s.AccessKeyId == "" || s.AccessKeySecret == "" {
		return nil, errors.New("access key required")
	}

	method := args.Method
	if method == "" {
		method = s.Method
		if method == "" {
			if len(args.Body) > 0 {
				method = "POST"
			} else {
				method = "GET"
			}
		}
	}

	scheme := args.Scheme
	if scheme == "" {
		scheme = s.Scheme()
	}

	domain := GetDomain(args.Region, args.Product)
	if domain == "" {
		domain = s.Domain
		if domain == "" {
			return nil, errors.New("domain required")
		}
	}

	query := args.Query
	if query == nil {
		query = make(url.Values)
	}
	query.Set("Version", args.Version)

	header := args.Header
	if header == nil {
		header = make(http.Header)
	}
	if header.Get("User-Agent") == "" {
		header.Set("User-Agent", UserAgent)
	}
	if header.Get("Connection") == "" {
		header.Set("Connection", "keep-alive")
	}
	if header.Get("Accept-Encoding") == "" {
		header.Set("Accept-Encoding", "gzip, deflate")
	}
	if header.Get("Cache-Control") == "" {
		header.Set("Cache-Control", "no-cache")
	}
	header.Set("Date", time.Now().UTC().Format(http.TimeFormat))

	req := &http.Request{
		Method: method,
		URL: &url.URL{
			Scheme: scheme,
			Host:   domain,
			Path:   args.Expand(),
		},
		Proto:      "HTTP/1.1", // 29
		ProtoMajor: 1,
		ProtoMinor: 1,
		Host:       domain,
	}

	if len(args.Body) > 0 {
		b := []byte(args.Body.Encode())
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Content-Md5", Md5sum(b))
		req.Body = ioutil.NopCloser(bytes.NewReader(b))
	}

	switch args.Style {
	case "RPC":
		SignatureRPC(s.AccessKeyId, s.AccessKeySecret, method, query)
	case "ROA":
		SignatureROA(s.AccessKeyId, s.AccessKeySecret, method, req.URL.Path, query, header)
	default:
		return nil, errors.New("unknown style")
	}

	req.Header = header
	req.URL.RawQuery = query.Encode()

	return req, nil
}
