package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
    "github.com/Hasan-Kilici/kawethradb"
    "strings"
)


    func RegisterPage(c *gin.Context){
        userToken, err := c.Cookie("Token")
        if err != nil {
            c.HTML(http.StatusOK, "register.tmpl", gin.H{
                "title":"Kayıt ol",
            })
            return
        }
        user, err := kawethradb.Find("./data/Users.csv", "Token", userToken)
        if err != nil {
            fmt.Println(user)
            c.Redirect(http.StatusFound, "/")
            return
        } else {
            c.HTML(http.StatusOK, "register.tmpl", gin.H{
                "title":"Kayıt ol",
            })
            return
        }
    }

	func HomePage(ctx *gin.Context) {
        infos, _ := kawethradb.FindAll("./data/Info.csv","Status", "open")
        userToken, err := ctx.Cookie("Token")
        if err != nil {
            ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
                "title": "Anasayfa",
                "infos": infos,
            })
            return
        }
        user , err := kawethradb.Find("./data/Users.csv", "Token", userToken)
        if err != nil {
            ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
                "title": "Anasayfa",
                "infos": infos,
            })
            return
        }

        fmt.Println(user["Name"],user["Email"],user["Url"],user["Perm"])
        username := user["Name"]
        email := user["Email"]
        url := user["Url"]
        perm := user["Perm"]
        fmt.Println(infos)
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Anasayfa",
            "username": username,
            "email": email,
            "url": url,
            "perm": perm,
            "infos": infos,
            "userstatus": true,
		})
	}

    func InfoPage(c *gin.Context){
        id := c.Param("id")

        info, err := kawethradb.Find("./data/Info.csv","ID",id)
        fmt.Println(info)
        if err != nil {
            c.HTML(http.StatusOK, "hata.tmpl", gin.H{
                "title": "Bilgi bulunamadı",
                "error": "",
                "body": "",
                "id":id,
            })   
            return
        }
            title := info["Title"]
            galery := strings.Split(info["GaleryPhotos"], " > ")
            property1 := info["Property1"]
            property2 := info["Property2"]
            property3 := info["Property3"]
            banner := info["Banner"]
            description := info["Description"]
            types := info["Type"]

            userToken, err := c.Cookie("Token")
            if err != nil {
                c.HTML(http.StatusOK, "info.tmpl", gin.H{
                    "title": title,
                    "galery": galery,
                    "property1": property1,
                    "property2": property2,
                    "property3": property3,
                    "banner": banner,
                    "description": description,
                    "types": types,
                    "id":id,
                })
                return
        }

        user , err := kawethradb.Find("./data/Users.csv","Token", userToken)
        if err != nil {
            c.HTML(http.StatusOK, "info.tmpl", gin.H{
                "title": title,
                "galery": galery,
                "property1": property1,
                "property2": property2,
                "property3": property3,
                "banner": banner,
                "description": description,
                "types": types,
                "id":id,
            })
            return
        }

        username := user["Name"]
        email := user["Email"]
        url := user["Url"]
        perm := user["Perm"]

        c.HTML(http.StatusOK, "info.tmpl", gin.H{
            "title": title,
            "galery": galery,
            "property1": property1,
            "property2": property2,
            "property3": property3,
            "banner": banner,
            "description": description,
            "username": username,
            "email": email,
            "url": url,
            "perm": perm,
            "userstatus": true,
            "types": types,
            "id":id,
        })
    }


    func AcceptRegisterPage(c *gin.Context){
        hashedPassword := c.Param("hashedPassword")
        token := c.Param("userToken")
        c.HTML(http.StatusOK, "controlEmail.tmpl", gin.H{
            "title":"Emailinize Gönderdiğimiz Kodu Doğrulayın",
            "hash":hashedPassword,
            "token":token,
        })
    }

    func ProfilePage(c *gin.Context){
        purl := c.Param("url")
        profile , err := kawethradb.Find("./data/Users.csv","Url",purl)
        if err != nil {
            c.HTML(http.StatusOK, "hata.tmpl", gin.H{
                "title": "Kullanıcı bulunamadı",
                "error": "",
                "body": "",
            })   
            return
        }
        profileName := profile["Name"]
        profileFollowers := profile["Followers"]
        profileFollowing := profile["Following"]
        profilePhoto := profile["ProfilePhoto"]
        profilePerm := profile["Perm"]
        title := profileName+" Kullanıcısının Profili"
        userToken, err := c.Cookie("Token")

        if err != nil {
            c.HTML(http.StatusOK, "userProfile.tmpl", gin.H{
                "title": title,
                "profileName": profileName,
                "profileFollowers": profileFollowers,
                "profileFollowing": profileFollowing,
                "profilePhoto": profilePhoto,
                "profilePerm": profilePerm,
            })
            return
        }
        user , err := kawethradb.Find("./data/Users.csv", "Token", userToken)
        if err != nil {
            c.HTML(http.StatusOK, "userProfile.tmpl", gin.H{
                "title": title,
                "profileName": profileName,
                "profileFollowers": profileFollowers,
                "profileFollowing": profileFollowing,
                "profilePhoto": profilePhoto,
                "profilePerm": profilePerm,
            })
            return
        }

        username := user["Name"]
        email := user["Email"]
        url := user["Url"]
        perm := user["Perm"]

        c.HTML(http.StatusOK, "userProfile.tmpl", gin.H{
            "title": title,
            "username": username,
            "email": email,
            "url": url,
            "perm": perm,
            "userstatus": true,
            "profileName": profileName,
            "profileFollowers": profileFollowers,
            "profileFollowing": profileFollowing,
            "profilePhoto": profilePhoto,
            "profilePerm": profilePerm,
        })
    }

    func SavedInfosPage(c*gin.Context){
        userToken, err := c.Cookie("Token")
        if err != nil {
            c.HTML(http.StatusOK, "saved.tmpl", gin.H{
                "userstatus": false,
                "title": "Kayedilenler",
            })
            return
        }

        User, err := kawethradb.Find("./data/Users.csv", "Token", userToken)
        if err != nil {
            c.HTML(http.StatusOK, "saved.tmpl", gin.H{
                "userstatus": false,
                "title": "Kayedilenler",
            })
            return
        }

        username := User["Name"]
        email := User["Email"]
        url := User["Url"]
        perm := User["Perm"]

        c.HTML(http.StatusOK, "saved.tmpl", gin.H{
            "userstatus": true,
            "username": username,
            "email": email,
            "url": url,
            "perm": perm,
        })
    }