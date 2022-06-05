package controllers

import (
	"net/http"

	"github.com/avored/go-ecommerce/providers"
)

//Create a struct that holds information to be displayed in our HTML file
type DashboardViewModel struct {
	providers.BaseFlashes
}

func AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	dashboardViewModel := DashboardViewModel{}

	providers.RenderAdminView("templates/admin/dashboard.html", dashboardViewModel, w)
}
