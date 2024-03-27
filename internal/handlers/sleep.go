package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pidanou/lifetrack/internal/datastore"
	"github.com/pidanou/lifetrack/internal/types"
	// "github.com/pidanou/lifetrack/pkg/timeutil"
)

type SleepHandler struct {
	Datastore datastore.Datastore
}

func (h *SleepHandler) HandlePostSleep(c echo.Context) error {

	type typeReqBody struct {
		Quantity float64
	}

	var reqBody typeReqBody

	err := c.Bind(&reqBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	sleep := types.Sleep{Quantity: reqBody.Quantity}

	h.Datastore.CreateSleep(&sleep)

	return c.String(http.StatusOK, "OK")
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
