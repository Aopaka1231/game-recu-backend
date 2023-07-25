// db package
package db

//Todoモデルを定義
//PrimaryKeyはID
type Rec struct {
	Id uint `gorm:"primarykey" json:"id"`
	User string `json:"user"`
	App string `json:"app"`
	Time string `json:"time"`
}