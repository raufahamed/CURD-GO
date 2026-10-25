package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {
}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) HandleGetWorkoutById(w http.ResponseWriter,r*http.Request){

paramsWorkOutId := chi.URLParam(r,"id")
if paramsWorkOutId == ""{
	http.NotFound(w,r)
	return 
}

workoutID , err := strconv.ParseInt(paramsWorkOutId,10,64)

if err !=nil{
	http.NotFound(w,r)
	return
}
fmt.Fprintf(w,"this is the workoutId %d\n",workoutID)

}

func (wh*WorkoutHandler)HandleCreateWorkOut(w http.ResponseWriter,r*http.Request){

	fmt.Fprintln(w,"create a workout")
}