type IngredientRepository struct {
	db *sqlx.DB
}

func (repo *IngredientRepository) FindAll() []model.Ingredient {
	int id
	string label
	string description
}