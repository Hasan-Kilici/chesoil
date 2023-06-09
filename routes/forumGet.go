package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Hasan-Kilici/kawethradb"
)

func ForumHomePage(c * gin.Context){
	userToken, err := c.Cookie("Token")
	if err != nil {
		c.HTML(http.StatusOK, "forumhome.tmpl", gin.H{
			"userstatus":false,
		})
		return
	} 

	User, err := kawethradb.Find("./data/Users.csv", "Token", userToken)
	if err != nil {
		c.HTML(http.StatusOK, "forumhome.tmpl", gin.H{
			"userstatus":false,
		})
		return
	}

	c.HTML(http.StatusOK, "forumhome.tmpl", gin.H{
		"userstatus":true,
		"username":User["Name"],
		"email":User["Email"],
		"ID":User["ID"],
		"followers":User["Followers"],
		"following":User["Following"],
		"perm":User["Perm"],
		"profilePhoto":User["ProfilePhoto"],
	})
}

func ForumPage(c *gin.Context){
	forumID := c.Param("id")
	forum,err := kawethradb.Find("./data/Forums.csv", "ID", forumID)
	if err != nil {
		c.HTML(http.StatusOK, "hata.tmpl", gin.H{
			"title": "Bilgi bulunamadı",
			"error": "",
			"body": "",
			"id":forumID,
		})   
		return
	}

	userToken,err := c.Cookie("Token")
	if err != nil {
		c.HTML(http.StatusOK, "forum.tmpl", gin.H{
			"title":forum["Title"],
			"question":forum["Question"],
			"userstatus":false,
			"id":forumID,
		})
		return
	}

	User, err := kawethradb.Find("./data/Users.csv", "Token", userToken)
	if err != nil {
		c.HTML(http.StatusOK, "forum.tmpl", gin.H{
			"title":forum["Title"],
			"question":forum["Question"],
			"userstatus":false,
			"id":forumID,
		})
		return
	}

	c.HTML(http.StatusOK, "forum.tmpl", gin.H{
		"title":forum["Title"],
		"question":forum["Question"],
		"userstatus":true,
		"username":User["Name"],
		"email":User["Email"],
		"ID":User["ID"],
		"followers":User["Followers"],
		"following":User["Following"],
		"perm":User["Perm"],
		"profilePhoto":User["ProfilePhoto"],
		"id":forumID,
	})
}