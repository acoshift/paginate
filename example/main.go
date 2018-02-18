package main

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/acoshift/paginate"
)

func main() {
	tmpl := template.Must(template.New("").Parse(`
		<!doctype html>
		<meta charset="utf-8">
		<style>
			div {
				display: flex;
				flex-wrap: wrap;
			}

			a {
				display: inline-block;
				padding: 8px 14px;
				color: black;
			}

			a:not(:first-child) {
				border-top: 1px solid blue;
				border-right: 1px solid blue;
				border-bottom: 1px solid blue;
			}

			a:first-child {
				border: 1px solid blue;
			}

			a.disabled {
				color: gray;
				pointer-events: none;
			}

			a:not(.disabled) {
				cursor: pointer;
			}

			a:not(.disabled):hover {
				background-color: rgb(230, 230, 230);
			}

			a.active {
				background-color: blue;
				color: white;
				pointer-events: none;
			}
		</style>
		<div>
			<a href="?page=1">First</a>
			<a href="?page={{.Prev}}">Prev</a>
			{{range .Pages 2 2}}
				{{if eq . 0}}
					<a class="disabled">...</a>
				{{else if eq $.Page .}}
					<a class="active">{{.}}</a>
				{{else}}
					<a href="?page={{.}}">{{.}}</a>
				{{end}}
			{{end}}
			<a href="?page={{.Next}}">Next</a>
			<a href="?page={{.MaxPage}}">Last</a>
		</div>
	`))
	http.ListenAndServe(":3000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page, _ := strconv.ParseInt(r.FormValue("page"), 10, 64)
		pn := paginate.New(page, 10, 123)
		w.Header().Set("Content-Type", "text/html")
		tmpl.Execute(w, pn)
	}))
}
