package model

type Unit struct {
	Id    int
	Label string
}

type Ingredient struct {
	Id          int
	Label       string
	Description string
}

type Step struct {
	Id          int
	RecipeId    int
	Label       string
	Description string
}

type StepIngredient struct {
	StepId       int
	IngredientId int
	UnitId       int
	Quantity     int
}

type Recipe struct {
	Id          int
	Label       string
	Description string
}
