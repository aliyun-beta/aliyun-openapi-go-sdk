// Copyright 2015 Chen Xianren. All rights reserved.
// Code generated from endpoints.xml; DO NOT EDIT

package openapi

// Region List
const (
	RegionCNBeijing           = "cn-beijing"
	RegionCNHangzhou          = "cn-hangzhou"
	RegionCNHangzhouD         = "cn-hangzhou-d"
	RegionCNHongkong          = "cn-hongkong"
	RegionCNShanghaiEt2B01    = "cn-shanghai-et2-b01"
	RegionCNShanghai          = "cn-shanghai"
	RegionUSWest1             = "us-west-1"
	RegionCNShanghaiEt2Test01 = "cn-shanghai-et2-test01"
	RegionAPSoutheast1        = "ap-southeast-1"
)

// Domains [region][product]
var Domains = map[string]map[string]string{
	"cn-hangzhou": {
		"Aas":         "aas.aliyuncs.com",
		"Ace":         "ace.cn-hangzhou.aliyuncs.com",
		"Alert":       "alert.aliyuncs.com",
		"Bss":         "bss.aliyuncs.com",
		"Cdn":         "cdn.aliyuncs.com",
		"Cms":         "metrics.aliyuncs.com",
		"Crm":         "crm-cn-hangzhou.aliyuncs.com",
		"Dqs":         "dqs.aliyuncs.com",
		"Drc":         "drc.aliyuncs.com",
		"Drds":        "drds.aliyuncs.com",
		"Dts":         "dts.aliyuncs.com",
		"Ecs":         "ecs-cn-hangzhou.aliyuncs.com",
		"Emr":         "emr.aliyuncs.com",
		"Location":    "location.aliyuncs.com",
		"M-kvstore":   "m-kvstore.aliyuncs.com",
		"Ocs":         "pop-ocs.aliyuncs.com",
		"Oms":         "oms.aliyuncs.com",
		"Ons":         "ons.aliyuncs.com",
		"Oss":         "oss-cn-hangzhou.aliyuncs.com",
		"OssAdmin":    "oss-admin.aliyuncs.com",
		"Ots":         "ots-pop.aliyuncs.com",
		"PTS":         "pts.aliyuncs.com",
		"Push":        "cloudpush.aliyuncs.com",
		"R-kvstore":   "r-kvstore-cn-hangzhou.aliyuncs.com",
		"ROS":         "ros.aliyuncs.com",
		"Ram":         "ram.aliyuncs.com",
		"Rds":         "rds.aliyuncs.com",
		"Risk":        "risk-cn-hangzhou.aliyuncs.com",
		"Slb":         "slb-pop.aliyuncs.com",
		"Sts":         "sts.aliyuncs.com",
		"Ubsms":       "ubsms.aliyuncs.com",
		"Ubsms-inner": "ubsms-inner.aliyuncs.com",
		"Yundun":      "yundun-cn-hangzhou.aliyuncs.com",
	},
	"cn-qingdao": {
		"Aas":          "aas.aliyuncs.com",
		"Ace":          "ace.cn-hangzhou.aliyuncs.com",
		"Alert":        "alert.aliyuncs.com",
		"BatchCompute": "batchcompute.cn-qingdao.aliyuncs.com",
		"Bss":          "bss.aliyuncs.com",
		"Cdn":          "cdn.aliyuncs.com",
		"Cms":          "metrics.aliyuncs.com",
		"Crm":          "crm-cn-hangzhou.aliyuncs.com",
		"Dqs":          "dqs.aliyuncs.com",
		"Drc":          "drc.aliyuncs.com",
		"Drds":         "drds.aliyuncs.com",
		"Dts":          "dts.aliyuncs.com",
		"Ecs":          "ecs-cn-hangzhou.aliyuncs.com",
		"Emr":          "emr.aliyuncs.com",
		"Location":     "location.aliyuncs.com",
		"M-kvstore":    "m-kvstore.aliyuncs.com",
		"Ocs":          "pop-ocs.aliyuncs.com",
		"Oms":          "oms.aliyuncs.com",
		"Ons":          "ons.aliyuncs.com",
		"Oss":          "oss-cn-hangzhou.aliyuncs.com",
		"OssAdmin":     "oss-admin.aliyuncs.com",
		"Ots":          "ots-pop.aliyuncs.com",
		"PTS":          "pts.aliyuncs.com",
		"Push":         "cloudpush.aliyuncs.com",
		"R-kvstore":    "r-kvstore-cn-hangzhou.aliyuncs.com",
		"ROS":          "ros.aliyuncs.com",
		"Ram":          "ram.aliyuncs.com",
		"Rds":          "rds.aliyuncs.com",
		"Risk":         "risk-cn-hangzhou.aliyuncs.com",
		"Slb":          "slb-pop.aliyuncs.com",
		"Sts":          "sts.aliyuncs.com",
		"Ubsms":        "ubsms.aliyuncs.com",
		"Ubsms-inner":  "ubsms-inner.aliyuncs.com",
		"Yundun":       "yundun-cn-hangzhou.aliyuncs.com",
	},
	"cn-shenzhen": {
		"Aas":          "aas.aliyuncs.com",
		"Ace":          "ace.cn-hangzhou.aliyuncs.com",
		"Alert":        "alert.aliyuncs.com",
		"BatchCompute": "batchcompute.cn-shenzhen.aliyuncs.com",
		"Bss":          "bss.aliyuncs.com",
		"Cdn":          "cdn.aliyuncs.com",
		"Cms":          "metrics.aliyuncs.com",
		"Crm":          "crm-cn-hangzhou.aliyuncs.com",
		"Dqs":          "dqs.aliyuncs.com",
		"Drc":          "drc.aliyuncs.com",
		"Drds":         "drds.aliyuncs.com",
		"Dts":          "dts.aliyuncs.com",
		"Ecs":          "ecs-cn-hangzhou.aliyuncs.com",
		"Emr":          "emr.aliyuncs.com",
		"Location":     "location.aliyuncs.com",
		"M-kvstore":    "m-kvstore.aliyuncs.com",
		"Ocs":          "pop-ocs.aliyuncs.com",
		"Oms":          "oms.aliyuncs.com",
		"Ons":          "ons.aliyuncs.com",
		"Oss":          "oss-cn-hangzhou.aliyuncs.com",
		"OssAdmin":     "oss-admin.aliyuncs.com",
		"Ots":          "ots-pop.aliyuncs.com",
		"PTS":          "pts.aliyuncs.com",
		"Push":         "cloudpush.aliyuncs.com",
		"R-kvstore":    "r-kvstore-cn-hangzhou.aliyuncs.com",
		"ROS":          "ros.aliyuncs.com",
		"Ram":          "ram.aliyuncs.com",
		"Rds":          "rds.aliyuncs.com",
		"Risk":         "risk-cn-hangzhou.aliyuncs.com",
		"Slb":          "slb-pop.aliyuncs.com",
		"Sts":          "sts.aliyuncs.com",
		"Ubsms":        "ubsms.aliyuncs.com",
		"Ubsms-inner":  "ubsms-inner.aliyuncs.com",
		"Yundun":       "yundun-cn-hangzhou.aliyuncs.com",
	},
}

// GetDomain returns the access domain given a region and product.
func GetDomain(region, product string) string {
	if region == "" {
		region = "cn-hangzhou"
	}
	if v, ok := Domains[region]; ok {
		return v[product]
	}
	return ""
}
