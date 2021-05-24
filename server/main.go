package main

import (
	"flag"

	"github.com/PrimaShouji/Orbital/server/app"
	"github.com/PrimaShouji/Orbital/server/routes/callback"
	"github.com/PrimaShouji/Orbital/server/routes/login"
	"github.com/PrimaShouji/Orbital/server/routes/logout"
)

func main() {
	// Parse command-line args
	port := flag.Uint64("port", 7558, "application binding port")

	flag.Parse()

	// Initialize application
	a := app.Init(*port)
	a.Get("/callback", callback.CallbackHandler)
	a.Get("/login", login.LoginHandler)
	a.Get("/logout", logout.LogoutHandler)
	a.Run()
}
