package app

import(

	"github.com/uprari/contextualSimulator/controllers"
)

func mapUrls() {

	router.GET("/ping", controllers.Ping)
	router.GET("/", controllers.LandingPage)
    router.GET("/categorise", controllers.Categorise)
}
