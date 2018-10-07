package protocol

type ClientConnect struct {
	ID int `json:"id"`
	Username string `json:"username"`
}

type ClientEnterWorld struct {
	ID int `json:"id"`
	WorldID int `json:"world_id"`
}