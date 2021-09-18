package main

import (
	"iplocator/httpd/handler"
	"iplocator/httpd/middleware"
	"os"
	"io"
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"bytes"
	"log"
)

func main() {

	gin.DisableConsoleColor()

    // separate request and error logs and write to separate files.
	logfile, _ := os.Create("request.log")
	gin.DefaultWriter = io.MultiWriter(logfile)

	errlogfile, _ := os.Create("error.log")
	gin.DefaultErrorWriter = io.MultiWriter(errlogfile)

	// set to release mode for clean logs
	gin.SetMode( gin.ReleaseMode )
	r := gin.New()
	r.Use( gin.Recovery() ) //recover in case of panic
	r.Use(RequestLogger()) // Use custom middleware for logging the payload and other info


	//create the cache for the requests
	myCache := cache.New(10*time.Minute, 10*time.Minute)

	//endpoints
	r.POST( "/locate", middleware.CacheCheck(myCache) ,handler.LocateIp( ) )

	r.Run( "0.0.0.0:8080" ) // listen and serve on 0.0.0.0:8080
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

		body, _ := ioutil.ReadAll(rdr1)

		c.Request.Body = rdr2
		c.Next()
		duration := GetDurationInMillseconds(start)

		log.SetOutput(gin.DefaultWriter)
		log.Println( fmt.Sprintf("ipAddress: %s, duration: %f, payload: %s\n",
							c.ClientIP(),
							duration,
							body,
		));
	}
}

func GetDurationInMillseconds(start time.Time) float64 {
	end := time.Now()
	duration := end.Sub(start)
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	return rounded
}