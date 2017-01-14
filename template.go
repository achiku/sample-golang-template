package sampletemplate

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"os"
	"reflect"
	"strings"
)

type dict map[string]interface{}

var funcMap = template.FuncMap{
	"title": strings.Title,
	"default": func(arg, value interface{}) interface{} {
		v := reflect.ValueOf(value)
		switch v.Kind() {
		case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
			if v.Len() == 0 {
				return arg
			}
		case reflect.Bool:
			if !v.Bool() {
				return arg
			}
		default:
			return value
		}
		return value
	},
}

func newDict(data []byte) (dict, error) {
	var d dict
	if err := json.Unmarshal(data, &d); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *dict) toJSON() (string, error) {
	s, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

func fill(msg string, d dict) (string, error) {
	b := &bytes.Buffer{}
	t, err := template.New("").Funcs(funcMap).Parse(msg)
	if err != nil {
		return "", err
	}
	if err := t.Execute(b, d); err != nil {
		return "", err
	}
	return b.String(), nil
}

func te() {
	// First we create a FuncMap with which to register the function.
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"title": strings.Title,
	}

	// A simple template definition to test our function.
	// We print the input text several ways:
	// - the original
	// - title-cased
	// - title-cased and then printed with %q
	// - printed with %q and then title-cased.
	const templateText = `
Input: {{printf "%q" .}}
Output 0: {{title .}}
Output 1: {{title . | printf "%q"}}
Output 2: {{printf "%q" . | title}}
`

	// Create a template, add the function map, and parse the text.
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// Run the template to verify the output.
	err = tmpl.Execute(os.Stdout, "the go programming language")
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
}
