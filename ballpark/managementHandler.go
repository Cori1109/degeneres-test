package ballpark

import (
	"net/http"

	"github.com/Cori1109/degeneres-test/data"
	"github.com/Cori1109/degeneres-test/helpers"

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
