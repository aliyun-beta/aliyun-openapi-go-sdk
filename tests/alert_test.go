// Copyright 2015 Chen Xianren. All rights reserved.

package openapi

import (
	"testing"

	"github.com/cxr29/aliyun-openapi-go-sdk"
	"github.com/cxr29/aliyun-openapi-go-sdk/apis/alert/v20150815"
)

func TestAlert(t *testing.T) {
	api := alert.API{newService()}

	{ // CreateAlert
		_, err := api.CreateAlert("aliyun-openapi-go-sdk", openapi.M{
			"_region": "cn-qingdao",
		})
		e, ok := err.(openapi.Error)
		if !(ok && e.Code == "ServiceUnavailable") {
			t.Fatal("expected ServiceUnavailable")
		}
	}
}
