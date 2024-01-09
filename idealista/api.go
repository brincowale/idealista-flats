package idealista

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	Client       *http.Client
	RequestDelay time.Duration
}

func New(requestDelay time.Duration) *Client {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	return &Client{
		Client:       client,
		RequestDelay: requestDelay,
	}
}

type Property struct {
	Id            string
	Title         string
	URL           string
	Price         string
	Description   string
	M2            string
	Neighbourhood string
}

func (ic *Client) ScrapeProperties(URL string) ([]*Property, error) {
	var properties []*Property
	doc, err := ic.fetchHTMLDocument(URL)
	if err != nil {
		return nil, fmt.Errorf("error fetching search results: %w", err)
	}

	doc.Find("article.item").Each(func(i int, s *goquery.Selection) {
		detailsURL, _ := s.Find("a.item-link ").Attr("href")
		property := &Property{
			Title: strings.TrimSpace(s.Find("a.item-link ").Text()),
			URL:   "https://www.idealista.com" + detailsURL,
			Price: s.Find(".item-price").Text(),
			M2:    s.Find(".item-detail:contains('m²')").Text(),
		}
		property.Id = fmt.Sprintf("%s_%s_%s", property.Title, property.Price, property.M2)
		properties = append(properties, property)
	})
	return properties, nil
}

func (ic *Client) ScrapeAdditionalDetails(property *Property) error {
	doc, err := ic.fetchHTMLDocument(property.URL)
	if err != nil {
		return fmt.Errorf("error fetching property page (%s): %w", property.URL, err)
	}
	property.Description = strings.TrimSpace(doc.Find(".comment").Text())
	property.Neighbourhood = doc.Find(".main-info__title-minor").Text()
	return nil
}

func (ic *Client) fetchHTMLDocument(URL string) (*goquery.Document, error) {
	time.Sleep(ic.RequestDelay)

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	headers := map[string]string{
		"sec-ch-ua":                 `"Chromium";v="116", "Not)A;Brand";v="24", "Google Chrome";v="116"`,
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        `"Windows"`,
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Sec-Fetch-Site":            "none",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-User":            "?1",
		"Sec-Fetch-Dest":            "document",
		"Accept-Language":           "en-US,en;q=0.9",
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	response, err := ic.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch HTML document: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML document: %w", err)
	}

	return doc, nil
}
