package handler

import (
	"encoding/json"
	"github.com/SawitProRecruitment/UserService/service"
	"github.com/google/uuid"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) CreateEstate(ec echo.Context) error {
	var payload service.PayloadCreateEstate

	if err := json.NewDecoder(ec.Request().Body).Decode(&payload); err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}

	res, err := s.Service.CreateEstate(ec.Request().Context(), payload)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, err)
	}

	return ec.JSON(http.StatusCreated, map[string]interface{}{
		"id": res,
	})
}

func (s *Server) CreateTree(ec echo.Context) error {
	var payload service.PayloadCreateTree

	if err := json.NewDecoder(ec.Request().Body).Decode(&payload); err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}

	id := ec.Param("id")
	payload.EstateId = uuid.MustParse(id)

	res, err := s.Service.CreateTree(ec.Request().Context(), payload)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, err)
	}

	return ec.JSON(http.StatusCreated, map[string]interface{}{
		"id": res,
	})
}
