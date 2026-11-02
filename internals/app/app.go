package app

import (
	"database/sql"
	"fem_project/internals/api"
	"fem_project/internals/store"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Application struct {

	Logger *log.Logger  

	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication()(*Application , error){

	pgDb,err := store.Open()
	if err !=nil{
		return nil , err 
	}

logger := log.New(os.Stdout,"",log.Ldate|log.Ltime)

workoutHandler := api.NewWorkoutHandler()

 app := &Application{
	Logger: logger,
	WorkoutHandler: workoutHandler,
	DB: pgDb,
 }
 return app,nil

}


func (a*Application) HealthCheck(w http.ResponseWriter , r*http.Request){
	fmt.Println("status is available")
}