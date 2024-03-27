package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pidanou/lifetrack/internal/datastore"
	"github.com/pidanou/lifetrack/internal/types"
	"github.com/pidanou/lifetrack/pkg/timeutil"
)

type HydrationHandler struct {
	Datastore datastore.Datastore
}

func (h *HydrationHandler) HandlePostHydration(c echo.Context) error {

	type typeReqBody struct {
		Name     string
		Quantity float64
	}

	var reqBody typeReqBody

	err := c.Bind(&reqBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	fmt.Println(reqBody.Name, reqBody.Quantity)

	water := types.Hydration{Quantity: reqBody.Quantity, Name: reqBody.Name}

	error := h.Datastore.CreateHydration(&water)

	fmt.Println(error)

	return c.String(http.StatusOK, "OK")
}

func (h *HydrationHandler) HandleGetHydration(c echo.Context) error {

	type typeReqBody struct {
		Date string
	}

	var reqBody typeReqBody

	err := c.Bind(&reqBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	date, _ := timeutil.ParseDayDateString(reqBody.Date)

	hydration, _ := h.Datastore.GetHydrationByDay(date)

	fmt.Println(hydration)

	return c.String(http.StatusOK, "OK")

}
