package router

import (
	"github.com/gorilla/mux"

	"transit-backend/internal/middleware"
	"transit-backend/internal/transport/handler"
	"transit-backend/internal/transport/model"
	"transit-backend/internal/transport/repository"
	"transit-backend/internal/transport/service"
)

func Setup(data *model.Data) *mux.Router {
	r := mux.NewRouter()

	r.Use(
		middleware.Logging,
		middleware.CORS,
		middleware.Recoverer,
	)

	api := r.PathPrefix("/api/v1").Subrouter()

	//Stations end points

	repo := repository.NewTransportRepository(data)
	linesService := service.NewLinesService(repo)
	linesHandler := handler.NewLinesHandler(linesService)

	api.HandleFunc("/transport", linesHandler.GetAll).Methods("GET")
	api.HandleFunc("/transport/{type}", linesHandler.GetByType).Methods("GET")

	api.HandleFunc("/transport/{type}/{line}", linesHandler.GetLineByType).Methods("GET")
	api.HandleFunc("/closest", linesHandler.GetClosestStation).Methods("GET")
	api.HandleFunc("/transport/{type}/{line}/{flag}", linesHandler.GetLineByType).Methods("GET")

	//Graph end points

	graphService := service.GetGraph(data.Lines)
	graphHandler := handler.NewGraphHandler(graphService)

	api.HandleFunc("/graph", graphHandler.GetGraphStructure).Methods("GET")
	api.HandleFunc("/graph/station", graphHandler.GetStationDetails).Methods("GET")
	api.HandleFunc("/graph/connections", graphHandler.GetStationConnections).Methods("GET")

	//route planner

	routePlanner := service.NewRoutePlanner(graphService)
	routeHandler := handler.NewRouteHandler(routePlanner)
	api.HandleFunc("/routes", routeHandler.GetRoutes).Methods("GET")

	return r
}
