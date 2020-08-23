package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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
	logger    *zap.Logger
	apiConfig config.APIConfig
}

// NewService returns a new BotService instance
func NewService(logger *zap.Logger, cfg config.APIConfig) Service {
	return BotService{
		logger:    logger,
		apiConfig: cfg,
	}
}

// GetDefinition makes a call to Dictionaries API and returns an instance of api.Response
func (b BotService) GetDefinition(query string) (api.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://od-api.oxforddictionaries.com:443/api/v2/entries/en/"+strings.ToLower(query), nil)
	if err != nil {
		b.logger.Error(fmt.Sprintf("[Service] [BotService] [GetDefinition] [NewRequest] %v", err))
		return api.Response{}, err
	}

	req.Header.Add(apiAppIDHeader, b.apiConfig.GetID())
	req.Header.Add(apiAppKeyHeader, b.apiConfig.GetKey())

	res, err := client.Do(req)
	if err != nil {
		b.logger.Error(fmt.Sprintf("[Service] [BotService] [GetDefinition] [Do] %v", err))
		return api.Response{}, err
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
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		b.logger.Error(fmt.Sprintf("[Service] [BotService] [GetDefinition] [ReadAll] %v", err))
		return api.Response{}, err
	}
	var r api.Response
	err = json.Unmarshal(data, &r)
	if err != nil {
		b.logger.Error(fmt.Sprintf("[Service] [BotService] [GetDefinition] [Unmarshal] %v", err))
		return api.Response{}, err
	}
	return r, nil
}