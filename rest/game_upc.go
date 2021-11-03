package rest

import (
	"context"
	"fmt"
	"strings"

	"github.com/kevinanthony/bzen/config"
	"github.com/kevinanthony/bzen/http"
)

type GameUPC interface {
	GetIDFromUPC(ctx context.Context, upc string) (string, error)
}

type gameUPC struct {
	client http.Client
	api    string
	url    string
	prefix string
}

func NewGameUPC(cfg config.AppConfig, client http.Client) GameUPC {
	if client == nil {
		panic("http client is required")
	}

	return gameUPC{
		api:    cfg.GameUPC.APIKey,
		url:    cfg.GameUPC.URL,
		prefix: cfg.GameUPC.PathPrefix,
		client: client,
	}
}

func (g gameUPC) GetIDFromUPC(ctx context.Context, upc string) (string, error) {
	r := http.NewRequest(g.client).
		Get().
		Domain(g.url).
		Path(fmt.Sprintf("/%s/%s", strings.Trim(g.prefix, `/`), "upc/:upc")).
		Parameter(":upc", upc).
		Header("x-api-key", g.api)

	var resp GetUPCResponse

	if err := r.Go(ctx, &resp); err != nil {
		return "", err
	}

	return "", nil
}

type GetUPCResponse struct {
	UPC         string       `json:"upc"`
	Name        string       `json:"name"`
	SearchedFor string       `json:"searched_for"`
	Stage       string       `json:"stage"`
	Status      string       `json:"status"`
	BGGStatus   string       `json:"bgg_info_status"`
	Info        []UPCBGGInfo `json:"bgg_info"`
}

type UPCBGGInfo struct {
	BGGID         int             `json:"id"`
	Name          string          `json:"name"`
	PublishedData string          `json:"published"`
	ThumbnailURL  string          `json:"thumbnail_url"`
	ImageURL      string          `json:"image_url"`
	PageURL       string          `json:"page_url"`
	DataURL       string          `json:"data_url"`
	UpdateURL     string          `json:"update_url"`
	VersionStatus string          `json:"version_status"`
	Versions      []UPCBGGVersion `json:"versions"`
	Confidence    int             `json:"confidence"`
}

type UPCBGGVersion struct {
	VersionID    int    `json:"version_id"`
	Name         string `json:"name"`
	Published    string `json:"published"`
	Language     string `json:"language"`
	LanguageID   int    `json:"language_id"`
	ThumbnailURL string `json:"thumbnail_url"`
	ImageURL     string `json:"image_url"`
	UpdateURL    string `json:"update_url"`
	Confidence   int    `json:"confidence"`
}
