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
	m := make(map[string]map[string]string)
	for _, i := range v.Endpoint {
		if i.Name == "cn-hangzhou" {
			buf.WriteString("\n// Region List\n")
			buf.WriteString("const (\n")
			for _, j := range i.RegionIds.RegionId {
				buf.WriteString(fmt.Sprintf("Region%s = %q\n", constName(j), j))
			}
			buf.WriteString(")\n")
		} else if !(len(i.RegionIds.RegionId) == 1 && i.RegionIds.RegionId[0] == i.Name) {
			Exitln(i.Name)
		}

		if _, ok := m[i.Name]; ok {
			Exitln(i.Name)
		}
		x := make(map[string]string, len(i.Products.Product))
		for _, j := range i.Products.Product {
			if _, ok := m[j.ProductName]; ok {
				Exitln(i.Name, j.ProductName)
			}
			x[j.ProductName] = j.DomainName
		}
		m[i.Name] = x
	}
	a := make([]string, 0, len(m))
	for k := range m {
		a = append(a, k)
	}
	sort.Strings(a)
	buf.WriteString("\n// Domains [region][product]\n")
	buf.WriteString("var Domains =  map[string]map[string]string{\n")
	for _, k := range a {
		v := m[k]
		b := make([]string, 0, len(v))
		for i := range v {
			b = append(b, i)
		}
		sort.Strings(b)
		buf.WriteString(fmt.Sprintf("%q: {\n", k))
		for _, i := range b {
			buf.WriteString(fmt.Sprintf("%q: %q,\n", i, v[i]))
		}
		buf.WriteString("},\n")
	}
	buf.WriteString("}\n")
	buf.WriteString(`
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
	return strings.Join(a, "")
}
