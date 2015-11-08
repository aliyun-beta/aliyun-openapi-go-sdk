Aliyun OpenAPI Go SDK

### Feature
- All Aliyun OpenAPI
- Parameter Validation
- XML and JSON
- Struct Result
- Tested
- Simple and Easy to use
- ...

### Install
```
$ cd ${GOPATH}/src # change the GOPATH if required
$ mkdir -p git.oschina.net/cxr29 && cd git.oschina.net/cxr29
$ git clone https://git.oschina.net/cxr29/aliyun-openapi-go-sdk.git
$ go install git.oschina.net/cxr29/aliyun-openapi-go-sdk
```
or
```$ go get git.oschina.net/cxr29/aliyun-openapi-go-sdk.git # it isn't working now```

### Usage
```
package main

import (
	"fmt"
	"os"

	"git.oschina.net/cxr29/aliyun-openapi-go-sdk"
	"git.oschina.net/cxr29/aliyun-openapi-go-sdk/apis/cdn/v20141111"
)

func fatal(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	api := cdn.New("YourAccessKeyId", "YourAccessKeySecret")

	_, err := api.OpenCdnService("PayByBandwidth", nil)
	if err != nil {
		e, ok := err.(openapi.Error)
		if !(ok && e.Code == "CdnService.HasOpened") {
			fatal(err)
		}
	}

	{ // ModifyCdnService
		_, err := api.ModifyCdnService("PayByTraffic", nil)
		fatal(err)
	}

	{ // DescribeCdnService
		result, err := api.DescribeCdnService(nil)
		fatal(err)
		fmt.Printf("DescribeCdnService: %+v\n", result)
	}

	{ // DescribeUserDomains
		optional := openapi.M{
			"PageSize":   int64(10),
			"PageNumber": int64(2),
			"_scheme":    "http",
		}
		result, err := api.DescribeUserDomains(optional)
		fatal(err)
		fmt.Printf("DescribeUserDomains: %+v\n", result)
	}

	// ...
}
```

### API Doc
https://godoc.org/git.oschina.net/cxr29/aliyun-openapi-go-sdk.git

### Test
```
$ export OpenAPITestAccessKeyId=YourAccessKeyId
$ export OpenAPITestAccessKeySecret=YourAccessKeySecret
$ go test -test.v ./tests
```

### Author
Chen Xianren &lt;cxr29@foxmail.com&gt;
