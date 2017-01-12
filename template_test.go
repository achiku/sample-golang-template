package sampletemplate

import "testing"

func TestFillTemplate(t *testing.T) {
	d := dict{
		"key1": "this is the first key",
		"key2": "this is the second key",
		"key3": "111",
		"key4": "日本語",
	}
	msg := "{{.key1}} and {{.key2}}, {{.key3}}, {{.key4}}, {{.noKey}}"
	s, err := fill(msg, d)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(s)
}

func TestToJSON(t *testing.T) {
	d := dict{
		"key1": "this is the first key",
		"key2": "this is the second key",
		"key3": "111",
		"key4": "日本語",
	}
	j, err := d.toJSON()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(j)
}
