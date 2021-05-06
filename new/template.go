package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"time"
)

const temp1 = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

type Detail struct {
	Number    int
	User      string
	Title     string
	CreatedAt time.Time
}

type Report struct {
	TotalCount int
	Items      []Detail
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	data := Report{
		TotalCount: 1,
		Items: []Detail{
			{Number: 1, User: "zhang", Title: "teacher", CreatedAt: time.Now()},
			{Number: 2, User: "li", Title: "docker", CreatedAt: time.Now()},
		},
	}
	fmt.Println(data)

	report, err := template.New("data").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(temp1)
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
