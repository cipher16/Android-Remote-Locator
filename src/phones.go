package src

import (
	"appengine"
	"appengine/user"
	"html/template"
	"net/http"
)

func Phones(w http.ResponseWriter, r *http.Request, lang string, user *user.User, context appengine.Context, url string) {
	page := r.URL.Path
	content := map[string]interface{}{
		"i18n":        i18n[lang],
		"user":        user,
		"urlLogin":    url,
		"currentpage": page,
	}
	ctmpl := "templates/phones.html"
	Tmpl := template.Must(template.ParseFiles("templates/base.html", ctmpl))
	Tmpl.Execute(w, content)
}
