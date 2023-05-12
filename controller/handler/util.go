package handler

import (
	"github.com/TwiN/gatus/v5/config"
	"net/http"
	"strconv"
)

const (
	// DefaultPage is the default page to use if none is specified or an invalid value is provided
	DefaultPage = 1

	// DefaultPageSize is the default page siZE to use if none is specified or an invalid value is provided
	DefaultPageSize = 20
)

func extractPageAndPageSizeFromRequest(r *http.Request, cfg *config.Config) (page int, pageSize int) {
	var err error
	if pageParameter := r.URL.Query().Get("page"); len(pageParameter) == 0 {
		page = DefaultPage
	} else {
		page, err = strconv.Atoi(pageParameter)
		if err != nil {
			page = DefaultPage
		}
		if page < 1 {
			page = DefaultPage
		}
	}
	if pageSizeParameter := r.URL.Query().Get("pageSize"); len(pageSizeParameter) == 0 {
		pageSize = DefaultPageSize
	} else {
		pageSize, err = strconv.Atoi(pageSizeParameter)
		if err != nil {
			pageSize = DefaultPageSize
		}
		if pageSize > cfg.Storage.MaximumNumberOfResults {
			pageSize = cfg.Storage.MaximumNumberOfResults
		} else if pageSize < 1 {
			pageSize = DefaultPageSize
		}
	}
	return
}
