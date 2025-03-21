/*
Copyright 2023 The K8sGPT Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ai

import (
	"bytes"
	"context"
	"encoding/json"

	"errors"
	"net/http"
)

const customRAGAIClientName = "customrag"

type CustomRagAIClient struct {
	nopCloser
}

func (c *CustomRagAIClient) Configure(_ IAIConfig) error {
	return nil
}

type RagResponse struct {
	RagRes string `json:"rag_response" bson:"rag_response"`
}

func (c *CustomRagAIClient) GetCompletion(_ context.Context, prompt string) (string, error) {
	values := map[string]string{"query": prompt} //#TODO understand the best way to generate promts
	json_data, _ := json.Marshal(values)

	resp, err := http.Post("http://localhost:8000/completion", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		return "error happened", errors.New("error geting response from customrag client")
	}

	defer resp.Body.Close()

	var res RagResponse

	json.NewDecoder(resp.Body).Decode(&res)

	return res.RagRes, nil
}

func (c *CustomRagAIClient) GetName() string {
	return customRAGAIClientName
}
