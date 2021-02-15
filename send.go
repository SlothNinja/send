package send

import (
	"context"
	"os"

	"github.com/SlothNinja/log"
	"github.com/SlothNinja/sn"
	"github.com/mailjet/mailjet-apiv3-go"
)

func getKeys() (string, string) {
	return os.Getenv("MJ_API_KEY_PUB"), os.Getenv("MJ_API_KEY_PRIV")
}

func Messages(c context.Context, msgInfo ...mailjet.InfoMessagesV31) (*mailjet.ResultsV31, error) {
	log.Debugf("Entering")
	defer log.Debugf("Entering")

	pub, priv := getKeys()
	mailjetClient := mailjet.NewMailjetClient(pub, priv)
	msgs := mailjet.MessagesV31{Info: msgInfo}
	if sn.IsProduction() {
		return mailjetClient.SendMailV31(&msgs)
	}
	for _, msg := range msgInfo {
		log.Debugf("sent message: %#v", msg)
	}
	return nil, nil
}
