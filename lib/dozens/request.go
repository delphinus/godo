package dozens

import (
	"net/http"

	"github.com/delphinus/godo/lib/dozens/path"
	"github.com/pkg/errors"
)

// MakeGet returns request for dozens
func MakeGet(p path.Path) (*http.Request, error) {
	req, err := http.NewRequest("GET", p.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error in NewRequest")
	}
	req.Header.Set("X-Auth-Token", Token())
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
