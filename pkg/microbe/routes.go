package microbe

func (app *App) setupRoutes() {
	app.router.GET("/ping", app.pingHandler)
	app.router.POST("/api/base64/encode", app.encodeBase64Handler)
	app.router.POST("/api/base64/decode", app.decodeBase64Handler)
}
