package ballpark

import (
	"net/http"

	"github.com/rms1000watt/degeneres-test/data"
	"github.com/rms1000watt/degeneres-test/helpers"

	log "github.com/sirupsen/logrus"
)

func TicketHandlerGET(w http.ResponseWriter, r *http.Request) {
	log.Debug("Starting TicketHandlerGET")
	defer log.Debug("Finished TicketHandlerGET")

	ticketIn, err := data.GetTicketIn(r)
	if err != nil {
		log.Error(err)
		http.Error(w, helpers.ErrorJSON(err.Error()), http.StatusInternalServerError)
		return
	}
	log.Debug(ticketIn)

	ticketOut := data.TicketOut{}

	// Developer do stuff here

	helpers.WriteJSON(ticketOut, w)

}

func TicketHandlerPOST(w http.ResponseWriter, r *http.Request) {
	log.Debug("Starting TicketHandlerPOST")
	defer log.Debug("Finished TicketHandlerPOST")

	ticketIn, err := data.GetTicketIn(r)
	if err != nil {
		log.Error(err)
		http.Error(w, helpers.ErrorJSON(err.Error()), http.StatusInternalServerError)
		return
	}
	log.Debug(ticketIn)

	ticketOut := data.TicketOut{}

	// Developer do stuff here

	helpers.WriteJSON(ticketOut, w)

}

func TicketHandlerPUT(w http.ResponseWriter, r *http.Request) {
	log.Debug("Starting TicketHandlerPUT")
	defer log.Debug("Finished TicketHandlerPUT")

	ticketIn, err := data.GetTicketIn(r)
	if err != nil {
		log.Error(err)
		http.Error(w, helpers.ErrorJSON(err.Error()), http.StatusInternalServerError)
		return
	}
	log.Debug(ticketIn)

	ticketOut := data.TicketOut{}

	// Developer do stuff here

	helpers.WriteJSON(ticketOut, w)

}
