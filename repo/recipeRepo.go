type RecipeRepository struct {
	db *sqlx.DB
}

func (repo *RecipeRepository) FindAll() []model.Recipe {
	int id
	string label
	string description
}