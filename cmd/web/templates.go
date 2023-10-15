package main

import (
	"github.com/SegniAdebaGodsSon/snippetbox/pkg/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
