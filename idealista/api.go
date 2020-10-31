package idealista

import (
	"github.com/getsentry/sentry-go"
	"github.com/go-resty/resty"
)

type Idealista struct {
	client *resty.Client
	Token  string
}

func New() *Idealista {
	return &Idealista{
		client: resty.New(),
	}
}

func (c Idealista) GetToken() string {
	var login *Login
	_, err := c.client.R().
		SetHeader("Authorization", "Basic NWI4NWMwM2MxNmJiYjg1ZDk2ZTIzMmIxMTJlZTg1ZGM6aWRlYSUzQmFuZHIwMWQ=").
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
		SetBody(`grant_type=client_credentials&scope=write`).
		SetResult(&login).
		Post("https://secure.idealista.com/api/oauth/token")
	if err != nil {
		sentry.CaptureException(err)
	}
	return login.AccessToken
}

func (c Idealista) GetProperties(search string) *Results {
	var properties *Results
	_, err := c.client.R().
		SetHeader("Authorization", "Bearer " + c.Token).
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
		SetBody(search).
		SetResult(&properties).
		Post("https://secure.idealista.com/api/3.5/es/search?k=5b85c03c16bbb85d96e232b112ee85dc")
	if err != nil {
		sentry.CaptureException(err)
	}
	return properties
}

func (c Idealista) GetProperty(propertyId string) PropertyDetails {
	var property PropertyDetails
	_, err := c.client.R().
		SetResult(&property).
		Get("https://secure.idealista.com/api/3/es/detail/" + propertyId + "?k=5b85c03c16bbb85d96e232b112ee85dc")
	if err != nil {
		sentry.CaptureException(err)
	}
	return property
}