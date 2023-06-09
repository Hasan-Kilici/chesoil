package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Hasan-Kilici/kawethradb"
	"time"
	"strconv"
	"chesoil/utils"
)

func UploadForumPOST(c*gin.Context){
	userToken, err := c.Cookie("Token")
	if err != nil {
		c.Redirect(http.StatusFound,"/")
		return
	}

	User,err := kawethradb.Find("./data/Users.csv", "Token", userToken)
	if err != nil {
		c.Redirect(http.StatusFound,"/")
		return
	}
	t := time.Now()
	forumID := kawethradb.Count("./data/Forums.csv")
	forumIDForNewForum := strconv.Itoa(forumID)
	forumName := c.PostForm("name")
	forumQuestion := c.PostForm("question")
	authorID := User["ID"]
	authorUrl := User["Url"]
	authorPP := User["ProfilePhoto"]
	AnswersCount := "0"
	Created_At := t.Format("2006-01-02 15:04:05")
	Updated_At := t.Format("2006-01-02 15:04:05")
	Upvotes := "0"
	Downvotes := "0"
	Status := "Opened"
	Token := utils.GenerateToken(22)
	newForumData := []string{forumIDForNewForum,Token,authorID,authorUrl,authorPP,forumName,forumQuestion,Created_At,Updated_At,AnswersCount,Upvotes,Downvotes,Status}
	kawethradb.Insert("./data/Users.csv", newForumData)
	url := "/forum/"+forumIDForNewForum
	c.Redirect(http.StatusFound, url)
}

func CreateCommentPOST(){
	fmt.Println("31")
}

func UpvoteCommentPOST(){
	fmt.Println("31")
}

func UpvoteForumPOST(){
	fmt.Println("31")
}

func DownvoteCommentPOST(){
	fmt.Println("31")
}

func DownvoteForumPOST(){
	fmt.Println("31")
}