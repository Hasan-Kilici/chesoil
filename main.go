package main

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
    routes "chesoil/routes"
)

func main(){

	store := persistence.NewInMemoryStore(time.Second)

	r := gin.Default()
	r.LoadHTMLGlob("src/*.tmpl")
	r.Static("/static", "./static/")
	r.Static("/uploads","./uploads")
	r.Static("/components", "./components/")

	config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://localhost:5000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
    config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}

	r.GET("/", cache.CachePage(store, time.Second*5,routes.HomePage)) 
	r.GET("/register", cache.CachePage(store, time.Second*5,routes.RegisterPage))
	r.GET("/forum", cache.CachePage(store, time.Second*5,routes.ForumHomePage))
    r.GET("/info/:id", cache.CachePage(store, time.Second*5,routes.InfoPage))
    r.GET("/accept/register/:userToken/:hashedPassword", routes.AcceptRegisterPage)
    r.GET("/profile/:url", cache.CachePage(store, time.Second*5,routes.ProfilePage))
	r.GET("/settings", cache.CachePage(store, time.Second*5,routes.UserDashboardPage))
	r.GET("/saved/infos", cache.CachePage(store, time.Second*5,routes.SavedInfosPage))
	r.GET("/forum/:id", cache.CachePage(store, time.Second*5,routes.ForumPage))

	r.POST("/upload", routes.UploadPOST)
    r.POST("/register", routes.RegisterPOST)
    r.POST("/login", routes.LoginPOST)
    r.POST("/accept/register/:userToken/:password", routes.AcceptRegisterPOST)
	r.POST("/change/password", routes.ChangePasswordPOST)
	r.POST("/delete/account", routes.DeleteAccountPOST)
	r.POST("/archive/account", routes.ArchiveAccountPOST)
	r.POST("/change/username", routes.ChangeUsernamePOST)

	r.GET("/signout", routes.SignOutGET)
	
	r.Run(":5000")
}