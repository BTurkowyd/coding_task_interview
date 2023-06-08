package funcs

import (
	"github.com/gin-gonic/gin"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func SetCookie(c *gin.Context) {
	c.SetCookie("my-first-cookie", "123456", 24*60*60, "/", "localhost", false, true)
}
