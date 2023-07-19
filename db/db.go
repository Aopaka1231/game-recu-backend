package db

//gormとdriverのインポート
import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Ssql *gorm.DB
)

//データベースを初期化する関数の定義
func InitDB() (err error){
	//ここでsqliteを開く
	Ssql, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//エラーハンドリング
	if err != nil {
		return
	}
	//さっき作ったRecモデルでテーブルを作成する
	if err = Ssql.AutoMigrate(&Rec{}); err != nil {
		return
	}

	return
}