package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine){
	r.Use(cors.New(cors.Config{
		//アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"*",
		},
		//アクセスを許可したいhttpメソッド
		AllowMethods: []string{
			"POST",
			"GET",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		//許可したいhttpリクエストヘッダ
		AllowHeaders: []string{
			"*",
		},
	}))
}