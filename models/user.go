package models

type User struct {
	User string `json:"name"`
	Birthday string `json:"birthday"`
	Address string `json:"address"`
	Nick string `json:"nick"`
	Password  string `json:"password"`
}

/*type User struct {
	User string`form:"name"`
	Birthday string `form:"birthday"`
	Address string `form:"address"`
	Nick string `form:"nick"`
	Password string `form:"password"`
}*/