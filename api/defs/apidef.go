package defs


type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}

type VideoInfo struct {
	Id string `json:"id"`
	AuthorId int `json:"author_id"`
	Name string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}