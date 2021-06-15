package app

import (
  "github.com/gin-gonic/gin"
  "github.com/uprari/contextualSimulator/controllers"
  "github.com/uprari/contextualSimulator/globalData"
  "os"
  "fmt"
)

var (

  router = gin.Default()
)





func StartApplication(){

    glb.Init_response_data()
    controllers.InitController()
    path, err := os.Getwd()
	fmt.Println(path)	
    addr, err := determineListenAddress()
 	if err != nil {
    	fmt.Println(err)
  	}
	
    router.LoadHTMLFiles("/html/landing_page.html")	
	mapUrls()
	router.Run(addr)

}


func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}
