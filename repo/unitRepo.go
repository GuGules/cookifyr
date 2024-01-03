type UnitRepository struct {
	db *sqlx.DB
}

func (repo *UnitRepository) FindAll() []model.Unit {
	int id
	string label
}