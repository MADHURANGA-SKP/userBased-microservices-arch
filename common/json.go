package common

import (
	"github.com/gin-gonic/gin"
)

func WriteJSON(ctx *gin.Context, status int, data any) {
	ctx.JSON(status, data)
}

func ReadJSON(ctx *gin.Context, data any) error {
	if err := ctx.ShouldBindJSON(&data); err != nil {
		return err
	}

	return nil
}

func WriteError(ctx *gin.Context, status int, message string){
	WriteJSON(ctx, status, gin.H{"error":message})
}