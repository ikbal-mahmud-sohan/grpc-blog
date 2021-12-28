-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS posts(
    id SERIAL NOT NULL,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    category_id INTEGER NOT NULL,
    image TEXT NOT NULL, 
    is_completed BOOLEAN DEFAULT false,

    PRIMARY KEY(id),
    CONSTRAINT fk_category
      FOREIGN KEY(category_id) 
	  REFERENCES categories(id)

	  ON DELETE SET NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS posts;
-- +goose StatementEnd
