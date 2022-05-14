package html

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func FindAvailableReservations(htmlString string) (map[string]map[string]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlString))
	if err != nil {
		return nil, err
	}

	results := make(map[string]map[string]string)

	// Find reservation rows
	doc.Find("tr > th[class^=slot]").Each(func(i int, s *goquery.Selection) {
		startTime := s.Nodes[0].FirstChild.Data

		// Find available reservations
		s.Parent().Find("button").Each(func(i int, s *goquery.Selection) {
			id := ""
			formaction := ""

			// Parse attributes of buttons
			// <button id="..." formaction="..."></button>
			for _, attr := range s.Nodes[0].Attr {
				switch attr.Key {
				case "id":
					id = attr.Val
				case "formaction":
					formaction = attr.Val
				}
			}

			// Save parsed data
			if id != "" && formaction != "" {
				parts := strings.Split(id, "_")
				date := parts[len(parts)-1]
				if results[date] == nil {
					results[date] = make(map[string]string)
				}

				results[date][startTime] = formaction
			}
		})
	})

	return results, nil
}
