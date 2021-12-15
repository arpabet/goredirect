package main

import (
	"log"
	"net/http"
	"strings"
	"html/template"
)

type redirectPage struct {
	tpl   *template.Template

	to    string
	user  string
}

func (t *redirectPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	referenceName := strings.Trim(r.URL.Path, "/ ")
	moduleName := strings.Trim(r.Host, " ")
	projectName := referenceName

	params := map[string]string {
		"ModuleName": moduleName,
		"ReferenceName": referenceName,
		"GitHost": t.to,
		"GitUser": t.user,
		"ProjectName": projectName,
	}

	err := t.tpl.Execute(w, params)
	if err != nil {
		log.Printf("Serving error %v\n", err)
	}

}

func NewRedirectPage(templateIndex, to, user string) (http.Handler, error) {
	tpl, err := template.New("index").Parse(string(templateIndex))
	if err != nil {
		return nil, err
	}
	return &redirectPage{
		tpl: tpl,
		to: to,
		user: user,
	}, nil
}



