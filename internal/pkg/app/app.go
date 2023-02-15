package app

import (
	getdoc "golangparser/internal/app/GetDoc"
	inserttext "golangparser/internal/app/InsertText"
	connecttogoogleservice "golangparser/internal/app/connectToGoogleService"
	findanddeletetable "golangparser/internal/app/findAndDeleteTable"
	inserttableintogoogledoc "golangparser/internal/app/insertTableIntoGoogleDoc"
	parsehtml "golangparser/internal/app/parseHtml"
	"log"
)

type App struct {
	p  *parsehtml.ParseHtml
	g  *getdoc.GetDoc
	i  *inserttableintogoogledoc.InsertTableIntoGoogleDoc
	it *inserttext.InsertTextIntoGoogleDocTable
	c  *connecttogoogleservice.ConnectToGoogleService
	f  *findanddeletetable.FindAndDeleteTable
}

func New() (*App, error) {
	a := &App{}

	a.g = getdoc.New()
	a.p = parsehtml.New()
	a.i = inserttableintogoogledoc.New()
	a.it = inserttext.New()
	a.c = connecttogoogleservice.New()
	a.f = findanddeletetable.New()

	return a, nil
}

func (a *App) Run() error {
	srv, err := a.c.Connect()
	if err != nil {
		log.Fatalf("Error while connect to Google service")
	}
	link := "https://confluence.hflabs.ru/pages/viewpage.action?pageId=1181220999"
	doc := a.g.Get(srv)
	if a.f.Delete(doc, srv) {
		table := a.p.Parse(link)
		a.i.Insert(srv, doc.DocumentId, table)
		a.it.Insert(srv, table, doc.DocumentId)
	} else {
		log.Fatal("Unable to delete table in doc")
	}
	return nil
}
