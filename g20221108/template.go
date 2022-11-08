package g20221108

import (
	"log"
	"os"
	"text/template"
)

type TestData struct {
	Name       string `json:"name"`
	Attributes map[string]interface{}
	Address    *struct {
		Street string
	}
}

func testTemplate() {
	var Template = `apiVersion: apps/v1
kind: {{.Name}}
metadata:
  name: nginx-deployment
spec:
  selector:
    matchLabels:
      app: nginx
  replicas: 2 # tells deployment to run 2 pods matching the template
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: nginx:1.14.2
          ports:
            - containerPort: 80


       {{ range $key, $value := .Attributes }}
        - name: {{ $key }}
          value: {{ $value }}
      {{ end }}

  {{if .Address}}
  Apt is {{.Address.Street}}.
  {{else}}
  Apt is NOT set.
  {{end}}
`
	tmpl, err := template.New("test").Parse(Template)
	if err != nil {
		log.Println(err)
		return
	}
	data := TestData{
		Name: "hello",
		Attributes: map[string]interface{}{
			"key":   "hello",
			"value": "json",
		},
		Address: &struct {
			Street string
		}{
			Street: "牛逼",
		},
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Println(err)
		return
	}

}
