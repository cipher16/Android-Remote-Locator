package src

// Launch
import (
	"net/http"
)

var TEMPLATE = "templates"
var TEMPLATE_I18N = TEMPLATE + "/i18n"

func init() {
	http.HandleFunc("/apiKeyStore", apiKey)
	http.HandleFunc("/gcm", GCM)
	http.HandleFunc("/phones", phones)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/", pages)
}
