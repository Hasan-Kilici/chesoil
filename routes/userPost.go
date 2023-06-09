package routes 

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
    "github.com/Hasan-Kilici/kawethradb"
	"chesoil/utils"
	"chesoil/discord"
)

func ChangePasswordPOST(c*gin.Context){
	clientIP := c.GetHeader("X-Real-Ip")
	if clientIP == "" {
		clientIP = c.GetHeader("X-Forwarded-For")
	}
	if clientIP == "" {
		clientIP = c.ClientIP()
	}
	userToken,err := c.Cookie("Token")
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	User,err := kawethradb.Find("./data/Users.csv", "Token", userToken)
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	UserPassword := User["Password"]
	InputPassword := c.PostForm("password")
	NewPasswordInput := c.PostForm("newPassword")
	if utils.CheckPassword(UserPassword, InputPassword){
		NewPassword := utils.HashPassword(NewPasswordInput)
		newUserValues := []string{User["ID"],User["Token"],User["Email"],NewPassword,User["ProfilePhoto"],User["Name"],User["Url"],User["Followers"],User["Following"],User["Perm"]}

		err := kawethradb.Update("./data/Users.csv", "ID", User["ID"], newUserValues)
		if err != nil {
			fmt.Println("Kayıt güncellenirken bir hata oluştu:", err)
			return
		}

		c.HTML(http.StatusOK, "settings.tmpl", gin.H{
			"message": "Şifreniz başarıyla değiştirildi!",
			"Username":User["Name"],
			"Email":User["Email"],
			"ID":User["ID"],
			"Followers":User["Followers"],
			"Following":User["Following"],
			"Perm":User["Perm"],
			"ProfilePhoto":User["ProfilePhoto"],
			"redirect":"#",
		})
		ExtraMessage := "\n" + "**Şifresini değiştiren hesap** : " + User["Email"] + "\n **Yeni şifre** : "+NewPassword
		fmt.Println(discord.SendLogMessage("Anonim", clientIP, User["Email"], "Kullanıcı Şifresini değiştirdi", ExtraMessage))
		return
	}

	c.HTML(http.StatusOK, "settings.tmpl", gin.H{
		"error": "Eski şifreniz Doğru değil!",
		"Username":User["Name"],
		"Email":User["Email"],
		"ID":User["ID"],
		"Followers":User["Followers"],
		"Following":User["Following"],
		"Perm":User["Perm"],
		"ProfilePhoto":User["ProfilePhoto"],
		"redirect":"#",
	})
}

func DeleteAccountPOST(c*gin.Context){
	userToken,err := c.Cookie("Token")
	if err != nil {
		c.Redirect(http.StatusFound,"/")
		return
	}

	User,err := kawethradb.Find("./data/Users.csv", "Token", userToken)
	if err != nil {
		c.Redirect(http.StatusFound,"/")
		return
	}

	kawethradb.Delete("./data/Users.csv", "Token", User["Token"])
	c.SetCookie("Token", "", -1, "/", "", false, true) 
	c.Redirect(http.StatusFound, "/")
}

func ArchiveAccountPOST(c*gin.Context){
	userToken,err := c.Cookie("Token")
	if err != nil {
		c.Redirect(http.StatusFound,"/")
		return
	}

	User,err := kawethradb.Find("./data/Users.csv", "Token", userToken)
	if err != nil {
		c.Redirect(http.StatusFound,"/")
		return
	}

	ArchiveUser := []string{User["ID"],User["Token"],User["Email"],User["Password"],User["Name"],User["ProfilePhoto"],User["Url"],User["Followers"],User["Following"],User["Perm"]}

	kawethradb.Insert("./data/ArchivedUsers.csv", ArchiveUser)
	kawethradb.Delete("./data/Users.csv", "Token", userToken)
	c.SetCookie("Token", "", -1, "/", "", false, true) 
	c.Redirect(http.StatusFound, "/")
}

func SignOutGET(c*gin.Context){
	clientIP := c.GetHeader("X-Real-Ip")
	if clientIP == "" {
		clientIP = c.GetHeader("X-Forwarded-For")
	}
	if clientIP == "" {
		clientIP = c.ClientIP()
	}
	userToken,err := c.Cookie("Token")
	if err != nil {
		c.Redirect(http.StatusFound,"/")
		return
	}
	ExtraMessage := "\n" + "**Çıkış yapan Hesabın Tokeni** : ||" + userToken + " ||\n"
	fmt.Println(discord.SendLogMessage("Anonim", clientIP, "Email'lik bir şey yok.", "Kullanıcı Çıkış yaptı", ExtraMessage))
	c.SetCookie("Token", "", -1, "/", "", false, true) 
	c.Redirect(http.StatusFound, "/")
}

func ChangeUsernamePOST(c*gin.Context){
	userToken, err := c.Cookie("Token")
	if err != nil {
		c.Redirect(http.StatusFound,"/")
		return
	}

	User, err := kawethradb.Find("./data/Users.csv", "Token", userToken)
	if err != nil {
		c.Redirect(http.StatusFound,"/")
		return
	}

	NewUsername := c.PostForm("username")
	NewUserData := []string{User["ID"],User["Token"],User["Email"],User["Password"],NewUsername,User["ProfilePhoto"],User["Url"],User["Followers"],User["Following"],User["Perm"]}
	kawethradb.Update("./data/Users.csv", "Token", userToken, NewUserData)

	c.HTML(http.StatusOK, "settings.tmpl", gin.H{
		"message": "Kullanıcı adı değiştirildi",
		"Username":User["Name"],
		"Email":User["Email"],
		"ID":User["ID"],
		"Followers":User["Followers"],
		"Following":User["Following"],
		"Perm":User["Perm"],
		"ProfilePhoto":User["ProfilePhoto"],
		"redirect":"#",
	})
}

func ChangeStatusPOST(){
	fmt.Println("Zort")
}

func SetStatusState(){
	fmt.Println("Zort")	
}