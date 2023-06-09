package discord

import (
	"github.com/gtuk/discordwebhook"
	"log"
)

func SendLogMessage(Username, IP, Email, Title, ExtraMessage string) string {
	var username = "Site Kayıtları"
	var content = "**"+Username+"** Kullanıcısı **"+Title+"**\n**Kullanıcının Email adresi** : "+Email+"\n**Kullanıcının IP Adresi** : "+IP+ExtraMessage
	var url = "yourWebhookApiURL"
 
	message := discordwebhook.Message{
		Username: &username,
		Content: &content,
	}
 
	err := discordwebhook.SendMessage(url, message)
	if err != nil {
		log.Fatal(err)
	}
	return content;
}