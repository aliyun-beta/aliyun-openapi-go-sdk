// Copyright 2015 Chen Xianren. All rights reserved.

package openapi

import (
	"os"
	"runtime"
	"testing"

	"git.oschina.net/cxr29/aliyun-openapi-go-sdk"
)

func stack(t *testing.T) {
	buf := make([]byte, 10*1024)
	t.Logf("%s", buf[:runtime.Stack(buf, false)])
}

func equal(t *testing.T, what string, expected, got interface{}) {
	if expected != got {
		stack(t)
		t.Fatal(what, "expected", expected, "but got", got)
	}
}

func fatal(t *testing.T, err error) {
	if err != nil {
		stack(t)
		t.Fatal(err)
	}
}

func newService() openapi.Service {
	return openapi.NewService(
		os.Getenv("OpenAPITestAccessKeyId"),
		os.Getenv("OpenAPITestAccessKeySecret"))
}
