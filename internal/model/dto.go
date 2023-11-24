package model

type UnitDto struct {
	Id    int    `json:"id"`
	Label string `json:"label"`
}

type IngredientDto struct {
	Id          int    `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

type StepDto struct {
	Id              int                 `json:"id"`
	Label           string              `json:"label"`
	Description     string              `json:"description"`
	StepIngredients []StepIngredientDto `json:"step_ingredients"`
}

type StepIngredientDto struct {
	Ingredient IngredientDto `json:"ingredient"`
	Unit       UnitDto       `json:"unit"`
	Quantity   int           `json:"quantity"`
}

type RecipeDto struct {
	Id          int       `json:"id"`
	Label       string    `json:"label"`
	Description string    `json:"description"`
	Steps       []StepDto `json:"steps"`
}
