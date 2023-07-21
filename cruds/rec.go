package cruds

import (
	"game-rec-back/db"
)

//データベースに保存する関数
func CreateRec(user string, app string, time string)(res_rec db.Rec, err error){
	res_rec = db.Rec{User: user, App: app, Time: time}
	err = db.Ssql.Create(&res_rec).Error
	return
}

//すべてのRecの配列を返す関数
func GetAllRec()(res_rec []db.Rec, err error){
	err = db.Ssql.Find(&res_rec).Error
	return
}

//引数のidのRecを削除する関数
func DeleteRed(id uint)(err error){
	err = db.Ssql.Where("id = ?", id).Delete(&db.Rec{}).Error

	return
}