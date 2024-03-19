package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pidanou/lifetrack/internal/datastore"
	"github.com/pidanou/lifetrack/internal/types"
	// "github.com/pidanou/lifetrack/pkg/timeutil"
)

type SleepHandler struct {
	Datastore datastore.Datastore
}

func (h *SleepHandler) HandlePostSleep(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)

	type typeReqBody struct {
		Quantity float64
	}

	var reqBody typeReqBody

	json.Unmarshal(body, &reqBody)

	fmt.Println(reqBody.Quantity)

	sleep := types.Sleep{Quantity: reqBody.Quantity}

	error := h.Datastore.CreateSleep(&sleep)

	fmt.Println(error)

}

func (h *SleepHandler) HandleGetSleep(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	type typeReqBody struct {
		Date string
	}

	var reqBody typeReqBody

	json.Unmarshal(body, &reqBody)

	// date, _ := timeutil.ParseDayDateString(reqBody.Date)

	// hydration, _ := h.Datastore.GetSleepByDay(date)

	// fmt.Println(hydration)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(reqBody.Date))

}
