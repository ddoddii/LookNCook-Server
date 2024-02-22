package model

type Ingredient struct {
	Name                string `json:"name"`
	LocationDescription string `json:"locationDescription"`
}

type FridgeContent struct {
	Ingredients []Ingredient
}

type Step struct {
	Index  int      `json:"index"`
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Threat []string `json:"threat"`
}

type Recipe struct {
	Name           string       `json:"name"`
	ThumbnailImage string       `json:"thumbnailImage"`
	Steps          []Step       `json:"steps"`
	Level          string       `json:"level"`
	LikeCount      int          `json:"likeCount"`
	Ingredients    []Ingredient `json:"ingredients"`
}

type ApiResponse struct {
	Ingredients []Ingredient `json:"ingredients"`
	RecipeList  []Recipe     `json:"recipeList"`
}

type RecipeResponse struct {
	Fridge string `json:"fridge"`
	Recipe string `json:"recipe"`
}
