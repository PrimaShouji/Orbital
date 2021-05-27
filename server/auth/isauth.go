package auth

import (
	"github.com/PrimaShouji/Orbital/server/app"
	"github.com/gin-gonic/gin"
)

func IsAuthenticated(c *gin.Context) (bool, error) {
	session, err := app.Store.Get(c.Request, "auth-session")
	if err != nil {
		return false, err
	}

	if _, ok := session.Values["profile"]; !ok {
		return false, nil
	}

	return true, nil
}
