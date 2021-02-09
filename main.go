package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/kr/pretty"
)

var (
	target string
	token  string
	auth   string
	u      *url.URL

	authTypes = map[string]interface{}{
		"basic": nil,
		"token": nil,
	}
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%#v\n", pretty.Formatter(r))

	r.Host = u.Host
	if auth == "basic" {
		r.SetBasicAuth(token, "x-oauth-basic")
	} else {
		r.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	}

	httputil.NewSingleHostReverseProxy(u).ServeHTTP(w, r)
}

func main() {
	flag.StringVar(&target, "target", "https://github.com", "target URL which is transported to")
	flag.StringVar(&token, "token", "", "GITHUB_TOKEN")
	flag.StringVar(&auth, "auth", "basic", "auth type (basic or token)")
	flag.Parse()

	var err error
	u, err = url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}

	if token == "" {
		log.Fatal("-token is required")
	}

	if _, ok := authTypes[auth]; !ok {
		log.Fatalf("invalid auth type: %s\n", auth)
	}

	http.HandleFunc("/", handler)
	if err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil); err != nil {
		log.Fatal(err)
	}
}
