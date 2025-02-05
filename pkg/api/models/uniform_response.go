package models

import "encoding/json"

type RegisterIntegrationResponse struct {
	ID string `json:"id"`
}

func (i *RegisterIntegrationResponse) ToJSON() ([]byte, error) {
	if i == nil {
		return nil, nil
	}
	return json.Marshal(i)
}

func (i *RegisterIntegrationResponse) FromJSON(b []byte) error {
	var res RegisterIntegrationResponse
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*i = res
	return nil
}

type CreateSubscriptionResponse struct {
	ID string `json:"id"`
}

func (s *CreateSubscriptionResponse) ToJSON() ([]byte, error) {
	if s == nil {
		return nil, nil
	}
	return json.Marshal(s)
}

func (s *CreateSubscriptionResponse) FromJSON(b []byte) error {
	var res CreateSubscriptionResponse
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*s = res
	return nil
}
