package model

type Ingredient struct {
	Name                string `json:"name"`
	LocationDescription string `json:"locationDescription"`
}

type Recipe struct {
	Title        string       `json:"title"`
	Difficulty   string       `json:"difficulty"`
	Ingredients  []Ingredient `json:"ingredients"`
	Instructions []Step       `json:"instructions"`
}

type Step struct {
	Step string `json:"step"`
}

type RecipeResponse struct {
	Ingredients []Ingredient `json:"ingredients"`
	Recipes     []Recipe     `json:"recipes"`
}
