package ballpark

import (
	"net/http"

	"github.com/rms1000watt/degeneres-test/data"
	"github.com/rms1000watt/degeneres-test/helpers"

	log "github.com/sirupsen/logrus"
)

func ManagementHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("Starting ManagementHandler")
	defer log.Debug("Finished ManagementHandler")

	managementIn, err := data.GetManagementIn(r)
	if err != nil {
		log.Error(err)
		http.Error(w, helpers.ErrorJSON(err.Error()), http.StatusInternalServerError)
		return
	}
	log.Debug(managementIn)

	managementOut := data.ManagementOut{}

	// Developer do stuff here

	helpers.WriteJSON(managementOut, w)

}
