package utils

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
	"math/rand"
	"time"
	"fmt"
	"github.com/spf13/viper"
	"crypto/sha256"
	"encoding/hex"
)

type Config struct {
	EmailSenderName      string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

func GenerateToken(length int) string {
	rand.Seed(time.Now().UnixNano())

	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func GenerateRandomPassword(length int) string {
	rand.Seed(time.Now().UnixNano())

	const letters = "0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func HashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CheckPassword(hashedPassword, password string) bool {
	return hashedPassword == HashPassword(password)
}

func ReplaceSpacesWithDash(text string) string {
    return strings.ReplaceAll(text, " ", "-")
}

func UnHashPassword(password string) (string, error) {
	hashBytes := []byte(password)
	err := bcrypt.CompareHashAndPassword(hashBytes, []byte(""))
	if err != nil {
		return "", err
	}
	return string(hashBytes), nil
}

func SetUrl(email string) string{
	url := strings.Split(email, "@");
	fmt.Println(url)
	return url[0]; 
}

func SetUrlForInfo(title string) string{
	url := strings.Split(title, "-");
	fmt.Println(url)
	return url[0]; 
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}