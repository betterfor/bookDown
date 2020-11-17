/**
 *Created by XieJian on 2020/11/17 17:05
 *@Desc: template files
 */
package tpl

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//
func LoadTemplate(pattern string) *template.Template {
	t := template.New("").Funcs(FuncMap())

	var files []string

	filepath.Walk(pattern, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		path = strings.ReplaceAll(path, "\\", "/")
		name := strings.TrimSuffix(strings.TrimPrefix(path, "templates/"), ".tmpl")
		bts, _ := ioutil.ReadFile(path)
		t, err = t.New(name).Parse(string(bts))
		if err != nil {
			return err
		}
		files = append(files, fmt.Sprintf("\t- %s", info.Name()))
		return nil
	})

	fmt.Printf("[GIN-debug] Loaded HTML Templates (%d)\n%s\n\n", len(files), strings.Join(files, "\n"))

	return t
}
