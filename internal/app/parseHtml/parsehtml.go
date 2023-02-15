package parsehtml

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ParseHtml struct {
}

func New() *ParseHtml {
	return &ParseHtml{}
}
func (p *ParseHtml) Parse(link string) [][]string {
	// Example usage: extract data from table on Wikipedia page
	tableData := extractTableData(link)
	return tableData
}

func extractTableData(link string) [][]string {
	// Make HTTP request to fetch web page content
	response, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Parse HTML using goquery
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the first table on the page
	table := doc.Find("table").First()

	// Extract data from table rows
	var rows [][]string
	table.Find("tr").Each(func(rowIndex int, row *goquery.Selection) {
		// Extract data from table cells in this row
		var cells []string
		row.Find("td,th").Each(func(cellIndex int, cell *goquery.Selection) {
			cellText := strings.TrimSpace(cell.Text())
			cells = append(cells, cellText)
		})
		rows = append(rows, cells)
	})

	return rows
}
