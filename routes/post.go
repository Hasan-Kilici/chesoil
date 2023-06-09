package routes

import (
	db "chesoil/db"
	discord "chesoil/discord"
	mail "chesoil/mail"
	"chesoil/utils"
	"fmt"
	"image"
	"net/http"
	"path/filepath"
	"time"

	"github.com/Hasan-Kilici/kawethradb"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID           int
	Token        string
	Email        string
	Password     string
	Name         string
	ProfilePhoto string
	Url          string
	Followers    int
	Following    int
	Perm         string
}

type Info struct {
	ID           int
	Token        string
	Name         string
	GaleryPhotos string
	Banner       string
	Property1    string
	Property2    string
	Property3    string
	Description  string
	Title        string
	Url          string
	Status       string
	Type         string
}

func UploadPOST(c *gin.Context) {
	clientIP := c.GetHeader("X-Real-Ip")
	if clientIP == "" {
		clientIP = c.GetHeader("X-Forwarded-For")
	}
	if clientIP == "" {
		clientIP = c.ClientIP()
	}
	infos, _ := kawethradb.FindAll("./data/Info.csv", "Status", "open") // JS Yavaş.
	userToken, err := c.Cookie("Token")
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	user, err := kawethradb.Find("./data/Users.csv", "Token", userToken)
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	fmt.Println(user["Perm"])
	if user["Perm"] == "User" {
		ExtraMessage := "\n" + "**UYARI, APTAL BI OROSPU EVLADI BILGI YUKLEMEYE ÇALIYOR** : @everyone"
		fmt.Println(discord.SendLogMessage(user["Name"], clientIP, user["Email"], "Bilgi Yükledi", ExtraMessage))
		c.Redirect(http.StatusFound, "/")
		return
	}

	file, err := c.FormFile("banner")
	if err != nil {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"redirect": "#/admin/dashboard",
			"error":    "Fotoğraf yüklenmemiş",
		})
		return
	}

	imageFile, err := file.Open()
	if err != nil {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"redirect": "#/admin/dashboard",
			"error":    "Dosya açılamadı",
			"infos":    infos,
		})
		return
	}
	defer imageFile.Close()

	img, _, err := image.Decode(imageFile)
	if err != nil {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"redirect": "#/admin/dashboard",
			"error":    "Dosya paketlenemedi",
			"infos":    infos,
		})
		return
	}

	ext := filepath.Ext(file.Filename)
	filename := file.Filename[:len(file.Filename)-len(ext)]
	bannerName := utils.ReplaceSpacesWithDash("banner_" + time.Now().Format("20060102150405") + "_" + filename + "_resized" + ext)

	if err := utils.SaveResizedImage(img, "./uploads/"+bannerName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"redirect": "#/admin/dashboard",
			"error":    "Eksik Dosya",
			"infos":    infos,
		})
		return
	}
	files := form.File["gallery[]"]

	var galleryPaths []string

	for _, file := range files {
		imageFile, err := file.Open()
		if err != nil {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"redirect": "#/admin/dashboard",
				"error":    "Fotoğraf Yüklenmemiş",
				"infos":    infos,
			})
			return
		}
		defer imageFile.Close()
		img, _, err := image.Decode(imageFile)
		if err != nil {
			c.HTML(http.StatusOK, "index.tmpl", gin.H{
				"redirect": "#/admin/dashboard",
				"error":    "Dosya paketlenemedi",
				"infos":    infos,
			})
			return
		}

		ext := filepath.Ext(file.Filename)
		filename := file.Filename[:len(file.Filename)-len(ext)]
		newFileName := utils.ReplaceSpacesWithDash("gallery_" + time.Now().Format("20060102150405") + "_" + filename + "_resized" + ext)
		if err := utils.SaveResizedImage(img, "./uploads/"+newFileName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		galleryPaths = append(galleryPaths, "/uploads/"+newFileName)
	}

	var galleryPhotoPaths string
	for i, path := range galleryPaths {
		if i == 0 {
			galleryPhotoPaths = path
		} else {
			galleryPhotoPaths += " > " + path
		}
	}

	title := c.PostForm("title")
	types := c.PostForm("types")
	description := c.PostForm("description")
	property1 := c.PostForm("property1")
	property2 := c.PostForm("property2")
	property3 := c.PostForm("property3")
	length := kawethradb.Count("./data/Info.csv") + 1
	token := utils.GenerateToken(36)
	bannerPath := "/uploads/" + bannerName
	data := Info{
		ID:           length,
		Token:        token,
		Name:         title,
		GaleryPhotos: galleryPhotoPaths,
		Banner:       bannerPath,
		Property1:    property1,
		Property2:    property2,
		Property3:    property3,
		Description:  description,
		Title:        title,
		Url:          utils.SetUrlForInfo(title),
		Status:       "open",
		Type:         types,
	}

	ExtraMessage := "\n" + "**Eklenen Bitki/Hayvan** : " + title + "\n" + "**Açıklama** : " + description + "\n"
	kawethradb.Insert("./data/Info.csv", data)
	fmt.Println(discord.SendLogMessage(user["Name"], clientIP, user["Email"], "Bilgi Yükledi", ExtraMessage))
	c.Redirect(http.StatusFound, "/#/admin/dashboard")
}

func RegisterPOST(c *gin.Context) {
	clientIP := c.GetHeader("X-Real-Ip")
	if clientIP == "" {
		clientIP = c.GetHeader("X-Forwarded-For")
	}
	if clientIP == "" {
		clientIP = c.ClientIP()
	}
	infos, _ := kawethradb.FindAll("./data/Info.csv", "Status", "open")
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm-password")
	username := c.PostForm("username")
	if password == confirmPassword {
		findUserByEmail, err := kawethradb.Find("./data/Users.csv", "Email", email)
		if err != nil {
			findUserByName, err := kawethradb.Find("./data/Users.csv", "Name", username)
			if err != nil {
				randomPassword := utils.GenerateRandomPassword(6)
				hashedPassword := utils.HashPassword(randomPassword)
				fmt.Println(randomPassword)
				token := db.RegisterUser(username, email, password)
				fmt.Println(email)
				mail.SendAuthMail(email, token, hashedPassword, randomPassword)
				link := "/accept/register/" + token + "/" + randomPassword
				ExtraMessage := "\n" + "**Kaydedilecek Kullanıcı adı** : " + username + "\n" + "**Kaydedilecek Email** : " + email + "\n **Kayıt durumu** : Başarılı, Doğrulama mesajı gönderildi"
				fmt.Println(discord.SendLogMessage("Anonim", clientIP, email, "Kayıt olmaya çalışıyor", ExtraMessage))
				c.Redirect(http.StatusFound, link)
				return
			}
			fmt.Println(findUserByName)
			return
		}
		ExtraMessage := "\n" + "**Kaydedilecek Kullanıcı adı** : " + username + "\n" + "**Kaydedilecek Email** : " + email + "\n **Kayıt durumu** : Başarısız, Bu Email Kullanılıyor"
		fmt.Println(discord.SendLogMessage("Anonim", clientIP, email, "Kayıt olmaya çalışıyor", ExtraMessage))
		fmt.Println(findUserByEmail)
		c.HTML(http.StatusOK, "register.tmpl", gin.H{
			"redirect": "#/register",
			"error":    "Bu Email Kullanılıyor.",
			"infos":    infos,
		})
	} else {
		c.HTML(http.StatusOK, "register.tmpl", gin.H{
			"redirect": "#/register",
			"error":    "Şifreleriniz uyuşmuyor.",
			"infos":    infos,
		})
		return
	}
}

func LoginPOST(c *gin.Context) {
	clientIP := c.GetHeader("X-Real-Ip")
	if clientIP == "" {
		clientIP = c.GetHeader("X-Forwarded-For")
	}
	if clientIP == "" {
		clientIP = c.ClientIP()
	}
	infos, _ := kawethradb.FindAll("./data/Info.csv", "Status", "open")
	email := c.PostForm("email")
	user, err := kawethradb.Find("./data/Users.csv", "Email", email)
	if err != nil {
		ExtraMessage := "\n" + "**Giriş yapılacak Email** : " + email + "\n **Giriş durumu** : Başarısız, Kullanıcı bulunamadı"
		fmt.Println(discord.SendLogMessage("Anonim", clientIP, email, "Giriş yapmaya çalışıyor", ExtraMessage))
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"redirect": "#/login",
			"error":    "Kullanıcı Bulunamadı",
		})
		return
	}
	fmt.Println(user["Password"])
	password := c.PostForm("password")
	hashedPassword := user["Password"]
	fmt.Println(utils.CheckPassword(hashedPassword, password))
	if utils.CheckPassword(hashedPassword, password) {
		c.SetCookie("Token", user["Token"], 36000, "/", "", false, true)
		ExtraMessage := "\n" + "**Giriş yapılacak Email** : " + email + "\n **Giriş durumu** : Başarılı"
		fmt.Println(discord.SendLogMessage("Anonim", clientIP, email, "Giriş yapmaya çalışıyor", ExtraMessage))
		c.Redirect(http.StatusFound, "/")
		return
	} else {
		ExtraMessage := "\n" + "**Giriş yapılacak Email** : " + email + "\n **Giriş durumu** : Başarısız, Şifre yanlış"
		fmt.Println(discord.SendLogMessage("Anonim", clientIP, email, "Giriş yapmaya çalışıyor", ExtraMessage))
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"redirect": "#/login",
			"error":    "Email ya da Şifre yanlış",
			"infos":    infos,
		})
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func AcceptRegisterPOST(c *gin.Context) {
	clientIP := c.GetHeader("X-Real-Ip")
	if clientIP == "" {
		clientIP = c.GetHeader("X-Forwarded-For")
	}
	if clientIP == "" {
		clientIP = c.ClientIP()
	}
	password := c.Param("password")
	passwordInput := c.PostForm("password")
	token := c.Param("userToken")
	fmt.Println(utils.CheckPassword(password, passwordInput))
	if password == passwordInput {
		fmt.Println("Şifre Doğru")
		user, _ := kawethradb.Find("./data/UnSignedUsers.csv", "Token", token)
		userEmail, _ := user["Email"]
		userPassword, _ := user["Password"]
		userName, _ := user["Name"]
		c.SetCookie("Token", db.CrateUser(userEmail, userPassword, userName), 36000, "/", "", false, true)
		ExtraMessage := "\n" + "**Hesabı onaylanan Email** : " + userEmail + "\n **Giriş durumu** : Başarılış"
		fmt.Println(discord.SendLogMessage("Anonim", clientIP, userEmail, "Kullanıcı Hesabı onaylandı", ExtraMessage))
		kawethradb.Delete("./data/UnsignedUsers.csv", "Token", token)
		c.Redirect(http.StatusFound, "/")
		return
	} else {
		c.Redirect(http.StatusFound, "/")
		return
	}
}

func DeleteInfoPOST(){

}

func EditInfoPOST(){

}