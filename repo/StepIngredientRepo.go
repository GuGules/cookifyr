type StepIngredientRepository struct {
	db *sqlx.DB
}

func (repo *StepIngredientRepository) FindAll() []model.StepIngredient {
	int id_step
	int id_ingredient
	int unit_id
	int quantity
}