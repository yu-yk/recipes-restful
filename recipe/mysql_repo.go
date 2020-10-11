package recipe

import (
	"database/sql"
	"log"
	"time"
)

var mySQLRepository *MySQLRepository

// MySQLRepository persists a mysql database
type MySQLRepository struct {
	*sql.DB
}

// NewMySQLRepository return the mysql repository
func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	if mySQLRepository != nil {
		return mySQLRepository
	}
	mySQLRepository = &MySQLRepository{db}
	return mySQLRepository
}

// InsertRecipe is the mysql implementation of insert recipe
func (db *MySQLRepository) InsertRecipe(r *Recipe) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO recipes(title, making_time, serves, ingredients, cost, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(r.Title, r.MakingTime, r.Serves, r.Ingredients, r.Cost, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

// GetRecipies is the mysql implementation of get all recipes
func (db *MySQLRepository) GetRecipies() ([]Recipe, error) {
	rows, err := db.Query("SELECT * FROM recipes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []Recipe

	for rows.Next() {
		var r Recipe
		err := rows.Scan(&r.ID, &r.Title, &r.MakingTime, &r.Serves, &r.Ingredients, &r.Cost, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, r)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return recipes, nil
}

// GetRecipieByID is the mysql implementation of get recipe by id
func (db *MySQLRepository) GetRecipieByID(id string) (*Recipe, error) {
	stmt, err := db.Prepare("SELECT * FROM recipes WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var r Recipe
	err = stmt.QueryRow(id).Scan(&r.ID, &r.Title, &r.MakingTime, &r.Serves, &r.Ingredients, &r.Cost, &r.CreatedAt, &r.UpdatedAt)
	if err != nil {
		return &r, err
	}
	return &r, nil
}

// UpdateRecipe is the mysql implementation of update recipe by id
func (db *MySQLRepository) UpdateRecipe(id string, r *Recipe) (int64, error) {
	stmt, err := db.Prepare("UPDATE recipes SET title = ?, making_time = ?, serves = ?, ingredients = ?, cost = ? WHERE id = ?")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(r.Title, r.MakingTime, r.Serves, r.Ingredients, r.Cost, id)
	if err != nil {
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affected, nil
}

// DeleteRecipe is the mysql implementation of delete recipe by id
func (db *MySQLRepository) DeleteRecipe(id string) (int64, error) {
	stmt, err := db.Prepare("DELETE FROM recipes WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affected, err
}
