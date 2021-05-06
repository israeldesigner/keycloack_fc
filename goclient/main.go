package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/net/html"
	"golang.org/x/oauth2"
)

var (
	clientID     = "myclient"
	clientSecret = "835cbe86-b0e3-43f7-9ecc-828e1a46e0e8"
)

func main() {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "http://127.0.0.1:8080/auth/realms/myrealm")
	if err != nil {
		log.Fatal(err)
	}

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:9900/auth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	state := "123"

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, config.AuthCodeURL(state), http.StatusFound)
		fmt.Fprintf(writer, "Hello, %q", html.EscapeString(request.URL.Path))
	})

	log.Fatal(http.ListenAndServe("9900", nil))

}
