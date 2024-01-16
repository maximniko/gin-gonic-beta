package handlers

import (
	"awesomeProject/src/interfaces/errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func HandleRuntimePanic(c *gin.Context) {
	if r := recover(); r != nil {
		switch r.(type) {
		case errors.PanicResponse:
			tmp := r.(errors.PanicResponse)
			log.Println(fmt.Printf("PanicResponse with code: %d\n", tmp.Code()))
			c.JSON(tmp.Code(), tmp.Object())
			c.Abort()
		case string:
			tmp := r.(string)
			log.Println(tmp)
			c.JSON(1, tmp)
			c.Abort()
		}
	}
}
