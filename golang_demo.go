package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

var (
	authStr string
)

func askForAuth(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Please Login"`)
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusUnauthorized)
}

func checkAuth(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")

	switch auth {
	case "": // no such header
		askForAuth(w)
		w.Write([]byte("no auth header received"))
	case authStr: // passed
		switch r.RequestURI {
		case "/hello":
			w.Write([]byte("hello"))
		case "/world":
			w.Write([]byte("world"))
		default:
			w.Write([]byte("test"))
		}
	default: // header existed but invalid
		askForAuth(w)
		w.Write([]byte("not authenticated"))
	}
}

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s port user password\n", os.Args[0])
		os.Exit(0)
	}

	authStr = calcAuthStr(os.Args[2], os.Args[3])

	host := fmt.Sprintf(":%s", os.Args[1])
	http.HandleFunc("/", checkAuth)
	http.ListenAndServe(host, nil)
}

func calcAuthStr(usr, pwd string) string {
	msg := []byte(usr + ":" + pwd)
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(msg)))
	base64.StdEncoding.Encode(encoded, msg)
	return "Basic " + string(encoded)
}
