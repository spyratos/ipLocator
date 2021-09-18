package handler

import (
	"net/http"
	"log"
	"net"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"encoding/json"
)

type Login struct {
	Ip     string `json:"ip" binding:"required"`
}


func LocateIp( ) gin.HandlerFunc { //handler function for the actual request
	return func( c *gin.Context ) {
		log.SetOutput(gin.DefaultErrorWriter) //all logs here are related to errors

		var paylaod Login

		if err := c.ShouldBindJSON(&paylaod); err != nil { //cover that ip field exists in the POST request
			log.Println( err.Error() );
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ip := paylaod.Ip

		if !checkIPAddress(ip) { //check validity of ip address, because the public API never returns an error
			c.JSON( http.StatusOK, gin.H {// and takes the ip from headers (our service ip) if no ip is provided by the user
				"error": "ip is invalid",
			})
			return
		}

		res, err := http.Get("https://freegeoip.app/json/"+ip)//request to 3rd party

		if err != nil {// just in  case 3rd party goes down
			log.Println( err.Error() );
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//read body into map
		data, _ := ioutil.ReadAll( res.Body )
		res.Body.Close()
		var dat map[string]interface{}

		err = json.Unmarshal(data, &dat)
		if err != nil {
			log.Println( err.Error() );
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON( http.StatusOK, gin.H {//final response
			"data": map[string]string {
				"country_name": dat["country_name"].(string),
				"city": dat["city"].(string),
			},
		})

	}
}

func checkIPAddress(ip string) bool {
	if net.ParseIP(ip) == nil {
			return false
	} else {
			return true
	}
}
