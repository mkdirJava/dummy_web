package server

import (
	"io"
	"net/http"

	getstring "github.com/mkdirjava/dummy_domain"
)

var getStringer = getstring.NewGetStringerService()

func StartServer() {
	http.HandleFunc("/test", func(write http.ResponseWriter, request *http.Request) {
		io.WriteString(write, getStringer.GetString())
	})
	http.ListenAndServe(":8080", nil)
}
