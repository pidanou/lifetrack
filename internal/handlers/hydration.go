package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pidanou/lifetrack/internal/datastore"
	"github.com/pidanou/lifetrack/internal/types"
	"github.com/pidanou/lifetrack/pkg/timeutil"
)

type HydrationHandler struct {
	Datastore datastore.Datastore
}

func (h *HydrationHandler) HandlePostHydration(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)

	type typeReqBody struct {
		Name     string
		Quantity float64
	}

	var reqBody typeReqBody

	json.Unmarshal(body, &reqBody)

	fmt.Println(reqBody.Name, reqBody.Quantity)

	water := types.Hydration{Quantity: reqBody.Quantity, Name: reqBody.Name}

	error := h.Datastore.CreateHydration(&water)

	fmt.Println(error)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(reqBody.Name))

}

func (h *HydrationHandler) HandleGetHydration(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)

	type typeReqBody struct {
		Date string
	}

	var reqBody typeReqBody

	json.Unmarshal(body, &reqBody)

	date, _ := timeutil.ParseDayDateString(reqBody.Date)

	hydration, _ := h.Datastore.GetHydrationByDay(date)

	fmt.Println(hydration)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(reqBody.Date))

}
