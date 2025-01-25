package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name" binding:"required"`        //必須
	Email string `json:"email" binding:"required,email"` //必須、有効なメールアドレス形式
	Age   int    `json:"age" binding:"gte=18"`           //18以上
}

func main() {
	r := gin.Default()

	// POST エンドポイント
	r.POST("/gin", func(c *gin.Context) {
		var user User

		// ShouldBindJSON は、JSON データをバインドして構造体に格納します。
		if err := c.ShouldBindJSON(&user); err != nil {
			// バインドに失敗した場合、エラーを返します。
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
	})

	r.Run(":8081")
}
