-- Create unit table
CREATE TABLE unit (
    id SERIAL PRIMARY KEY,
    label VARCHAR(255)
);

-- Create ingredient table
CREATE TABLE ingredient (
    id SERIAL PRIMARY KEY,
    label VARCHAR(255) NOT NULL,
    description TEXT
);

-- Create recipe table
CREATE TABLE recipe (
    id SERIAL PRIMARY KEY,
    label VARCHAR(255) NOT NULL,
    description TEXT
);

-- Create step table
CREATE TABLE step (
    id SERIAL PRIMARY KEY,
    recipe_id INT NOT NULL REFERENCES recipe(id),
    label VARCHAR(255) NOT NULL,
    description TEXT NOT NULL
);

-- Create step_ingredient table
CREATE TABLE step_ingredient (
    id_step INT NOT NULL REFERENCES step(id),
    id_ingredient INT NOT NULL REFERENCES ingredient(id),
    unit_id INT NOT NULL REFERENCES unit(id),
    quantity INT NOT NULL,
    PRIMARY KEY (id_step, id_ingredient)
);
