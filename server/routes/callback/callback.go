package callback

import (
	"log"
	"net/http"

	"github.com/PrimaShouji/Orbital/server/app"
	"github.com/PrimaShouji/Orbital/server/auth"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
)

func CallbackHandler(c *gin.Context) {
	r := c.Request
	w := c.Writer

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.URL.Query().Get("state") != session.Values["state"] {
		http.Error(w, "Invalid state parameter", http.StatusBadRequest)
		return
	}

	authenticator, err := auth.NewAuthenticator()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := authenticator.Config.Exchange(authenticator.Ctx, r.URL.Query().Get("code"))
	if err != nil {
		log.Printf("no token found: %v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}

	oidcConfig := &oidc.Config{
		ClientID: "dwjT0bzwLM41mboU3VCZq6Hbb8uVSgsi", // Extract to config
	}

	idToken, err := authenticator.Provider.Verifier(oidcConfig).Verify(authenticator.Ctx, rawIDToken)

	if err != nil {
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Getting now the userInfo
	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["id_token"] = rawIDToken
	session.Values["access_token"] = token.AccessToken
	session.Values["profile"] = profile
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to logged-in page
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
