package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sendgrid/sendgrid-go/controllers"
	"github.com/sendgrid/sendgrid-go/models"
)

func main() {
	db, err := gorm.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}

	campaignGORM := models.NewCampaignGORM(db)
	if err != nil {
		log.Fatal(err)
	}

	campaignGORM.AutoMigrate()

	campaignC := controllers.NewCampaigns(campaignGORM)

	r := mux.NewRouter()

	r.NewRoute().
		Name("ServeCampaignForm").
		Methods("GET").
		Path("/campaigns").
		HandlerFunc(campaignC.ServeCampaignForm)

	r.NewRoute().
		Name("CreateCampaign").
		Methods("POST").
		Path("/campaigns").
		HandlerFunc(campaignC.CreateCampaign)

	r.NewRoute().
		Name("ServeUpdateCampaignForm").
		Methods("GET").
		Path("/campaigns/{id:[0-9]+}/edit").
		HandlerFunc(campaignC.ServeUpdateCampaignForm)

	r.NewRoute().
		Name("UpdateCampaign").
		Methods("POST").
		Path("/campaigns/{id:[0-9]+}/edit").
		HandlerFunc(campaignC.UpdateCampaign)

	err = http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		log.Fatal(err)
	}

}
