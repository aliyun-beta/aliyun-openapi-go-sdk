package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	// "regexp"
	"sort"
	"strings"
)

var Types = map[string]string{
	"Boolean": "bool",
	"Integer": "int",
	"Long":    "int64",
	"Float":   "float32",
	"Double":  "float64",
	"String":  "string",
	"List":    "[]string",
}

type VersionInfo struct {
	Product  string
	Name     string
	APIStyle string
	APIs     struct {
		APIs []struct {
			Name string
		}
	}
	Pattern string // ROA
}

func (vi *VersionInfo) Package() string {
	n := len(vi.Product)
	a := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		if x := vi.Product[i]; 'A' <= x && x <= 'Z' {
			a = append(a, x-'A'+'a')
		} else if 'a' <= x && x <= 'z' {
			a = append(a, x)
		} else if '0' <= x && x <= '9' {
			a = append(a, x)
		}
	}
	if len(a) == 0 {
		Exitln("package", vi.Product)
	}
	return string(a)
}

type API struct {
	IsvProtocol struct {
		Method   string
		Pattern  string
		Protocol string
	}
	Name       string
	Parameters struct {
		Parameters []APIParameter
	}
	Product       string
	ResultMapping APIResultMapping
	Version       string
}

func (api *API) Clean() {
	api.IsvProtocol.Method = strings.ToUpper(CleanValues(api.IsvProtocol.Method, "|"))
	api.IsvProtocol.Protocol = strings.ToLower(CleanValues(api.IsvProtocol.Protocol, "|"))
	for k, v := range api.Parameters.Parameters {
		api.Parameters.Parameters[k].Value = CleanValues(v.Value, ",")
		api.Parameters.Parameters[k].TagPosition = strings.ToLower(strings.TrimSpace(v.TagPosition))
	}
}

func (api *API) GenerateParameters() (data [4][]string) {
	// method
	if strings.Contains(api.IsvProtocol.Method, "|") {
		data[2] = append(data[2], fmt.Sprintf(`if v, ok := optional["_method"]; ok {
	if s, ok := v.(string); ok {
		if !openapi.IsIn(s, %q) {
			return nil, errors.New("_method must be %s")
		}
		args.Method = s
	} else {
		return nil, errors.New("_method must be string")
	}
}`, api.IsvProtocol.Method, api.IsvProtocol.Method))
		data[3] = append(data[3], fmt.Sprintf("name: _method, type: string, optional values: %s", api.IsvProtocol.Method))
	} else if api.IsvProtocol.Method != "" {
		data[2] = append(data[2], fmt.Sprintf("args.Method = %q", api.IsvProtocol.Method))
	}

	// scheme
	if strings.Contains(api.IsvProtocol.Protocol, "|") {
		data[2] = append(data[2], fmt.Sprintf(`if v, ok := optional["_scheme"]; ok {
	if s, ok := v.(string); ok {
		if !openapi.IsIn(s, %q) {
			return nil, errors.New("_scheme must be %s")
		}
		args.Scheme = s
	} else {
		return nil, errors.New("_scheme must be string")
	}
}`, api.IsvProtocol.Protocol, api.IsvProtocol.Protocol))
		data[3] = append(data[3], fmt.Sprintf("name: _scheme, type: string, optional values: %s", api.IsvProtocol.Protocol))
	} else if api.IsvProtocol.Protocol != "" {
		data[2] = append(data[2], fmt.Sprintf("args.Scheme = %q", api.IsvProtocol.Protocol))
	}

	// region
	data[2] = append(data[2], `if v, ok := optional["_region"]; ok {
	if s, ok := v.(string); ok {
		args.Region = s
	} else {
		return nil, errors.New("_region must be string")
	}
}`)
	data[3] = append(data[3], fmt.Sprintf("name: _region, type: string"))

	// pattern
	if api.IsvProtocol.Pattern != "" {
		data[2] = append(data[2], fmt.Sprintf("args.Pattern = `%s`", api.IsvProtocol.Pattern))
	}

	// var dotted []APIParameter
	// statusKeyValueRegexp := regexp.MustCompile(`^Status(Key|Value)\d+$`)

	for _, i := range api.Parameters.Parameters {
		// if statusKeyValueRegexp.MatchString(i.TagName) { // ubsms, ubsms-inner
		// 	i.TagName = statusKeyValueRegexp.ReplaceAllStringFunc(i.TagName, func(s string) string {
		// 		return s + "."
		// 	})
		// }
		// if strings.Contains(i.TagName, ".") {
		// 	dotted = append(dotted, i)
		// }

		if i.Required == "true" {
			data[0] = append(data[0], i.Comment())
			data[1] = append(data[1], i.Parameter())
			if s := i.Validate(); s != "" {
				data[2] = append(data[2], s)
			}
		} else {
			data[2] = append(data[2], i.Validate())
			data[3] = append(data[3], i.Comment())
		}
	}

	for _, i := range data {
		sort.Strings(i)
	}

	n := len(data[1]) - 1
	for i := 0; i < n; i++ {
		x := strings.Index(data[1][i], " ")
		if data[1][i][x+1:] == data[1][i+1][strings.Index(data[1][i+1], " ")+1:] {
			data[1][i] = data[1][i][:x]
		}
	}

	data[1] = append(data[1], "optional openapi.M")

	return
}

func (api *API) GenerateResultMapping() (string, string) {
	if !(api.ResultMapping.TagName == "" && api.ResultMapping.ItemName == "") {
		Exitln("tag name and item name")
	}
	s := api.ResultMapping.Generate()
	if s == "" {
		return "openapi.Response", ""
	}
	buf := new(bytes.Buffer)
	name := MustExport(api.Name) + "Response"
	buf.WriteString(fmt.Sprintf("type %s struct {\n", name))
	// buf.WriteString("RequestId string\n")
	buf.WriteString(s)
	buf.WriteString("}\n")
	return name, buf.String()
}

type APIParameter struct {
	MaxValue    string `json:",omitempty"`
	MinValue    string `json:",omitempty"`
	Required    string
	TagName     string
	TagPosition string
	Type        string
	Value       string          `json:",omitempty"`
	ValueSwitch *APIValueSwitch `json:",omitempty"`
}

func (p APIParameter) ToArgs() string {
	t := MustType(p.Type)
	switch p.TagPosition {
	case "query", "body":
		s := "Query"
		if p.TagPosition == "body" {
			s = "Body"
		}
		switch t {
		case "int", "int64", "bool", "float32", "float64":
			return fmt.Sprintf("args.%s.Set(%q, fmt.Sprint(%s))", s, p.TagName, p.Name())
		case "string":
			return fmt.Sprintf("args.%s.Set(%q, %s)", s, p.TagName, p.Name())
		case "[]string":
			return fmt.Sprintf("args.%s[%q] = %s", s, p.TagName, p.Name())
		}
	case "path":
		switch t {
		case "int", "int64":
			return fmt.Sprintf("args.Path[%q] = fmt.Sprint(%s)", p.TagName, p.Name())
		case "string":
			return fmt.Sprintf("args.Path[%q] = %s", p.TagName, p.Name())
		case "bool", "float32", "float64", "[]string":
			Exitln("path parameter", t)
		}
	default:
		Exitln("position", p.TagPosition)
	}
	panic("never")
}

func (p APIParameter) Validate() string {
	var a []string
	if s := p.ValidateMinValue(); s != "" {
		a = append(a, s)
	}
	if s := p.ValidateMaxValue(); s != "" {
		a = append(a, s)
	}
	if s := p.ValidateValue(); s != "" {
		a = append(a, s)
	}

	a = append(a, p.ToArgs())

	if p.Required == "true" {
		return strings.Join(a, "\n")
	}

	t := MustType(p.Type)
	return fmt.Sprintf(`if v, ok := optional[%q]; ok {
	if %s, ok := v.(%s); ok {
		%s
	} else {
		return nil, errors.New("%s must be %s")
	}
}`, p.TagName, p.Name(), t, strings.Join(a, "\n"), p.TagName, t)
}

func IsNumber(t string) bool {
	switch MustType(t) {
	case "int", "int64", "float32", "float64":
		return true
	}
	return false
}

func (p APIParameter) ValidateMinValue() string {
	if p.MinValue != "" {
		if !IsNumber(p.Type) {
			Exitln(p.TagName, "must be number")
		}
		return fmt.Sprintf(`if %s < %s {
	return nil, errors.New("%s must be equal or greater than %s")
}`, p.Name(), p.MinValue, p.TagName, p.MinValue)
	}
	return ""
}

func (p APIParameter) ValidateMaxValue() string {
	if p.MaxValue != "" {
		if !IsNumber(p.Type) {
			Exitln(p.TagName, "must be number")
		}
		return fmt.Sprintf(`if %s > %s {
	return nil, errors.New("%s must be equal or less than %s")
}`, p.Name(), p.MaxValue, p.TagName, p.MaxValue)
	}
	return ""
}

func (p APIParameter) ValidateValue() string {
	if p.Value != "" {
		var s string
		switch MustType(p.Type) {
		case "int", "int64":
			s = "if !openapi.IsIn(fmt.Sprint(%s), %q) {\n"
		case "string":
			s = "if !openapi.IsIn(%s, %q) {\n"
		default:
			Exitln(p.TagName, "must be integers or string")
		}
		s += `return nil, errors.New("%s must be %s")`
		s += "\n}\n"
		return fmt.Sprintf(s, p.Name(), p.Value, p.TagName, p.Value)
	}
	return ""
}

func (p APIParameter) Name() string {
	name := p.TagName
	if name == "type" { // ace-ops
		name = "typ"
	} else {
		name = strings.Replace(name, ".", "", -1)
	}
	return name
}

func (p APIParameter) Parameter() string {
	return p.Name() + " " + MustType(p.Type)
}

func (p APIParameter) Comment() string {
	s := fmt.Sprintf("name: %s, type: %s", p.TagName, MustType(p.Type))
	if p.MinValue != "" {
		s += ", min value: " + p.MinValue
	}
	if p.MaxValue != "" {
		s += ", max value: " + p.MaxValue
	}
	if p.Value != "" {
		s += ", optional values: " + p.Value
	}
	return s
}

type APIValueSwitch struct {
	Cases []struct {
		TagValue string
	}
}

func (s *APIValueSwitch) Value() (a []string) {
loop:
	for _, v := range s.Cases {
		i := strings.TrimSpace(v.TagValue)
		if i != "" {
			for _, j := range a {
				if i == j {
					continue loop
				}
			}
			a = append(a, i)
		}
	}
	sort.Strings(a)
	return
}

type APIResultMapping struct {
	Arrays   []APIResultMapping `json:",omitempty"`
	Lists    APILists           `json:",omitempty"`
	Members  APIMembers         `json:",omitempty"`
	Structs  []APIResultMapping `json:",omitempty"`
	ItemName string             `json:",omitempty"`
	TagName  string             `json:",omitempty"`
}

func (r APIResultMapping) Generate() string {
	var a []string
	if len(r.Members) > 0 {
		a = append(a, r.Members.Generate()...)
	}
	if len(r.Lists) > 0 {
		a = append(a, r.Lists.Generate()...)
	}
	for _, i := range r.Arrays {
		tagName, itemName := MustExport(i.TagName), MustExport(i.ItemName)
		s := fmt.Sprintf("%s struct {\n", tagName)
		s += fmt.Sprintf("%s []struct {\n", itemName)
		s += i.Generate()
		if itemName == i.ItemName {
			s += "}\n"
		} else {
			s += fmt.Sprintf("} `json:%q xml:%q`\n", i.ItemName, i.ItemName)
		}
		if tagName == i.TagName {
			s += "}\n"
		} else {
			s += fmt.Sprintf("} `json:%q xml:%q`\n", i.TagName, i.TagName)
		}
		a = append(a, s)
	}
	for _, i := range r.Structs {
		if i.ItemName != "" {
			Exitln("item name", i.ItemName)
		}
		tagName := MustExport(i.TagName)
		s := fmt.Sprintf("%s struct {\n", tagName)
		s += i.Generate()
		if tagName == i.TagName {
			s += "}\n"
		} else {
			s += fmt.Sprintf("} `json:%q xml:%q`\n", i.TagName, i.TagName)
		}
		a = append(a, s)
	}
	a = RemoveDuplicate(a)
	return strings.Join(a, "")
}

type APIMembers []struct {
	TagName     string
	Type        string
	ValueSwitch *APIValueSwitch `json:",omitempty"`
}

func (a APIMembers) Generate() (b []string) {
	for _, i := range a {
		var value string
		if i.ValueSwitch != nil {
			value = strings.Join(i.ValueSwitch.Value(), "|")
			if value != "" {
				value = " // optional values: " + value
			}
		}
		tagName := MustExport(i.TagName)
		if tagName == i.TagName {
			b = append(b, fmt.Sprintf("%s %s%s\n", tagName, MustType(i.Type), value))
		} else {
			b = append(b, fmt.Sprintf("%s %s `json:%q xml:%q`%s\n", tagName, MustType(i.Type), i.TagName, i.TagName, value))
		}
	}
	b = RemoveDuplicate(b)
	return
}

type APILists []struct {
	ItemName string
	TagName  string
}

func (a APILists) Generate() (b []string) {
	for _, i := range a {
		tagName, itemName := MustExport(i.TagName), MustExport(i.ItemName)
		s := fmt.Sprintf("%s []struct{\n", tagName)
		if i.ItemName == itemName {
			s += fmt.Sprintf("%s string\n", itemName)
		} else {
			s += fmt.Sprintf("%s string `json:%q xml:%q`\n", itemName, i.ItemName, i.ItemName)
		}
		s += "}\n"
	}
	b = RemoveDuplicate(b)
	return
}

func RemoveDuplicate(a []string) []string {
	sort.Strings(a)
	n := len(a) - 1
	for i := 0; i < n; i++ {
		if a[i][:strings.Index(a[i], " ")] == a[i+1][:strings.Index(a[i+1], " ")] {
			// fmt.Println("duplicate", a[i], a[i+1])
			copy(a[i:], a[i+1:])
			n--
			i--
		}
	}
	return a[:n+1]
}

func MustType(t string) string {
	v, ok := Types[t]
	if !ok {
		Exitln("type", t)
	}
	return v
}

func MustExport(s string) string {
	var x byte
	if len(s) > 0 {
		x = s[0]
	}
	if 'a' <= x && x <= 'z' {
		return string(x-'a'+'A') + s[1:]
	} else if !('A' <= x && x <= 'Z') {
		Exitln("export", s)
	}
	return s
}

func Exitln(args ...interface{}) {
	fmt.Println(args...)
	os.Exit(1)
}

func ErrFatal(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func CleanValues(s string, sep string) string {
	a := strings.Split(s, sep)
	b := make([]string, 0, len(a))
loop:
	for _, i := range a {
		i = strings.TrimSpace(i)
		if i != "" {
			if strings.Contains(i, "|") {
				Exitln("values", s)
			}
			for _, j := range b {
				if i == j {
					continue loop
				}
			}
			b = append(b, i)
		}
	}
	sort.Strings(b)
	return strings.Join(b, "|")
}

func readJSON(f string, v interface{}) {
	data, err := ioutil.ReadFile(f)
	ErrFatal(err)

	{ // fix
		buf := new(bytes.Buffer)
		in := false
		for _, i := range data {
			switch i {
			case '"':
				buf.WriteByte(i)
				in = !in
			case '\n':
				if in {
					//buf.WriteString("\\n")
				} else {
					buf.WriteByte(i)
				}
			case '\t':
				if in {
					//buf.WriteString("\\t")
				} else {
					buf.WriteByte(i)
				}
			default:
				buf.WriteByte(i)
			}
		}
		if in {
			Exitln("fix", f)
		}
		data = buf.Bytes()
	}

	ErrFatal(json.Unmarshal(data, v))
}
