package endpoint

import (
	"fmt"
	"net/url"
	"strings"
)

var baseURL *url.URL

func init() {
	url, err := url.Parse("http://dozens.jp/api")
	if err != nil {
		panic(err)
	}
	baseURL = url
}

// Endpoint means the path struct
type Endpoint struct {
	Base  *url.URL
	Chunk string
}

// NewEndpoint returns Endpoint struct
func NewEndpoint(chunk string) Endpoint {
	return Endpoint{baseURL, chunk}
}

func (p Endpoint) String() string {
	u, err := url.Parse(p.Base.String())
	if err != nil {
		panic(err)
	}
	u.Path = strings.Join([]string{u.Path, p.Chunk}, "/")
	return u.String()
}

// Authorize means `http://dozens.jp/api/authorize.json`
func Authorize() Endpoint {
	return NewEndpoint("authorize.json")
}

// ZoneList means `http://dozens.jp/api/zone.json`
func ZoneList() Endpoint {
	return NewEndpoint("zone.json")
}

// ZoneCreate means `http://dozens.jp/api/zone/create.json`
func ZoneCreate() Endpoint {
	return NewEndpoint("zone/create.json")
}

// ZoneUpdate means `http://dozens.jp/api/zone/update/:zone_id.json`
func ZoneUpdate(zoneID string) Endpoint {
	return NewEndpoint(fmt.Sprintf("zone/update/%s.json", zoneID))
}

// ZoneDelete means `http://dozens.jp/api/zone/delete/:zone_id.json`
func ZoneDelete(zoneID string) Endpoint {
	return NewEndpoint(fmt.Sprintf("zone/delete/%s.json", zoneID))
}
