package send

import (
	"io/ioutil"
	"net/http"

	"github.com/SlothNinja/codec"
	"github.com/SlothNinja/log"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/mail"
)

func Mail(c *gin.Context) {
	log.Debugf("Entering")
	defer log.Debugf("Exiting")

	encoded, err := ioutil.ReadAll(c.Request.Body)
	c.Request.Body.Close()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	m := new(mail.Message)
	err = codec.Decode(m, encoded)
	if err != nil {
		log.Errorf(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = mail.Send(appengine.NewContext(c.Request), m)
	if err != nil {
		log.Errorf(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusOK)
}
