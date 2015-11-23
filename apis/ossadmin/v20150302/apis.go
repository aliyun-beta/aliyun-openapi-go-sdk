// Copyright 2015 Chen Xianren. All rights reserved.
// Code generated by openapi-meta; DO NOT EDIT

package ossadmin // import "github.com/cxr29/aliyun-openapi-go-sdk/apis/ossadmin/v20150302"

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
	Product = "OssAdmin"
	Style   = "RPC"
	Version = "2015-03-02"
)

// CreateOssInstance version 2015-03-02
//
// optional parameters:
//  name: OwnerAccount, type: string
//  name: OwnerId, type: int64
//  name: ResourceOwnerAccount, type: string
//  name: ResourceOwnerId, type: int64
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
//  name: region, type: string
func (api API) CreateOssInstance(optional openapi.M) (*CreateOssInstanceResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "CreateOssInstance")
	if v, ok := optional["OwnerAccount"]; ok {
		if OwnerAccount, ok := v.(string); ok {
			args.Query.Set("OwnerAccount", OwnerAccount)
		} else {
			return nil, errors.New("OwnerAccount must be string")
		}
	}
	if v, ok := optional["OwnerId"]; ok {
		if OwnerId, ok := v.(int64); ok {
			args.Query.Set("OwnerId", fmt.Sprint(OwnerId))
		} else {
			return nil, errors.New("OwnerId must be int64")
		}
	}
	if v, ok := optional["ResourceOwnerAccount"]; ok {
		if ResourceOwnerAccount, ok := v.(string); ok {
			args.Query.Set("ResourceOwnerAccount", ResourceOwnerAccount)
		} else {
			return nil, errors.New("ResourceOwnerAccount must be string")
		}
	}
	if v, ok := optional["ResourceOwnerId"]; ok {
		if ResourceOwnerId, ok := v.(int64); ok {
			args.Query.Set("ResourceOwnerId", fmt.Sprint(ResourceOwnerId))
		} else {
			return nil, errors.New("ResourceOwnerId must be int64")
		}
	}
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
	if v, ok := optional["region"]; ok {
		if region, ok := v.(string); ok {
			args.Query.Set("region", region)
		} else {
			return nil, errors.New("region must be string")
		}
	}

	result := new(CreateOssInstanceResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateOssInstanceResponse represents the response of the api CreateOssInstance.
type CreateOssInstanceResponse struct {
	AliUid         int64
	EndTime        string
	InstacneStatus string
	InstanceId     string
	InstanceName   string
	StartTime      string
}

// ReleaseOssInstance version 2015-03-02
//
// optional parameters:
//  name: OwnerAccount, type: string
//  name: OwnerId, type: int64
//  name: ResourceOwnerAccount, type: string
//  name: ResourceOwnerId, type: int64
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
//  name: region, type: string
func (api API) ReleaseOssInstance(optional openapi.M) (*openapi.Response, error) {
	args := NewParams()

	args.Query.Set("Action", "ReleaseOssInstance")
	if v, ok := optional["OwnerAccount"]; ok {
		if OwnerAccount, ok := v.(string); ok {
			args.Query.Set("OwnerAccount", OwnerAccount)
		} else {
			return nil, errors.New("OwnerAccount must be string")
		}
	}
	if v, ok := optional["OwnerId"]; ok {
		if OwnerId, ok := v.(int64); ok {
			args.Query.Set("OwnerId", fmt.Sprint(OwnerId))
		} else {
			return nil, errors.New("OwnerId must be int64")
		}
	}
	if v, ok := optional["ResourceOwnerAccount"]; ok {
		if ResourceOwnerAccount, ok := v.(string); ok {
			args.Query.Set("ResourceOwnerAccount", ResourceOwnerAccount)
		} else {
			return nil, errors.New("ResourceOwnerAccount must be string")
		}
	}
	if v, ok := optional["ResourceOwnerId"]; ok {
		if ResourceOwnerId, ok := v.(int64); ok {
			args.Query.Set("ResourceOwnerId", fmt.Sprint(ResourceOwnerId))
		} else {
			return nil, errors.New("ResourceOwnerId must be int64")
		}
	}
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
	if v, ok := optional["region"]; ok {
		if region, ok := v.(string); ok {
			args.Query.Set("region", region)
		} else {
			return nil, errors.New("region must be string")
		}
	}

	result := new(openapi.Response)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// RestartOssInstance version 2015-03-02
//
// optional parameters:
//  name: OwnerAccount, type: string
//  name: OwnerId, type: int64
//  name: ResourceOwnerAccount, type: string
//  name: ResourceOwnerId, type: int64
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
//  name: region, type: string
func (api API) RestartOssInstance(optional openapi.M) (*openapi.Response, error) {
	args := NewParams()

	args.Query.Set("Action", "RestartOssInstance")
	if v, ok := optional["OwnerAccount"]; ok {
		if OwnerAccount, ok := v.(string); ok {
			args.Query.Set("OwnerAccount", OwnerAccount)
		} else {
			return nil, errors.New("OwnerAccount must be string")
		}
	}
	if v, ok := optional["OwnerId"]; ok {
		if OwnerId, ok := v.(int64); ok {
			args.Query.Set("OwnerId", fmt.Sprint(OwnerId))
		} else {
			return nil, errors.New("OwnerId must be int64")
		}
	}
	if v, ok := optional["ResourceOwnerAccount"]; ok {
		if ResourceOwnerAccount, ok := v.(string); ok {
			args.Query.Set("ResourceOwnerAccount", ResourceOwnerAccount)
		} else {
			return nil, errors.New("ResourceOwnerAccount must be string")
		}
	}
	if v, ok := optional["ResourceOwnerId"]; ok {
		if ResourceOwnerId, ok := v.(int64); ok {
			args.Query.Set("ResourceOwnerId", fmt.Sprint(ResourceOwnerId))
		} else {
			return nil, errors.New("ResourceOwnerId must be int64")
		}
	}
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
	if v, ok := optional["region"]; ok {
		if region, ok := v.(string); ok {
			args.Query.Set("region", region)
		} else {
			return nil, errors.New("region must be string")
		}
	}

	result := new(openapi.Response)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// StopOssInstance version 2015-03-02
//
// optional parameters:
//  name: OwnerAccount, type: string
//  name: OwnerId, type: int64
//  name: ResourceOwnerAccount, type: string
//  name: ResourceOwnerId, type: int64
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
//  name: _scheme, type: string, optional values: http|https
//  name: region, type: string
func (api API) StopOssInstance(optional openapi.M) (*openapi.Response, error) {
	args := NewParams()

	args.Query.Set("Action", "StopOssInstance")
	if v, ok := optional["OwnerAccount"]; ok {
		if OwnerAccount, ok := v.(string); ok {
			args.Query.Set("OwnerAccount", OwnerAccount)
		} else {
			return nil, errors.New("OwnerAccount must be string")
		}
	}
	if v, ok := optional["OwnerId"]; ok {
		if OwnerId, ok := v.(int64); ok {
			args.Query.Set("OwnerId", fmt.Sprint(OwnerId))
		} else {
			return nil, errors.New("OwnerId must be int64")
		}
	}
	if v, ok := optional["ResourceOwnerAccount"]; ok {
		if ResourceOwnerAccount, ok := v.(string); ok {
			args.Query.Set("ResourceOwnerAccount", ResourceOwnerAccount)
		} else {
			return nil, errors.New("ResourceOwnerAccount must be string")
		}
	}
	if v, ok := optional["ResourceOwnerId"]; ok {
		if ResourceOwnerId, ok := v.(int64); ok {
			args.Query.Set("ResourceOwnerId", fmt.Sprint(ResourceOwnerId))
		} else {
			return nil, errors.New("ResourceOwnerId must be int64")
		}
	}
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
	if v, ok := optional["region"]; ok {
		if region, ok := v.(string); ok {
			args.Query.Set("region", region)
		} else {
			return nil, errors.New("region must be string")
		}
	}

	result := new(openapi.Response)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}
