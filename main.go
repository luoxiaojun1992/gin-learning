package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	// Without Middleware
	// r := gin.New()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	adminRouter := r.Group("/administrator/")
	adminRouter.GET("test", func(c *gin.Context) {
		var Person struct {
			Name    string `form:"name" binding:"required"`
			Address string `form:"address" binding:"required"`
		}

		err := c.BindQuery(&Person)
		if err == nil {
			c.JSON(http.StatusOK, Person)
		} else {
			c.String(http.StatusBadRequest, err.Error())
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	// Sample: curl -X POST "http://localhost:9999/admin" -u foo:bar -d "Value={\"foo\":\"bar\"}"
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func run(r *gin.Engine) {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}

	_, port, err := net.SplitHostPort(ln.Addr().String())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening:" + port)

	http.Serve(ln, r)
}

func main() {
	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	// r.Run(":9999")
	run(r)
}
