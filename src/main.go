package src

// Launch
import (
	"net/http"
)

func init() {
	http.HandleFunc("/", pages)
}
