package hi

import (
	"fmt"
	"io"
	"net/http"
)

func SayHi2(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	//notare che non gestiamo eventuali errori
	name := string(body)

	w.Header().Set("Content-type", "text/plain")
	fmt.Fprintf(w, "Hi %s!", name)

}
