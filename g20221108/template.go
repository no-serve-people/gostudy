package g20221108

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type TestData struct {
	Name       string `json:"name"`
	Attributes map[string]interface{}
}

func testTemplate() {
	var Template = `{{ range $key, $value := .Attributes }}
	{{ $.Name}} {{ $key }}: {{ $value }}{{- end }}`
	tmpl, err := template.New("test").Parse(strings.TrimSpace(Template))
	if err != nil {
		log.Println(err)
		return
	}
	data := TestData{
		"hello",
		map[string]interface{}{
			"key":   "hello",
			"value": "json",
		},
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Println(err)
		return
	}

}
