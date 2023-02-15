package inserttableintogoogledoc

import (
	"log"

	"google.golang.org/api/docs/v1"
)

type InsertTableIntoGoogleDoc struct {
}

func New() *InsertTableIntoGoogleDoc {
	return &InsertTableIntoGoogleDoc{}
}

func (i *InsertTableIntoGoogleDoc) Insert(srv *docs.Service, documentId string, table [][]string) error {

	req := &docs.BatchUpdateDocumentRequest{
		Requests: []*docs.Request{
			{
				InsertText: &docs.InsertTextRequest{
					Text: " ", // Insert a newline character to create a new paragraph
					Location: &docs.Location{
						Index: 1, // Insert at the beginning of the document
					},
				},
			},
		},
	}

	req.Requests = append(req.Requests, &docs.Request{
		InsertTable: &docs.InsertTableRequest{
			Columns: int64(len(table[0])),
			Rows:    int64(len(table)),
			Location: &docs.Location{
				Index: int64(1),
			},
		},
	})

	_, err := srv.Documents.BatchUpdate(documentId, req).Do()
	if err != nil {
		log.Fatalf("Unable to insert table: %v", err)
	}

	return nil
}
