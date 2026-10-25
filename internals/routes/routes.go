package routes

import (
	"fem_project/internals/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux{
	r := chi.NewRouter()
	r.Get("/health",app.HealthCheck)
	r.Get("/workout/{id}",app.WorkoutHandler.HandleGetWorkoutById)
	r.Post("/create",app.WorkoutHandler.HandleCreateWorkOut)
	return r 
}