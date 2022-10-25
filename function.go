package controller

import (
	"net/http"
	"sonarhook/config"
	"sonarhook/message"

	log "github.com/sirupsen/logrus"
)

var err error

func HandleWebhook(w http.ResponseWriter, r *http.Request) {


	cfg := config.NewConfig()

	var webhook *config.Webhook
	if webhook = getWebhook(cfg, r.URL.Path); webhook == nil {
		http.Error(w, "Webhook config not found.", http.StatusNotFound)
		return
	}

	// Message constructor
	mc, err := message.NewMessageConstructor(webhook, r.Body)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	
	// Send the message
	err = mc.SendMessage()
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func getWebhook(cfg *config.Config, path string) *config.Webhook {
	for _, u := range cfg.Webhooks {
		if u.Path == path {
			return &u
		}
	}
	return nil
}