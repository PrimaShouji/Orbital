package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/PrimaShouji/Orbital/server/app"
	"github.com/PrimaShouji/Orbital/server/middleware/isauth"
	"github.com/PrimaShouji/Orbital/server/routes/authenticated"
	"github.com/PrimaShouji/Orbital/server/routes/callback"
	"github.com/PrimaShouji/Orbital/server/routes/login"
	"github.com/PrimaShouji/Orbital/server/routes/logout"
	"github.com/gin-gonic/gin"
)

func proxyHandler() gin.HandlerFunc {
	r, err := url.Parse("http://localhost:4200")
	if err != nil {
		log.Fatalln(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(r)
	return gin.WrapH(proxy)
}

func main() {
	// Parse command-line args
	port := flag.Uint64("port", 7558, "application binding port")

	flag.Parse()

	// Initialize application
	a := app.Init(*port)

	a.Get("/callback", callback.CallbackHandler)
	a.Get("/login", login.LoginHandler)
	a.Get("/logout", logout.LogoutHandler)
	a.Get("/authenticated", authenticated.AuthenticatedHandler)

	admin := a.Group("/admin", isauth.IsAuthenticated)
	admin.Get("/test", func(c *gin.Context) {
		c.String(200, "Authenticated")
	})

	a.Any("/dashboard/*proxyPath", proxyHandler())
	a.Any("/", func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/dashboard")
	})

	a.Run()
}
