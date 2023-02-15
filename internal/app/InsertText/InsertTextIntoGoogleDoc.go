package inserttext

import (
	"log"

	"google.golang.org/api/docs/v1"
)

type InsertTextIntoGoogleDocTable struct {
}

func New() *InsertTextIntoGoogleDocTable {
	return &InsertTextIntoGoogleDocTable{}
}

func (it *InsertTextIntoGoogleDocTable) Insert(srv *docs.Service, table [][]string, documentId string) {
	doc, err := srv.Documents.Get(documentId).Do()
	if err != nil {
		log.Fatalf("Failed to retrieve document: %v", err)
	}

	// Find the first table in the document
	var gtable *docs.Table
	for _, elem := range doc.Body.Content {
		if elem.Table != nil {
			gtable = elem.Table
			break
		}
	}
	if gtable == nil {
		log.Fatal("No table found in document")
	}

	for i := len(table) - 1; i >= 0; i-- {
		if i >= len(gtable.TableRows) {
			log.Fatalf("Row index out of range: %v", i)
		}
		for j := len(table[i]) - 1; j >= 0; j-- {
			if j >= len(gtable.TableRows[i].TableCells) {
				log.Fatalf("Column index out of range: %v", j)
			}
			req := &docs.BatchUpdateDocumentRequest{Requests: []*docs.Request{}}
			req.Requests = append(req.Requests, &docs.Request{
				InsertText: &docs.InsertTextRequest{
					Text: table[i][j],
					Location: &docs.Location{
						Index: gtable.TableRows[i].TableCells[j].StartIndex + 1,
					},
				},
			})
			_, err = srv.Documents.BatchUpdate(documentId, req).Do()
			if err != nil {
				log.Fatalf("Unable to insert data into table: %v", err)
			}
		}
	}

}
