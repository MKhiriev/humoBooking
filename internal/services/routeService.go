package services

import (
	"go-booking-system/internal/database"
	"go-booking-system/internal/models"
)

type RouteService struct {
	repository database.RouteRepository
}

func NewRouteService(repository database.RouteRepository) *RouteService {
	return &RouteService{repository: repository}
}

func (r *RouteService) Create(route models.Route) (models.Route, error) {
	return r.repository.Create(route)
}

func (r *RouteService) GetAll() []models.Route {
	return r.repository.GetAll()
}

func (r *RouteService) GetRouteById(routeId int) (models.Route, error) {
	return r.repository.GetRouteById(routeId)
}

func (r *RouteService) GetRouteByURL(url string) (models.Route, error) {
	return r.repository.GetRouteByURL(url)
}

func (r *RouteService) Update(route models.Route) (models.Route, error) {
	return r.repository.Update(route)
}

func (r *RouteService) Delete(routeId int) (bool, error) {
	return r.repository.Delete(routeId)
}
