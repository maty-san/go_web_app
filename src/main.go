package main

import ( "github.com/gin-gonic/gin" )

func main() { r := gin.Default()

```text
r.GET("/", func(c *gin.Context) {
	c.String(200, "Hello world")
})

r.GET("/hogehoge", func(c *gin.Context) {
	c.String(200, "hogehoge")
})

r.Run(":8080111111")
```

}
