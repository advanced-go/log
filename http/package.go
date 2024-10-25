package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/httpx"
	"github.com/advanced-go/common/uri"
	"github.com/advanced-go/log/event1"
	"github.com/advanced-go/log/module"
	"net/http"
)

const (
	PkgPath = "github/advanced-go/log/http"
	ver1    = "v1"
	ver2    = "v2"

	log                 = "log"
	healthLivenessPath  = "health/liveness"
	healthReadinessPath = "health/readiness"
	versionPath         = "version"
	authorityPath       = "authority"
)

var (
	authorityResponse = NewAuthorityResponse(module.Authority)
)

// Exchange - HTTP exchange function
func Exchange(r *http.Request) (*http.Response, *core.Status) {
	h2 := make(http.Header)
	h2.Add(httpx.ContentType, httpx.ContentTypeText)

	if r == nil {
		status := core.NewStatusError(http.StatusBadRequest, errors.New("request is nil"))
		return httpx.NewResponse(status.HttpCode(), h2, status.Err)
	}
	p, err := uri.ValidateURL(r.URL, module.Authority)
	if err != nil {
		status := core.NewStatusError(http.StatusBadRequest, err)
		resp1, _ := httpx.NewResponse(status.HttpCode(), h2, status.Err)
		return resp1, status
	}
	core.AddRequestId(r.Header)
	switch p.Resource {
	case log:
		resp, status1 := logExchange(r, p)
		resp.Header.Add(core.XRoute, event1.Route)
		return resp, status1
	case versionPath:
		resp, status1 := NewVersionResponse(module.Version), core.StatusOK()
		return resp, status1
	case authorityPath:
		resp, status1 := authorityResponse, core.StatusOK()
		return resp, status1
	case healthReadinessPath, healthLivenessPath:
		return httpx.NewHealthResponseOK(), core.StatusOK()
	default:
		status := core.NewStatusError(http.StatusNotFound, errors.New(fmt.Sprintf("error invalid URI, testresource not found: [%v]", p.Resource)))
		return httpx.NewResponse(status.HttpCode(), h2, status.Err)
	}
}
