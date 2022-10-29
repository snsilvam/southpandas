package search

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	elastic "github.com/elastic/go-elasticsearch/v7"
	"southpandas.com/go/cqrs/models"
)

type ElasticSearchRepository struct {
	client *elastic.Client
}

func NewElastic(url string) (*ElasticSearchRepository, error) {
	client, err := elastic.NewClient(elastic.Config{
		Addresses: []string{url},
	})
	if err != nil {
		return nil, err
	}
	return &ElasticSearchRepository{client: client}, nil
}

func (r *ElasticSearchRepository) Close() {
	//Esta funcion, no esta disponible en la implementacion actual
}

func (r *ElasticSearchRepository) IndexworkerOfClient(ctx context.Context, workerOfClient models.WorkerOfClient) error {
	body, _ := json.Marshal(workerOfClient)
	_, err := r.client.Index(
		"workerOfClients",
		bytes.NewReader(body),
		r.client.Index.WithDocumentID(workerOfClient.ID),
		r.client.Index.WithContext(ctx),
		r.client.Index.WithRefresh("wait_for"),
	)
	return err
}

func (r *ElasticSearchRepository) SearchworkerOfClient(ctx context.Context, query string) (results []models.WorkerOfClient, err error) {
	var buf bytes.Buffer
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":            query,
				"fields":           []string{"description"},
				"fuzziness":        3,
				"cutoff_frequency": 0.0001,
			},
		},
	}
	if err = json.NewEncoder(&buf).Encode(searchQuery); err != nil {
		return nil, err
	}
	res, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex("workerOfClients"),
		r.client.Search.WithBody(&buf),
		r.client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			results = nil
		}
	}()
	if res.IsError() {
		return nil, errors.New(res.String())
	}

	var eRes map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&eRes); err != nil {
		return nil, err
	}

	var workerOfClients []models.WorkerOfClient
	for _, hit := range eRes["hits"].(map[string]interface{})["hits"].([]interface{}) {
		workerOfClient := models.WorkerOfClient{}
		source := hit.(map[string]interface{})["_source"]
		marshal, err := json.Marshal(source)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(marshal, &workerOfClient); err == nil {
			workerOfClients = append(workerOfClients, workerOfClient)
		}
	}
	return workerOfClients, nil
}
