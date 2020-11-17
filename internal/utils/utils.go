/**
 *Created by Xie Jian on 2019/8/9 13:17
 */
package utils

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func TmplFuncMap() template.FuncMap {
	return template.FuncMap{
		"trim":       strings.TrimSpace,
		"bold":       Bold,
		"headline":   MagentaBold,
		"foldername": RedBold,
		"endline":    EndLine,
		"tmpltostr":  TemplToString,
	}
}

func TemplToString(tmpl string, data interface{}) string {
	t := template.New("tmpl").Funcs(TmplFuncMap())
	template.Must(t.Parse(tmpl))

	var doc bytes.Buffer
	err := t.Execute(&doc, data)
	if err != nil {
		panic(err)
	}

	return doc.String()
}

func Tmpl(text string, data interface{}) {
	t := template.New("Usage").Funcs(TmplFuncMap())
	template.Must(t.Parse(text))

	err := t.Execute(os.Stderr, data)
	if err != nil {
		panic(err)
	}
}

func PrintErrorAndExit(message, errorTemplate string) {
	Tmpl(fmt.Sprintf(errorTemplate, message), nil)
	os.Exit(2)
}
