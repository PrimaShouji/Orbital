package authenticated

import (
	"net/http"

	"github.com/PrimaShouji/Orbital/server/auth"
	"github.com/gin-gonic/gin"
)

type AuthenticationState struct {
	Authenticated bool `json:"is_authenticated"`
}

func AuthenticatedHandler(c *gin.Context) {
	authenticated, err := auth.IsAuthenticated(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &AuthenticationState{Authenticated: authenticated})
}
