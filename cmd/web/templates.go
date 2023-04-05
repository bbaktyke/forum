package main

import (
	"text/template"
	"time"

	"git.01.alem.school/bbaktyke/forum.git/pkg/forms"
	"git.01.alem.school/bbaktyke/forum.git/pkg/models"
)

type templateData struct {
	IsAuthenticated  bool
	Form             *forms.Form
	Post             *models.Posts
	Posts            []*models.Posts
	Comments         []*models.Comment
	Reaction         *models.Reaction
	Comment_Reaction *models.Comment_Reaction
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}
