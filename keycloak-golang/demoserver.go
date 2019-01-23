package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

func main() {
	serverUrl := "http://localhost" + port()
	const keycloakBaseUrl = "http://localhost:8081"
	const keycloakRealm = "KeycloakExampleApp"

	configURL := keycloakBaseUrl+"/auth/realms/"+keycloakRealm
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, configURL)
	if err != nil {
		panic(err)
	}

	clientID := "login-app"
	clientSecret := "" //"cbfd6e04-a51c-4982-a25b-7aaba4f30c81"

	redirectURL := serverUrl+"/"+keycloakRealm+"/callback"
	// Configure an OpenID Connect aware OAuth2 client.
	oauth2Config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),
		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}
	state := "somestate"

	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	verifier := provider.Verifier(oidcConfig)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rawAccessToken := r.Header.Get("Authorization")
		if rawAccessToken == "" {
			http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
			return
		}

		parts := strings.Split(rawAccessToken, " ")
		if len(parts) != 2 {
			w.WriteHeader(400)
			return
		}
		_, err := verifier.Verify(ctx, parts[1])

		if err != nil {
			http.Redirect(w, r, oauth2Config.AuthCodeURL(state), http.StatusFound)
			return
		}

		w.Write([]byte("hellooo"))
	})

	http.HandleFunc("/"+keycloakRealm+"/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state {
			http.Error(w, "state did not match", http.StatusBadRequest)
			return
		}

		oauth2Token, err := oauth2Config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
			return
		}
		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		resp := struct {
			OAuth2Token   *oauth2.Token
			IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
		}{oauth2Token, new(json.RawMessage)}

		if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	fmt.Printf("Server url -> %v \n", serverUrl)
	err2 := http.ListenAndServe(port(), nil)
	checkError(err2)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8026"
	}
	return ":" + port
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

