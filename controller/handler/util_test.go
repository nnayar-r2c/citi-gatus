package handler

import (
	"fmt"
	"github.com/TwiN/gatus/v5/config"
	"github.com/TwiN/gatus/v5/storage"
	"net/http"
	"testing"
)

func TestExtractPageAndPageSizeFromRequest(t *testing.T) {
	type Scenario struct {
		Name             string
		Page             string
		PageSize         string
		ExpectedPage     int
		ExpectedPageSize int
	}
	scenarios := []Scenario{
		{
			Page:             "1",
			PageSize:         "20",
			ExpectedPage:     1,
			ExpectedPageSize: 20,
		},
		{
			Page:             "2",
			PageSize:         "10",
			ExpectedPage:     2,
			ExpectedPageSize: 10,
		},
		{
			Page:             "2",
			PageSize:         "10",
			ExpectedPage:     2,
			ExpectedPageSize: 10,
		},
		{
			Page:             "1",
			PageSize:         "999999",
			ExpectedPage:     1,
			ExpectedPageSize: storage.DefaultMaximumNumberOfResults,
		},
		{
			Page:             "-1",
			PageSize:         "-1",
			ExpectedPage:     DefaultPage,
			ExpectedPageSize: DefaultPageSize,
		},
		{
			Page:             "invalid",
			PageSize:         "invalid",
			ExpectedPage:     DefaultPage,
			ExpectedPageSize: DefaultPageSize,
		},
	}
	storageCfg := &storage.Config{}
	storageCfg.ValidateAndSetDefaults()
	cfg := &config.Config{
		Storage: storageCfg,
	}
	for _, scenario := range scenarios {
		t.Run("page-"+scenario.Page+"-pageSize-"+scenario.PageSize, func(t *testing.T) {
			request, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/statuses?page=%s&pageSize=%s", scenario.Page, scenario.PageSize), http.NoBody)
			actualPage, actualPageSize := extractPageAndPageSizeFromRequest(request, cfg)
			if actualPage != scenario.ExpectedPage {
				t.Errorf("expected %d, got %d", scenario.ExpectedPage, actualPage)
			}
			if actualPageSize != scenario.ExpectedPageSize {
				t.Errorf("expected %d, got %d", scenario.ExpectedPageSize, actualPageSize)
			}
		})
	}
}
