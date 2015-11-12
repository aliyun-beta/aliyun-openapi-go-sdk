package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Endpoint struct {
	Name      string `xml:"name,attr"`
	RegionIds struct {
		RegionId []string
	}
	Products struct {
		Product []struct {
			ProductName string
			DomainName  string
		}
	}
}

type Endpoints struct {
	Endpoint []Endpoint
}

func main() {
	data, err := ioutil.ReadFile("openapi-meta/endpoints.xml")
	ErrFatal(err)
	var v Endpoints
	err = xml.Unmarshal(data, &v)
	ErrFatal(err)

	buf := new(bytes.Buffer)
	buf.WriteString("// Copyright 2015 Chen Xianren. All rights reserved.\n")
	buf.WriteString("// Code generated from endpoints.xml; DO NOT EDIT\n\n")
	buf.WriteString("package openapi\n")

	buf.WriteString("\n// Region List\n")
	buf.WriteString("const (\n")
	{
		var a []string
		for _, i := range v.Endpoint {
			for _, j := range i.RegionIds.RegionId {
				j = strings.TrimSpace(j)
				a = append(a, fmt.Sprintf("%s = %q\n", constName(j), j))
			}
		}
		buf.WriteString(strings.Join(RemoveDuplicate(a), ""))
	}
	buf.WriteString(")\n")

	buf.WriteString("\n// Regions [Region]:Endpoint\n")
	buf.WriteString("var Regions = map[string]string{\n")
	{
		var a []string
		for _, i := range v.Endpoint {
			i.Name = strings.TrimSpace(i.Name)
			for _, j := range i.RegionIds.RegionId {
				j = strings.TrimSpace(j)
				a = append(a, fmt.Sprintf("%s: %s,\n", constName(j), constName(i.Name)))
			}
		}
		buf.WriteString(strings.Join(RemoveDuplicate(a), ""))
	}
	buf.WriteString("}\n")

	buf.WriteString("\n// Domains [Endpoint][Product]:Domain\n")
	buf.WriteString("var Domains =  map[string]map[string]string{\n")
	{
		var a []string
		for _, i := range v.Endpoint {
			i.Name = strings.TrimSpace(i.Name)
			var b []string
			for _, j := range i.Products.Product {
				j.ProductName = strings.TrimSpace(j.ProductName)
				j.DomainName = strings.TrimSpace(j.DomainName)
				b = append(b, fmt.Sprintf("%q: %q,\n", j.ProductName, j.DomainName))
			}
			a = append(a, fmt.Sprintf("%s: {\n%s},\n",
				constName(i.Name), strings.Join(RemoveDuplicate(b), "")))
		}
		sort.Strings(a)
		buf.WriteString(strings.Join(a, ""))
	}
	buf.WriteString("}\n")

	buf.WriteString(`
// GetDomain returns the access domain given a product and optional region.
func GetDomain(product, region string) string {
	if product == "" {
		return ""
	}
	if region == "" {
		region = RegionCNHangzhou
	}
	return Domains[Regions[region]][product]
}
`)
	ErrFatal(ioutil.WriteFile("../endpoints.go", buf.Bytes(), 0644))
}

func constName(s string) string {
	a := strings.Split(s, "-")
	a[0] = strings.ToUpper(a[0])
	for k, v := range a[1:] {
		if x := v[0]; 'a' <= x && x <= 'z' {
			a[k+1] = string(x-'a'+'A') + v[1:]
		}
	}
	return "Region" + strings.Join(a, "")
}
