package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sendgrid/sendgrid-go/models"
	"github.com/sendgrid/sendgrid-go/views"
)

func NewCampaigns(cs models.CampaignService) *Campaigns {
	return &Campaigns{
		CreateCampaignsTempl: views.NewView("bootstrap", "campaigns/campaigns"),
		UpdateCampaignsTempl: views.NewView("bootstrap", "campaigns/updatecampaigns"),

		CampaignService: cs,
	}
}

type Campaigns struct {
	models.CampaignService

	CreateCampaignsTempl *views.View
	UpdateCampaignsTempl *views.View
}

func (c Campaigns) ServeCampaignForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := c.CreateCampaignsTempl.Render(w, nil)
	if err != nil {
		ServeInternalServerError(w, err)
		return
	}
}

func (c Campaigns) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	body := struct {
		Title          string `schema:"campaign"`
		FormSender     string `schema:"fsender"`
		Subject        string `schema:"subject"`
		EmailPreheader string `schema:"emailph"`
		Content        string `schema:"content"`
	}{}
	if err := parseForm(r, &body); err != nil {
		panic(err)
	}

	e := models.Campaign{
		Title:      body.Title,
		FormSender: body.FormSender,
		Subject:    body.Subject,
		Content:    body.Content,
	}

	if err := c.CampaignService.Create(&e); err != nil {
		ServeInternalServerError(w, err)
		return
	}
}

func (c Campaigns) ServeUpdateCampaignForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ServeInternalServerError(w, err)
		return
	}

	_, err = c.CampaignService.ByID(uint(id))
	if err != nil {
		RecordNotFound(w, err)
		return
	}
	err = c.UpdateCampaignsTempl.Render(w, nil)
	if err != nil {
		ServeInternalServerError(w, err)
		return
	}
}

func (c Campaigns) UpdateCampaign(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ServeInternalServerError(w, err)
		return
	}
	senderID, err := c.CampaignService.ByID(uint(id))
	if err != nil {
		RecordNotFound(w, err)
		return
	}

	body := struct {
		Title          string `schema:"campaign"`
		FormSender     string `schema:"fsender"`
		Subject        string `schema:"subject"`
		EmailPreheader string `schema:"emailph"`
		Content        string `schema:"content"`
	}{}
	if err := parseForm(r, &body); err != nil {
		panic(err)
	}

	senderID.Title = body.Title
	senderID.FormSender = body.FormSender
	senderID.Subject = body.Subject
	senderID.EmailPreheader = body.EmailPreheader
	senderID.Content = body.Content

	err = c.CampaignService.Update(senderID)
	if err != nil {
		ServeInternalServerError(w, err)
		return
	}
}
