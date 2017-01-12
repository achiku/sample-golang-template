package sampletemplate

import (
	"bytes"
	"html/template"
)

type dict map[string]interface{}

func fill(msg string, d dict) (string, error) {
	b := &bytes.Buffer{}
	t, err := template.New("").Parse(msg)
	if err != nil {
		return "", err
	}
	if err := t.Execute(b, d); err != nil {
		return "", err
	}
	return b.String(), nil
}
