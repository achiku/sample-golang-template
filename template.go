package sampletemplate

import (
	"bytes"
	"encoding/json"
	"html/template"
)

type dict map[string]string

func (d *dict) toJSON() (string, error) {
	s, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

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
