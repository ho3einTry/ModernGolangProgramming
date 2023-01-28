package dynowebportal

import (
	"fmt"
	"net/http"
)

func RunWebPortal(addr string) error {
	http.HandleFunc("/", rootHandler)
	return http.ListenAndServe(addr, nil)
}
func rootHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Welcome to the Dino web portal %s", r.RemoteAddr)
}
