package ballpark

import (
	"net/http"

	"github.com/rms1000watt/degeneres-test/data"
	"github.com/rms1000watt/degeneres-test/helpers"

	log "github.com/sirupsen/logrus"
)

func PersonHandlerGET(w http.ResponseWriter, r *http.Request) {
	log.Debug("Starting PersonHandlerGET")
	defer log.Debug("Finished PersonHandlerGET")

	personIn, err := data.GetPersonIn(r)
	if err != nil {
		log.Error(err)
		http.Error(w, helpers.ErrorJSON(err.Error()), http.StatusInternalServerError)
		return
	}
	log.Debug(personIn)

	personOut := data.PersonOut{}

	// Developer do stuff here

	helpers.WriteJSON(personOut, w)

}

func PersonHandlerPOST(w http.ResponseWriter, r *http.Request) {
	log.Debug("Starting PersonHandlerPOST")
	defer log.Debug("Finished PersonHandlerPOST")

	personIn, err := data.GetPersonIn(r)
	if err != nil {
		log.Error(err)
		http.Error(w, helpers.ErrorJSON(err.Error()), http.StatusInternalServerError)
		return
	}
	log.Debug(personIn)

	personOut := data.PersonOut{}

	// Developer do stuff here

	helpers.WriteJSON(personOut, w)

}
