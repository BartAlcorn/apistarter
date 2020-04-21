package routes

import (
	// "fmt"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var html = `
<html>
<head>
  <title>Https Test</title>
</head>
<body style="padding:24px">
  <h2>There's no need to fear,</h2>
  <h1 style="color:#336699;">Bartman is here!</h1>
</body>
</html>
`

// Default displays the default page via GET
func Default(ctx *gin.Context) {
	ctx.Writer.WriteHeader(http.StatusOK)
	_, err := ctx.Writer.Write([]byte(html))
	if err != nil {
		fmt.Println("Writer Error", err)
	}
}
