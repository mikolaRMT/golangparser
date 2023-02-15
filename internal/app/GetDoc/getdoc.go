package getdoc

import (
	"log"

	"google.golang.org/api/docs/v1"
)

type GetDoc struct {
}

func New() *GetDoc {
	return &GetDoc{}
}

func (g *GetDoc) Get(srv *docs.Service) *docs.Document {
	resp, err := srv.Documents.Get("14H1wjM8s7qZfnvyxQVslp_WH1hDnRdzY1E7SlHdHCoQ").Do() //Сделать что-то с id
	if err != nil {
		doc := &docs.Document{
			Title: "Table",
		}
		// Create an empty document.
		doc, err = srv.Documents.Create(doc).Do()
		if err != nil {
			log.Fatalf("Unable to create document: %v", err)
		}
		return doc
	}
	return resp
}
