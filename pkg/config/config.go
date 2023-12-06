package config

import (
	"html/template"
	"log/slog"

	scs "github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	Log           *slog.Logger
	Debug         bool
	Session       *scs.SessionManager
}
