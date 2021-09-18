package middleware

import (
	"io/ioutil"
	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

// bodyCacheWriter is used to cache responses in gin.
type bodyCacheWriter struct {
	gin.ResponseWriter
	cache      *cache.Cache
	payload string
}

// Write a JSON response to gin and cache the response.
func (w bodyCacheWriter) Write(b []byte) (int, error) {
	// Write the response to the cache only if a success code
	status := w.Status()
	if 200 <= status && status <= 299 {
		w.cache.Set(w.payload, b, cache.DefaultExpiration)
	}

	// Then write the response to gin
	return w.ResponseWriter.Write(b)
}

// CacheCheck sees if there are any cached responses and returns
// the cached response if one is available.
func CacheCheck(cache *cache.Cache) gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

		body, _ := ioutil.ReadAll(rdr1)

		c.Request.Body = rdr2
		// See if we have a cached response
		response, exists := cache.Get(string(body))
		if exists {
			// If so, use it
			c.Data(200, "application/json", response.([]byte))
			c.Abort()
		} else {
			// If not, pass our cache writer to the next middleware
			bcw := &bodyCacheWriter{cache: cache, payload: string(body), ResponseWriter: c.Writer}
			c.Writer = bcw
			c.Next()
		}
	}
}