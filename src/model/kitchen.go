package model

type KitchenEnvironmentResponse struct {
	Hazard []Hazard `json:"hazards"`
}

type Hazard struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}
