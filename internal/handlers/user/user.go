package user

import (
	"net/http"

	"github.com/juju/errors"
)

func errorToHTTPCode(err error) int {
	httpCode := http.StatusInternalServerError
	if errors.Is(err, errors.BadRequest) {
		httpCode = http.StatusBadRequest
	} else if errors.Is(err, errors.NotFound) {
		httpCode = http.StatusNotFound
	}

	return httpCode
}
