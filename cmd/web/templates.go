package main

//to centralise template rendering and caching

import (
	"html/template"
	"path/filepath"
	"time"

	"snippetbox.setkaung.net/internal/models"
)

// for passing multiple data to a template
type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateChache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	//filepath.Glob generates string slice from the pattern passed.
	//in this case, it is all .html with * as catch-all.
	//this makes it easier to add new html files without worrying
	//about forgetting to add html files manually.

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	// the loop below create a template for each page name
	// and the parse the base into that template
	// then parse the partials
	// and finally parse the file associated with the page name.
	// this produces a template cache for each page
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
	// the returned cache is used in render from helpers.go
}
