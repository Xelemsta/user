package monitoring_test

import (
	"net/http"
	"testing"
	"user/testutils"

	"github.com/maxatome/go-testdeep/helpers/tdhttp"
	"github.com/maxatome/go-testdeep/td"
)

func TestGetMonPing(t *testing.T) {
	ta := tdhttp.NewTestAPI(t, testutils.InitAPI(t))

	ta.Get("/mon/ping").
		CmpStatus(http.StatusOK).
		CmpJSONBody(td.JSON(`{"status":"OK"}`))
}
