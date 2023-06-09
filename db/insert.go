package db

import (
	"github.com/Hasan-Kilici/kawethradb"
	models "chesoil/models"
	utils "chesoil/utils"
)




func RegisterUser(username, email, password string) string{
	token := utils.GenerateToken(36);
	user := models.User{
		ID:	kawethradb.Count("./data/UnsignedUsers.csv")+1,
		Token: token,
		Email: email,
		Password: password,
		Name: username,
		ProfilePhoto:"",
		Url: utils.SetUrl(email),
		Followers: 0,
		Following: 0,
		Perm: "User",	
	}

	kawethradb.Insert("./data/UnsignedUsers.csv", user)
	return token;
}

func CrateUser(email, password, username string) string{
	token := utils.GenerateToken(36);
	user := models.User{
		ID:	kawethradb.Count("./data/Users.csv")+1,
		Token: token,
		Email: email,
		Password: utils.HashPassword(password),
		Name: username,
		ProfilePhoto:"",
		Url: utils.SetUrl(email),
		Followers: 0,
		Following: 0,
		Perm: "User",	
	}
	kawethradb.Insert("./data/Users.csv", user)
	return token;
}