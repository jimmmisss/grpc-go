package server

type Server struct {
}

type user struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	AvatarURL  string `json:"avatarURL"`
	Location   string `json:"location"`
	Followers  int64  `json:"followers"`
	Following  int64  `json:"following"`
	Repos      string `json:"repos"`
	Gists      string `json:"gists"`
	URL        string `json:"URL"`
	StarredURL string `json:"starredURL"`
	ReposURL   string `json:"reposURL"`
}
