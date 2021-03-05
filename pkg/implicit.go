package pkg

import (
	"embed"
	_ "embed"
	"html/template"
	"net/http"
)

//go:embed implicit/index.html
var implicit string

//go:embed implicit/*
var static embed.FS

type ImplicitTemplateContext struct {
	ClientID     string
	DiscoveryURL string
	RootURL      string
}

func (c *OIDCClient) implicit(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.New("name").Parse(implicit)
	// If the provider URL ends in a slash, we have to remove it
	// since "openid-implicit-client" re-adds a trailing slash, which can cause 404s
	if string(c.providerURL[len(c.providerURL)-1]) == "/" {
		c.providerURL = c.providerURL[0:(len(c.providerURL) - 1)]
	}
	context := ImplicitTemplateContext{
		ClientID:     c.config.ClientID,
		DiscoveryURL: c.providerURL,
		RootURL:      c.rootURL,
	}
	tmpl.Execute(w, context)
}
