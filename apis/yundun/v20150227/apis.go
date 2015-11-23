// Copyright 2015 Chen Xianren. All rights reserved.
// Code generated by openapi-meta; DO NOT EDIT

package yundun // import "github.com/cxr29/aliyun-openapi-go-sdk/apis/yundun/v20150227"

import (
	"errors"
	"fmt"

	"github.com/cxr29/aliyun-openapi-go-sdk"
)

var (
	_ = errors.New("")
	_ = fmt.Sprint("")
)

type API struct {
	openapi.Service
}

func New(accessKeyId, accessKeySecret string) API {
	return API{openapi.NewService(accessKeyId, accessKeySecret)}
}

func NewParams() openapi.Params {
	args := openapi.NewParams()
	args.Product = Product
	args.Style = Style
	args.Version = Version
	return args
}

const (
	Product = "Yundun"
	Style   = "RPC"
	Version = "2015-02-27"
)

// AllMalwareNum version 2015-02-27
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) AllMalwareNum(optional openapi.M) (*AllMalwareNumResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "AllMalwareNum")
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(AllMalwareNumResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// AllMalwareNumResponse represents the response of the api AllMalwareNum.
type AllMalwareNumResponse struct {
	AllMalwareNum int64
}

// CurrentDdosAttackNum version 2015-02-27
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) CurrentDdosAttackNum(optional openapi.M) (*CurrentDdosAttackNumResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "CurrentDdosAttackNum")
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(CurrentDdosAttackNumResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// CurrentDdosAttackNumResponse represents the response of the api CurrentDdosAttackNum.
type CurrentDdosAttackNumResponse struct {
	CurrentDdosAttackNum int64
}

// TodayAegisOnlineRate version 2015-02-27
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) TodayAegisOnlineRate(optional openapi.M) (*TodayAegisOnlineRateResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "TodayAegisOnlineRate")
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(TodayAegisOnlineRateResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// TodayAegisOnlineRateResponse represents the response of the api TodayAegisOnlineRate.
type TodayAegisOnlineRateResponse struct {
	Rate int64
}

// TodayAllkbps version 2015-02-27
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) TodayAllkbps(optional openapi.M) (*TodayAllkbpsResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "TodayAllkbps")
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(TodayAllkbpsResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// TodayAllkbpsResponse represents the response of the api TodayAllkbps.
type TodayAllkbpsResponse struct {
	Kbps int64
}

// TodayAllpps version 2015-02-27
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) TodayAllpps(optional openapi.M) (*TodayAllppsResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "TodayAllpps")
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(TodayAllppsResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// TodayAllppsResponse represents the response of the api TodayAllpps.
type TodayAllppsResponse struct {
	Pps int64
}

// TodayBackdoor version 2015-02-27
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) TodayBackdoor(optional openapi.M) (*TodayBackdoorResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "TodayBackdoor")
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(TodayBackdoorResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// TodayBackdoorResponse represents the response of the api TodayBackdoor.
type TodayBackdoorResponse struct {
	Backdoor int64
}

// TodayCrackIntercept version 2015-02-27
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) TodayCrackIntercept(optional openapi.M) (*TodayCrackInterceptResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "TodayCrackIntercept")
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(TodayCrackInterceptResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// TodayCrackInterceptResponse represents the response of the api TodayCrackIntercept.
type TodayCrackInterceptResponse struct {
	InterceptNum int64
}

// TodayMalwareNum version 2015-02-27
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) TodayMalwareNum(optional openapi.M) (*TodayMalwareNumResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "TodayMalwareNum")
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(TodayMalwareNumResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// TodayMalwareNumResponse represents the response of the api TodayMalwareNum.
type TodayMalwareNumResponse struct {
	TodayMalwareNum int64
}

// TodayqpsByRegion version 2015-02-27
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) TodayqpsByRegion(optional openapi.M) (*TodayqpsByRegionResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "TodayqpsByRegion")
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(TodayqpsByRegionResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// TodayqpsByRegionResponse represents the response of the api TodayqpsByRegion.
type TodayqpsByRegionResponse struct {
	Data struct {
		Region []struct {
			RegionFlow   int64
			RegionId     string
			RegionNumber int64
		}
	}
}

// WebAttackNum version 2015-02-27
//
// optional parameters:
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
func (api API) WebAttackNum(optional openapi.M) (*WebAttackNumResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "WebAttackNum")
	if v, ok := optional["_method"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "GET|POST") {
				return nil, errors.New("_method must be GET|POST")
			}
			args.Method = s
		} else {
			return nil, errors.New("_method must be string")
		}
	}
	if v, ok := optional["_region"]; ok {
		if s, ok := v.(string); ok {
			args.Region = s
		} else {
			return nil, errors.New("_region must be string")
		}
	}
	if v, ok := optional["_scheme"]; ok {
		if s, ok := v.(string); ok {
			if !openapi.IsIn(s, "http|https") {
				return nil, errors.New("_scheme must be http|https")
			}
			args.Scheme = s
		} else {
			return nil, errors.New("_scheme must be string")
		}
	}

	result := new(WebAttackNumResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// WebAttackNumResponse represents the response of the api WebAttackNum.
type WebAttackNumResponse struct {
	WebAttackNum int64
}
