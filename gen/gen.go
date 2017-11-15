package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cloudogu/spinners"
)

const SPINNERS_URL = "https://raw.githubusercontent.com/sindresorhus/cli-spinners/master/spinners.json"

func main() {
	spinnerMap := readSpinnersFromUrl()

	model := struct {
		Timestamp time.Time
		URL       string
		Spinners  map[string]spinners.Configuration
	}{
		Timestamp: time.Now(),
		URL:       SPINNERS_URL,
		Spinners:  spinnerMap,
	}

	writeTemplate("spinners.tpl", model, "spinners.go")
	writeTemplate("readme.tpl", model, "README.md")
}

func readSpinnersFromUrl() map[string]spinners.Configuration {
	spinnerMap := make(map[string]spinners.Configuration)

	rsp, err := http.Get(SPINNERS_URL)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &spinnerMap)
	if err != nil {
		panic(err)
	}

	return spinnerMap
}

func writeTemplate(templateName string, model interface{}, outputPath string) {
	funcMap := template.FuncMap{
		"Title":   strings.Title,
		"ToUpper": strings.ToUpper,
		"Escape": func(str string) string {
			return strings.Replace(str, "\\", "\\\\", -1)
		},
	}

	tpl := template.Must(template.New(templateName).Funcs(funcMap).ParseFiles("gen/" + templateName))

	f, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tpl.Execute(f, model)
	if err != nil {
		panic(err)
	}
}
