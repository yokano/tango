package tango

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", top)
}