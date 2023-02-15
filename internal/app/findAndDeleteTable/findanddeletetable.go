package findanddeletetable

import (
	"log"

	"google.golang.org/api/docs/v1"
)

type FindAndDeleteTable struct {
}

func New() *FindAndDeleteTable {
	return &FindAndDeleteTable{}
}

func (f *FindAndDeleteTable) Delete(doc *docs.Document, srv *docs.Service) bool {
	var gtable *docs.Table
	for _, elem := range doc.Body.Content {
		if elem.Table != nil {
			gtable = elem.Table
			break
		}
	}
	if gtable == nil {
		log.Println("No table found in document")
		return true
	} else {
		endIndex := gtable.TableRows[gtable.Rows-1].EndIndex
		req := &docs.BatchUpdateDocumentRequest{Requests: []*docs.Request{}}
		req.Requests = append(req.Requests, &docs.Request{
			DeleteContentRange: &docs.DeleteContentRangeRequest{
				Range: &docs.Range{
					StartIndex: 1,
					EndIndex:   endIndex + 1,
				},
			},
		})
		_, err := srv.Documents.BatchUpdate(doc.DocumentId, req).Do()
		if err != nil {
			log.Printf("Unable to delete table in doc: %v", err)
			return false
		}
		return true
	}
}
