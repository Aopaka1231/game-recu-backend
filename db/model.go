// db package
package db

//Todoモデルを定義
//PrimaryKeyはID
type Rec struct {
	ID uint `gorm:"primarykey" json:"id"`
	App string `json:"app"`
	Time string `json:"time"`
}