package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/mereiamangeldin/transaction/model"
	"log"
	"net/http"
)

func (h *Manager) ServeTransaction(c echo.Context) error {
	log.Println("inside")
	var transaction model.Transaction
	if err := c.Bind(&transaction); err != nil {
		log.Fatalln(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if transaction.TransactionType != "rent" && transaction.TransactionType != "purchase" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("invalid transaction type"))
	}
	err := h.srv.Transaction.ServeTransaction(transaction)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, model.SuccessResponse{Message: "Transaction is successfully completed"})
}

func (h *Manager) GetTransaction(c echo.Context) error {
	transactions, err := h.srv.Transaction.GetTransaction()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, transactions)
}

func (h *Manager) GetStatistics(c echo.Context) error {
	statistics, err := h.srv.Transaction.GetStatistics()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, statistics)
}
