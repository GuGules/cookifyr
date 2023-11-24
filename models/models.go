type unit struct {
	id    int
	label string
}

type ingredient struct {
	id          int
	label       string
	description string
}

type step struct {
	id          int
	recipe      recipe.id
	label       string
	description string
}

type step_ingredient struct {
	id_step       (step.id)
	id_ingredient (ingredient.id)
	unit_id       unit.id
	quantity      int
}

type recipe struct {
	id          int
	label       string
	description string
}