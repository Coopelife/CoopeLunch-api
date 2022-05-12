package main

import (
	"CoopeLunch-api/infrastructure"
	"CoopeLunch-api/interfaces/controllers"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	sqlHandler := infrastructure.NewSqlHandler()

	// Plan
	planController := controllers.NewPlanController(sqlHandler)
	router.Handle("/plan/", http.HandlerFunc(planController.PlanListView))

	srv := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}