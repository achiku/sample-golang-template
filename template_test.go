package sampletemplate

import "testing"

func TestFillTemplate(t *testing.T) {
	d := dict{
		"key1":   "this is the first key",
		"key2":   "this is the second key",
		"key3":   "111",
		"key4":   "日本語",
		"key100": "無い",
	}
	msg := `{{.key1}} and {{.key2}}, {{.key3}}, {{.key4}}, {{.noKey}} {{default .noKey "val"}}`
	s, err := fill(msg, d)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(s)
}

func TestFillTemplateFuncDefault(t *testing.T) {
	d := dict{
		"key1": "this is the first key",
		"key2": "this is the second key",
		"key3": "",
	}
	msg := `{{"output" | title }} and {{.key2 | title }}, and {{ .key3 | default "default" }}`
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

func TestNewDict(t *testing.T) {
	raw := []byte(
		`{"key1":"this is the first key","key2":"this is the second key","key3":"111","key4":"日本語"}`)
	d, err := newDict(raw)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", d)
}

func TestTe(t *testing.T) {
	te()
}
