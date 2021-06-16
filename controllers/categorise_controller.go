package controllers

import(

	 "net/http"
	 "github.com/gin-gonic/gin"
	 "github.com/uprari/contextualSimulator/globalData"
	 "fmt"
	 "hash/crc32"
	 "encoding/json"
	 "net/url"
)

var table *crc32.Table

func InitController(){

	table = crc32.MakeTable(88)

}
func Ping(c *gin.Context){

	c.String(http.StatusOK,"pong")
   
}

func LandingPage(c *gin.Context){

	c.HTML(http.StatusOK,"landing_page.html", gin.H{
			"title": "Main website",
		})
}



func Categorise(c *gin.Context){


	fmt.Println(c.Request)

	var data []byte
    var modulus uint32
    var index uint32
    service_source := c.Query("service")
	page_url := c.Query("s")

	_ ,err := url.ParseRequestURI( page_url )
	if err !=nil {
		fmt.Println(err)
		sendFailure(c,"Invalid page url")
		return
	}

	if len(glb.GumGum_data) == 0 ||len(glb.Peer39_data) == 0 {

		sendFailure(c,"Path Failure ")
		return
	}
	switch(service_source){
		case "gumgum":
			modulus = uint32(len(glb.GumGum_data))
    		index = crc32.Checksum([]byte(page_url),table)%modulus	
			data = []byte(glb.GumGum_data[index])
    		var v map[string]interface{}
    		json.Unmarshal(data, &v)
			v["pageUrl"] = page_url; 
			c.JSON(http.StatusOK,v)
		case "peer39":
			modulus = uint32(len(glb.Peer39_data))
    		index = crc32.Checksum([]byte(page_url),table)%modulus	
			data  = []byte(glb.Peer39_data[index])
    		var v []map[string]interface{}
    		json.Unmarshal(data, &v)
			c.JSON(http.StatusOK,v)
		default :
			sendFailure(c,"Data deficient; conversion failure")
			return
	}

}


func sendFailure(c *gin.Context, msg string){

	c.String(http.StatusNotFound, msg)

}
