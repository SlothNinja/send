package send

import (
	"os"

	"github.com/SlothNinja/log"
	"github.com/gin-gonic/gin"
	"github.com/mailjet/mailjet-apiv3-go"
)

func getKeys() (string, string) {
	return os.Getenv("MJ_API_KEY_PUB"), os.Getenv("MJ_API_KEY_PRIV")
}

func Messages(c *gin.Context, msgInfo ...mailjet.InfoMessagesV31) (*mailjet.ResultsV31, error) {
	log.Debugf("Entering")
	defer log.Debugf("Entering")

	pub, priv := getKeys()
	mailjetClient := mailjet.NewMailjetClient(pub, priv)
	msgs := mailjet.MessagesV31{Info: msgInfo}
	return mailjetClient.SendMailV31(&msgs)
}
