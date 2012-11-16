package src

import (
	"html/template"
	"net/http"
)

func admin(w http.ResponseWriter, r *http.Request) {
	content, context := getDefaultVar(r)

	if content["isAdmin"].(string) != "true" {
		http.Redirect(w, r, "/notfound", http.StatusNotFound)
		return
	}

	content["ApiKey"] = getLastStoredKey(context)

	ctmpl := "templates/i18n/admin." + content["lang"].(string) + ".html"
	Tmpl := template.Must(template.ParseFiles("templates/base.html", ctmpl))
	Tmpl.Execute(w, content)
}

/*func to handle web storage*/
func apiKey(w http.ResponseWriter, r *http.Request) {
	user, _, c := getLoginData(r)
	if user == nil || !user.Admin {
		http.Redirect(w, r, "/notfound", http.StatusNotFound)
	}

	if r.Method == "POST" && r.FormValue("key") != "" {
		StoreApiKey(r.FormValue("key"), c)
		http.Redirect(w, r, "/admin?msg=Ok", http.StatusMovedPermanently)
	} else {
		http.Redirect(w, r, "/admin", http.StatusMovedPermanently)
	}
}
