package sampletemplate

import (
	"bytes"
	"encoding/json"
	"html/template"
)

type dict map[string]interface{}

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
	t, err := template.New("").Parse(msg)
	if err != nil {
		return "", err
	}
	if err := t.Execute(b, d); err != nil {
		return "", err
	}
	return b.String(), nil
}
