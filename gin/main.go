package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name  string `json:"name" binding:"required"`        // 必須
	Email string `json:"email" binding:"required,email"` // 必須、有効なメールアドレス形式
	Age   int    `json:"age" binding:"gte=18"`           // 18以上
}

// カスタムエラーメッセージ
var customMessages = map[string]string{
	"Name.required":  "名前は必須です",
	"Email.required": "メールアドレスは必須です",
	"Email.email":    "有効なメールアドレスを入力してください",
	"Age.gte":        "年齢は18歳以上である必要があります",
}

func main() {
	r := gin.Default()

	r.POST("/gin", func(c *gin.Context) {
		var user User

		// JSONを構造体にバインド
		if err := c.ShouldBindJSON(&user); err != nil {
			// バリデーションエラーを取得
			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				// カスタムエラーメッセージを適用
				errorsMap := make(map[string]string)
				for _, e := range validationErrors {
					// フィールド名とタグを組み合わせてカスタムメッセージを取得
					key := e.Field() + "." + e.Tag()
					if msg, exists := customMessages[key]; exists {
						errorsMap[e.Field()] = msg
					} else {
						// デフォルトメッセージ
						errorsMap[e.Field()] = e.Error()
					}
				}
				// カスタムエラーをレスポンスに返す
				c.JSON(http.StatusBadRequest, gin.H{"errors": errorsMap})
				return
			}

			// その他のエラー
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 正常時のレスポンス
		c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
	})

	r.Run(":8081")
}
