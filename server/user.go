package server

import (
	"encoding/json"
	"fmt"
	user "go-grpc/proto"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/context"
)

type Server struct {
}

type User struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	AvatarURL  string `json:"avatarURL"`
	Location   string `json:"location"`
	Followers  int64  `json:"followers"`
	Following  int64  `json:"following"`
	Repos      int64  `json:"repos"`
	Gists      int64  `json:"gists"`
	URL        string `json:"URL"`
	StarredURL string `json:"starredURL"`
	ReposURL   string `json:"reposURL"`
}

func (s *Server) GetUser(ctx context.Context, in *user.UserRequest) (*user.UserResponse, error) {
	log.Printf("Receive message from client: %s", in.Username)

	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%v", in.Username))
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	usr := User{}
	jsonErr := json.Unmarshal(body, &usr)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return &user.UserResponse{
		Id:        usr.ID,
		Name:      usr.Name,
		Username:  usr.Username,
		Avatarurl: usr.AvatarURL,
		Location:  usr.Location,
		Statistics: &user.Statistics{
			Followers: usr.Followers,
			Following: usr.Following,
			Repos:     usr.Repos,
			Gists:     usr.Gists,
		},
		ListURLs: []string{usr.URL, usr.StarredURL, usr.ReposURL},
	}, nil
}
