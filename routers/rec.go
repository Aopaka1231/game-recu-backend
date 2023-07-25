package routers

import (
	"game-rec-back/cruds"
	"game-rec-back/db"
	"game-rec-back/schema"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Routerを初期化する関数
func InitRecRouters(tr *gin.RouterGroup){
	//""というパスに対してGETメソッドで、getRecを設定
	tr.GET("", getRec)
	//これも同じ
	tr.POST("", postRec)
	//"/:id"とすることで任意の文字列をidが拾ってこれるようにする
	tr.DELETE("/:id", deleteRec)
}

//データベースからrecの配列を受け取りそれをフロントに返す関数
func getRec(c *gin.Context){
	var (
		recs []db.Rec
		err error
	)

	//crudsで定義したGetAllRecを使用
	if recs, err = cruds.GetAllRec(); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	//errorがない場合にrecsを返す
	c.JSON(http.StatusOK, &recs)
}

func postRec(c *gin.Context){
	var(
		err error
		payload schema.CreateRec
		rec db.Rec
	)

	//BindJSONは構造体とjsonを結び付けてくれる
	if err = c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	//errorがなければ登録したrecを返す
	c.JSON(http.StatusOK, &rec)
}

func deleteRec(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	//crudsで定義したDeleteRecを実行
	//引数にはuintでキャストしたidを渡す
	err = cruds.DeleteRec(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	//errorがない場合、{"message": "OK!!"}を返す
	c.JSON(http.StatusOK, gin.H{
		"message": "OK!!",
		})
}