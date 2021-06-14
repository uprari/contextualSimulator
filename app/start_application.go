package app

import (
  "github.com/gin-gonic/gin"
  "github.com/uprari/contextualSimulator/controllers"
  "github.com/uprari/contextualSimulator/globalData"
)

var (

  router = gin.Default()
)





func StartApplication(){


    glb.Init_response_data()
    controllers.InitController()
    router.LoadHTMLFiles("../html/landing_page.html")	
	mapUrls()
	router.Run(":9434")

}



