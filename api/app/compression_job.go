package app

import (
	"encoding/json"
	"net/http"

	"github.com/yeslayla/trythis.zip/api/api"
	"github.com/yeslayla/trythis.zip/api/core"
)

const (
	JobStatusUnknown    = "Unknown"
	JobStatusTriggered  = "Triggered"
	JobStatusInProgress = "In Progress"
	JobStatusComplete   = "Complete"
)

func (app *App) TriggerCompressionJob(w http.ResponseWriter, r *http.Request, params api.TriggerCompressionJobParams) {

	job := api.CompressionJob{
		Id:     core.GenerateID(),
		Status: JobStatusTriggered,
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(job)
	_, _ = w.Write(data)
}

func (app *App) GetCompressionJob(w http.ResponseWriter, r *http.Request, jobId string) {

	job := api.CompressionJob{
		Id:     jobId,
		Status: JobStatusUnknown,
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(job)
	_, _ = w.Write(data)
}
