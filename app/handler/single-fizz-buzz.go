package handler

import (
	"net/http"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/kevinsudut/single-fizz-buzz/pkg/lib/log"
	"github.com/kevinsudut/single-fizz-buzz/pkg/lib/monitoring"
)

type monitoringHandleSingleFizzBuzzWithRange struct {
	Request struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"request"`
	Response string `json:"response"`
}

func (h handler) handleSingleFizzBuzzWithRange(w http.ResponseWriter, r *http.Request) {
	var (
		fromStr  = r.URL.Query().Get("from")
		toStr    = r.URL.Query().Get("to")
		response = ""

		monitoringName  = "HandleSingleFizzBuzzWithRange"
		monitoringStart = time.Now()
	)
	defer func() {
		monitoringContent, err := jsoniter.MarshalToString(monitoringHandleSingleFizzBuzzWithRange{
			Request: struct {
				From string "json:\"from\""
				To   string "json:\"to\""
			}{
				From: fromStr,
				To:   toStr,
			},
			Response: response,
		})
		if err != nil {
			log.Warnln("handleSingleFizzBuzzWithRange failed record monitoring content")
		}

		// Log its request, response and latency to STDOUT
		monitoring.RecordMonitoring(monitoringName, monitoringStart, monitoringContent)
	}()

	from, err := strconv.ParseInt(fromStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid parameter 'from' " + err.Error()))
		return
	}

	to, err := strconv.ParseInt(toStr, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid parameter 'to' " + err.Error()))
		return
	}

	response, err = h.usecase.UseCaseSingleFizzBuzzWithRange(r.Context(), from, to)
	if err != nil {
		log.Errorln("h.usecase.UseCaseSingleFizzBuzzWithRange", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}
