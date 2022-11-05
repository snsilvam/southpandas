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

func (r *ElasticSearchRepository) IndexUserExternalWorker(ctx context.Context, userExternalWorker models.UserExternalWorker) error {
	body, _ := json.Marshal(userExternalWorker)
	_, err := r.client.Index(
		"userExternalWorkers",
		bytes.NewReader(body),
		r.client.Index.WithDocumentID(userExternalWorker.ID),
		r.client.Index.WithContext(ctx),
		r.client.Index.WithRefresh("wait_for"),
	)
	return err
}

func (r *ElasticSearchRepository) SearchUserExternalWorker(ctx context.Context, query string) (results []models.UserExternalWorker, err error) {
	var buf bytes.Buffer
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":            query,
				"fields":           []string{"contract_type", "work_experience", "work_remote", "willingnesstravel", "current_salary", "expected_salary", "possibility_of_rotation", "profile_linkedln", "workarea", "description_workarea"},
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
		r.client.Search.WithIndex("userExternalWorkers"),
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

	var userExternalWorkers []models.UserExternalWorker
	for _, hit := range eRes["hits"].(map[string]interface{})["hits"].([]interface{}) {
		userExternalWorker := models.UserExternalWorker{}
		source := hit.(map[string]interface{})["_source"]
		marshal, err := json.Marshal(source)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(marshal, &userExternalWorker); err == nil {
			userExternalWorkers = append(userExternalWorkers, userExternalWorker)
		}
	}
	return userExternalWorkers, nil
}
