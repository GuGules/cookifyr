type StepRepository struct {
	db *sqlx.DB
}

func (repo *StepRepository) FindAll() []model.Step {
	int id
	int recipe
	string label
	string description
}