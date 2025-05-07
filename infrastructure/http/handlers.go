package http

import (
	"encoding/json"
	"html/template"
	"log"
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

// func (h AppHandler) PorfolioHTTPHandler(w http.ResponseWriter, r *http.Request) {
// 	portfolio, err := h.PortfolioService.MyPortfolio()
// 	if err != nil {
// 		w.Header().Set("Content-Type", "text/html")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("<div class='error'>Internal Server Error</div>"))
// 		return
// 	}
// 	if portfolio == nil {
// 		w.Header().Set("Content-Type", "text/html")
// 		w.WriteHeader(http.StatusNotFound)
// 		w.Write([]byte("<div class='not-found'>Portfolio not found</div>"))
// 		return
// 	}
// 	// Set headers
// 	w.Header().Set("Content-Type", "text/html")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	// Parse template
// 	tmpl, _ := template.ParseFiles("templates/index_test.html")

//		// Execute and write the template
//		w.WriteHeader(http.StatusOK)
//		tmpl.Execute(w, portfolio)
//	}
var templates *template.Template

func init() { // Or in main/setup
	var err error
	// Parse *all* needed template files together
	templates, err = template.ParseFiles(
		"templates/layout.html",
		"templates/portfolio_card.html",
		"templates/sidebar.html",
		"templates/main.html",
		// Add other template files here if needed
	)
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
	log.Println("Templates parsed successfully.")
}

func (h AppHandler) PorfolioHTTPHandler(w http.ResponseWriter, r *http.Request) {
	portfolio, _ := h.PortfolioService.MyPortfolio()
	// ... (error handling for service call) ...

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// ... (other headers) ...
	w.WriteHeader(http.StatusOK)

	// Execute the specific template you want to render as the main entry point
	// Often this is the layout file or a specific page template.
	// The name usually defaults to the base name of the file.
	if err := templates.ExecuteTemplate(w, "layout.html", portfolio); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// Note: Inside layout.html, {{ template "portfolioCard" . }}
	// will find and render the template defined in portfolio_card.html,
	// passing the same 'portfolio' data down to it because we used '.'
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
