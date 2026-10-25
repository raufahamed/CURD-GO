package main

import (
	"fem_project/internals/app"
	"fem_project/internals/routes"

	"flag"
	"fmt"
	"net/http"
	"time"
)

func main(){
	
	var port int 
	flag.IntVar(&port,"port",8080,"go backend server")

	flag.Parse()

	app , err := app.NewApplication()

	if err !=nil {
		panic(err)
	}

	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr: fmt.Sprintf(":%d",port),
		Handler: r,
		IdleTimeout: time.Minute,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 30*time.Second,
	}

	app.Logger.Printf("Our app is started to run and runing on the port %d",port)


	err = server.ListenAndServe()

	if err!=nil{
		app.Logger.Fatal(err)
	}

}

