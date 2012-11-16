package src

import (
	"appengine/user"
	"html/template"
	"net/http"
)

func phones(w http.ResponseWriter, r *http.Request) {
	content, context := getDefaultVar(r)
	ctmpl := "templates/phones.html"
	atmpl := ""

	if content["user"] == nil {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

	switch r.URL.Query().Get("a") {
	case "geol":
		atmpl = "templates/phones/info.html"
		Data := get10LastType("geoloc", (content["user"].(*user.User)).Email, context)
		content["data"] = Data
	case "ring":
	case "block":
	default:
		atmpl = "templates/phones/info.html"
		Data := get10LastType("status", (content["user"].(*user.User)).Email, context)
		content["data"] = Data

	}

	Tmpl := template.Must(template.ParseFiles("templates/base.html", ctmpl, atmpl))
	Tmpl.Execute(w, content)
}
