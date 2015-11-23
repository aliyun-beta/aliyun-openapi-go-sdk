// Copyright 2015 Chen Xianren. All rights reserved.

package openapi

import (
	"testing"

	"github.com/cxr29/aliyun-openapi-go-sdk"
	"github.com/cxr29/aliyun-openapi-go-sdk/apis/cdn/v20141111"
)

func TestCdn(t *testing.T) {
	api := cdn.API{newService()}

	{ // OpenCdnService
		result, err := api.OpenCdnService("PayByBandwidth", nil)
		e, ok := err.(openapi.Error)
		if !(ok && e.Code == "CdnService.HasOpened") {
			fatal(t, err)
			equal(t, "Response", "OpenCdnServiceResponse", result.XMLName.Local)
		}
	}

	{ // ModifyCdnService
		result, err := api.ModifyCdnService("PayByTraffic", nil)
		fatal(t, err)
		equal(t, "Response", "ModifyCdnServiceResponse", result.XMLName.Local)
	}

	{ // DescribeCdnService
		result, err := api.DescribeCdnService(nil)
		fatal(t, err)
		equal(t, "InternetChargeType", "PayByTraffic", result.InternetChargeType)
	}

	{ // DescribeUserDomains
		result, err := api.DescribeUserDomains(openapi.M{
			"PageSize":   int64(10),
			"PageNumber": int64(2),
		})
		fatal(t, err)
		equal(t, "PageSize", int64(10), result.PageSize)
		equal(t, "PageNumber", int64(2), result.PageNumber)
	}

	{ // DescribeCdnMonitorData
		_, err := api.DescribeCdnMonitorData("InvalidDomain.NotFound", openapi.M{
			"_scheme": "http",
		})
		e, ok := err.(openapi.Error)
		if !(ok && e.Code == "InvalidDomain.NotFound") {
			t.Fatal("expected InvalidDomain.NotFound")
		}
	}
}
