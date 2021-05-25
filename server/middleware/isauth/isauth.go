package isauth

import (
	"log"
	"net/http"

	"github.com/PrimaShouji/Orbital/server/app"
	"github.com/gin-gonic/gin"
)

func IsAuthenticated(c *gin.Context) {
	r := c.Request
	w := c.Writer

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := session.Values["profile"]; !ok {
		log.Println("Unauthorized request made; redirecting...")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		c.Next()
	}
}
