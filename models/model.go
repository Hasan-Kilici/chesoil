package db

import (
	"fmt"
)

type User struct {
	ID           	int
	Token        	string
	Email        	string
	Password     	string
	Name         	string
	ProfilePhoto 	string
	Url          	string
	Followers    	int
	Following    	int
	Perm         	string
}

type Info struct {
	ID           int
	Token        string
	Name         int
	GaleryPhotos string
	Banner       string
	Property1    string
	Property2    string
	Property3    string
	Description  string
	Title        string
	Url          string
}

type Forums struct {
	ID           int
	Token        string
	AuthorID     int
	AuthorUrl	 string
	AuthorPP	 string
	Title		 string
	Question     string
	Created_At   string
	Updated_At   string
	AnswersCount string
	Upvote       int
	Downvote     int
	Status		 string
}

type Comment struct {
	ID 				int
	Token 			string
	AuthorID 		int
	AuthorName 		string
	AuthorPP 		string
	Comment 		string
	Upvote 			string
	Downvote 		string
	ForumID			int
}

func Models() {
	fmt.Println("....")
}