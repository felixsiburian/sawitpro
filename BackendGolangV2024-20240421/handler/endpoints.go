package handler

import (
	"encoding/json"
	"fmt"
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

func (s *Server) UpdateTree(ec echo.Context) error {
	var payload service.PayloadUpdateTree

	if err := json.NewDecoder(ec.Request().Body).Decode(&payload); err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}

	id := ec.Param("id")
	payload.Id = uuid.MustParse(id)

	res, err := s.Service.UpdateTree(ec.Request().Context(), payload)
	fmt.Println("err: ", err)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, err)
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"id": res,
	})
}

func (s *Server) TreeStatsByEstateId(ec echo.Context) error {
	id := ec.Param("id")
	uuidId := uuid.MustParse(id)

	res, err := s.Service.TreeStatsByEstateId(ec.Request().Context(), uuidId)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, err)
	}

	return ec.JSON(http.StatusOK, res)
}

func (s Server) Distance(ec echo.Context) error {
	estateId := ec.Param("id")
	uuidEstateId := uuid.MustParse(estateId)

	res, err := s.Service.DroneDistance(ec.Request().Context(), uuidEstateId)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, err)
	}

	return ec.JSON(http.StatusOK, map[string]interface{}{
		"distance": res,
	})
}
