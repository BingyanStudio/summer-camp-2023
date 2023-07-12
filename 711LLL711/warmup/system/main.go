package main

/*The entrypoint for the web app is main.go.
The file loads the application settings,
starts the session,
connects to the database,
sets up the templates,
loads the routes,
attaches the middleware,
and starts the web server.
*/
import (
	"system/app/route"
	"system/app/shared/database"
	"system/app/shared/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//load html
	server.Ginserver.LoadHTMLGlob("templates/*")
	// Connect to database
	database.Connect()
	//set the router
	route.Routes()
	//start the routes
	server.Ginserver.Run(":8080")
}
