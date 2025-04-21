package http

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/jmechavez/jmgofolio/internal/ports"
)

type AppHandler struct {
	PortfolioService ports.PortfolioService
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/about.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h AppHandler) PortfolioJSONHandler(w http.ResponseWriter, r *http.Request) {

	portfolio, err := h.PortfolioService.MyPortfolio()
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err)
		return
	}
	if portfolio == nil {
		writeResponse(w, http.StatusNotFound, nil)
		return
	}

	writeResponse(w, http.StatusOK, portfolio)
}

//	func ProjectsJSONHandler(w http.ResponseWriter, r *http.Request) {
//		json.NewEncoder(w).Encode(db.Projects)
//	}
// func ProjectsJSONHandler(w http.ResponseWriter, r *http.Request) {
// 	writeResponse(w, http.StatusOK, db.Projects)
// }

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow frontend
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
