package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

//User represents a person using the web app
type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nick      string    `json:"nick"`
	CreatedOn time.Time `json:"createdOn"`
	Followers []User    `json:"followers"`
	Following []User    `json:"following"`
	Posts     []Post    `json:"posts"`
}

//SearchCompleteUser make 4 requests in the API to build the user
func SearchCompleteUser(userId uint64, r *http.Request) (User, error) {
	userChannel := make(chan User)
	followersChannel := make(chan []User)
	followingChannel := make(chan []User)
	postsChannel := make(chan []Post)

	go SearchUserData(userChannel, userId, r)
	go SearchFollowers(followersChannel, userId, r)
	go SearchFollowing(followingChannel, userId, r)
	go SearchPosts(postsChannel, userId, r)

	var (
		user      User
		followers []User
		following []User
		posts     []Post
	)

	for i := 0; i < 4; i++ {
		select {
		case userLoaded := <-userChannel:
			if userLoaded.ID == 0 {
				return User{}, errors.New("ERROR while searching the user")
			}
			user = userLoaded
		case followersLoaded := <-followersChannel:
			if followersLoaded == nil {
				followers = []User{}
			} else {
				followers = followersLoaded
			}
		case followingLoaded := <-followingChannel:
			if followingLoaded == nil {
				following = []User{}
			} else {
				following = followingLoaded
			}
		case postsLoaded := <-postsChannel:
			if postsLoaded == nil {
				posts = []Post{}
			} else {
				posts = postsLoaded
			}
		}
	}

	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return user, nil
}

//SearchUserData calls the API to get the base data of a user
func SearchUserData(chnnl chan<- User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userId)
	response, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		chnnl <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		chnnl <- User{}
		return
	}

	chnnl <- user
}

//SearchFollowers calls the API to get a user's followers
func SearchFollowers(chnnl chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userId)
	response, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		chnnl <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if err = json.NewDecoder(response.Body).Decode(&followers); err != nil {
		chnnl <- nil
		return
	}

	chnnl <- followers
}

func SearchFollowing(chnnl chan<- []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.APIURL, userId)
	response, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		chnnl <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if err = json.NewDecoder(response.Body).Decode(&following); err != nil {
		chnnl <- nil
		return
	}

	chnnl <- following
}

func SearchPosts(chnnl chan<- []Post, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/posts", config.APIURL, userId)
	response, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		chnnl <- nil
		return
	}
	defer response.Body.Close()

	var posts []Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		chnnl <- nil
		return
	}

	chnnl <- posts
}
