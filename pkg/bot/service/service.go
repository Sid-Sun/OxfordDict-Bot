package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/store"

	"github.com/sid-sun/OxfordDict-Bot/cmd/config"
	"github.com/sid-sun/OxfordDict-Bot/pkg/bot/contract/api"
	"go.uber.org/zap"
)

// Service Interface defines a service spec
type Service interface {
	GetDefinition(string) (api.Response, error)
}

// BotService implements Service with logger
type BotService struct {
	client *http.Client
	logger *zap.Logger
	apiKey string
	store  store.Store
}

var ErrForbidden = errors.New("status 403 - access denied")

// NewService returns a new BotService instance
func NewService(logger *zap.Logger, cfg config.Config, str store.Store) Service {
	return BotService{
		client: &http.Client{},
		logger: logger,
		apiKey: cfg.APIKey,
		store:  str,
	}
}

// GetDefinition makes a call to Dictionaries API and returns an instance of api.Response
func (b BotService) GetDefinition(query string) (api.Response, error) {
	// Check if cache has definition
	r := b.store.Redis.Get(query)
	if !r.IsEmpty() {
		return r, nil
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://wordsapiv1.p.rapidapi.com/words/%s", query), nil)
	if err != nil {
		b.logger.Error(fmt.Sprintf("[Service] [BotService] [GetDefinition] [NewRequest] %v", err))
		return api.Response{}, err
	}

	req.Header.Add(apiHostHeader, "wordsapiv1.p.rapidapi.com")
	req.Header.Add(apiAppKeyHeader, b.apiKey)

	res, err := b.client.Do(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("[Service] [BotService] [GetDefinition] [Do] %v", err))
		return api.Response{}, err
	}
	if res.StatusCode == http.StatusForbidden {
		return api.Response{}, ErrForbidden
	}
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNotFound {
		err = fmt.Errorf("status code not OK: %d", res.StatusCode)
		b.logger.Error(fmt.Sprintf("[Service] [BotService] [GetDefinition] [StatusOK] %v", err))
		return api.Response{}, err
	}
	if contentType := strings.Split(res.Header.Get("Content-Type"), ";")[0]; contentType != "application/json" {
		err = fmt.Errorf("invalid response Content Type: %s", contentType)
		b.logger.Error(fmt.Sprintf("[Service] [BotService] [GetDefinition] [ContentTypeJSON] %v", err))
		return api.Response{}, err
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		b.logger.Error(fmt.Sprintf("[Service] [BotService] [GetDefinition] [ReadAll] %v", err))
		return api.Response{}, err
	}
	err = json.Unmarshal(data, &r)
	if err != nil {
		b.logger.Error(fmt.Sprintf("[Service] [BotService] [GetDefinition] [Unmarshal] %v", err))
		return api.Response{}, err
	}

	// Cache response in memory
	b.store.Redis.Put(query, r)
	return r, nil
}
