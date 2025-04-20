package db

import "github.com/jmechavez/jmgofolio/internal/domain"

var Projects = []domain.Project{
	{
		ID:          1,
		Title:       "Portfolio Website",
		Description: "Built with Go, HTMX, and Alpine.js",
		Link:        "https://github.com/yourname/portfolio",
	},
	{
		ID:          2,
		Title:       "Blog Engine",
		Description: "Markdown-powered blog built in Go",
		Link:        "https://github.com/yourname/blog",
	},
}
