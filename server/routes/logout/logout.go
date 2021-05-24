package logout

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {
	r := c.Request
	w := c.Writer

	domain := "prima-shouji.us.auth0.com" // TODO: extract to config

	logoutUrl, err := url.Parse("https://" + domain)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logoutUrl.Path += "/v2/logout"
	parameters := url.Values{}

	var scheme string
	if r.TLS == nil {
		scheme = "http"
	} else {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + r.Host)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", "dwjT0bzwLM41mboU3VCZq6Hbb8uVSgsi") // TODO: extract to config
	logoutUrl.RawQuery = parameters.Encode()

	http.Redirect(w, r, logoutUrl.String(), http.StatusTemporaryRedirect)
}
