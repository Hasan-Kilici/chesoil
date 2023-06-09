package routes 

import(
	"net/http"
	"github.com/gin-gonic/gin"
    "github.com/Hasan-Kilici/kawethradb"
)

func UserDashboardPage(c *gin.Context){
	userToken, err := c.Cookie("Token")
	if err != nil {
		c.Redirect(http.StatusFound,"/")
		return
	}

	User, err:= kawethradb.Find("./data/Users.csv", "Token", userToken)
	if err != nil {
		c.Redirect(http.StatusFound,"/")
		return
	}

	c.HTML(http.StatusOK, "settings.tmpl", gin.H{
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