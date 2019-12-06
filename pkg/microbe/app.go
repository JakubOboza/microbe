package microbe

import "github.com/gin-gonic/gin"

//App base app wrapping type
type App struct {
	router *gin.Engine
}

//NewMicrobe creates app instance
func NewMicrobe() *App {
	return &App{}
}

//Run stats main loop of the app
func (app *App) Run() error {
	err := app.initialize()
	if err != nil {
		return err
	}
	defer app.close()

	app.router.Run()

	return nil
}

func (app *App) initialize() error {
	// initialize stuff
	app.router = gin.Default()

	app.setupRoutes()

	return nil
}

func (app *App) close() {
	// cleanup
}
