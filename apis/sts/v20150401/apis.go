// Copyright 2015 Chen Xianren. All rights reserved.
// Code generated by openapi-meta; DO NOT EDIT

package sts

import (
	"errors"
	"fmt"

	"git.oschina.net/cxr29/aliyun-openapi-go-sdk"
)

var (
	_ = errors.New("")
	_ = fmt.Sprint("")
	_ = make(openapi.M)
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
	Product = "Sts"
	Style   = "RPC"
	Version = "2015-04-01"
)

// AssumeRole version 2015-04-01
//
// required parameters:
//  name: RoleArn, type: string
//  name: RoleSessionName, type: string
//
// optional parameters:
//  name: DurationSeconds, type: int64
//  name: Policy, type: string
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
func (api API) AssumeRole(RoleArn, RoleSessionName string, optional openapi.M) (*AssumeRoleResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "AssumeRole")
	args.Query.Set("RoleArn", RoleArn)
	args.Query.Set("RoleSessionName", RoleSessionName)
	args.Scheme = "https"
	if v, ok := optional["DurationSeconds"]; ok {
		if DurationSeconds, ok := v.(int64); ok {
			args.Query.Set("DurationSeconds", fmt.Sprint(DurationSeconds))
		} else {
			return nil, errors.New("DurationSeconds must be int64")
		}
	}
	if v, ok := optional["Policy"]; ok {
		if Policy, ok := v.(string); ok {
			args.Query.Set("Policy", Policy)
		} else {
			return nil, errors.New("Policy must be string")
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

	result := new(AssumeRoleResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// AssumeRoleResponse represents the response of the api AssumeRole.
type AssumeRoleResponse struct {
	AssumedRoleUser struct {
		Arn           string
		AssumedRoleId string
	}
	Credentials struct {
		AccessKeyId     string
		AccessKeySecret string
		Expiration      string
		SecurityToken   string
	}
}

// AssumeRoleWithServiceIdentity version 2015-04-01
//
// required parameters:
//  name: AssumeRoleFor, type: string
//  name: RoleArn, type: string
//  name: RoleSessionName, type: string
//
// optional parameters:
//  name: DurationSeconds, type: int64
//  name: Policy, type: string
//  name: _method, type: string, optional values: GET|POST
//  name: _region, type: string
func (api API) AssumeRoleWithServiceIdentity(AssumeRoleFor, RoleArn, RoleSessionName string, optional openapi.M) (*AssumeRoleWithServiceIdentityResponse, error) {
	args := NewParams()

	args.Query.Set("Action", "AssumeRoleWithServiceIdentity")
	args.Query.Set("AssumeRoleFor", AssumeRoleFor)
	args.Query.Set("RoleArn", RoleArn)
	args.Query.Set("RoleSessionName", RoleSessionName)
	args.Scheme = "https"
	if v, ok := optional["DurationSeconds"]; ok {
		if DurationSeconds, ok := v.(int64); ok {
			args.Query.Set("DurationSeconds", fmt.Sprint(DurationSeconds))
		} else {
			return nil, errors.New("DurationSeconds must be int64")
		}
	}
	if v, ok := optional["Policy"]; ok {
		if Policy, ok := v.(string); ok {
			args.Query.Set("Policy", Policy)
		} else {
			return nil, errors.New("Policy must be string")
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

	result := new(AssumeRoleWithServiceIdentityResponse)
	if err := api.Service.Do(result, args); err != nil {
		return nil, err
	}
	return result, nil
}

// AssumeRoleWithServiceIdentityResponse represents the response of the api AssumeRoleWithServiceIdentity.
type AssumeRoleWithServiceIdentityResponse struct {
	AssumedRoleUser struct {
		Arn           string
		AssumedRoleId string
	}
	Credentials struct {
		AccessKeyId     string
		AccessKeySecret string
		Expiration      string
		SecurityToken   string
	}
}
