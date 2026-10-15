package hi

import (
	"fmt"
	"net/http"
)

func SayHi(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hi %s", name)
}
