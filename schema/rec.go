package schema

type CreateRec struct{
	Id uint64 `json:"id"`
	User string `json:"user"`
	App string `json:"app"`
	Time string `json:"time"`
}
