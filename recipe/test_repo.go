package recipe

import (
	"strconv"
	"time"
)

var db = map[string]Recipe{
	"1": {
		1,
		"Chicken Curry",
		"45 min",
		"4 people",
		"onion, chicken, seasoning",
		1000,
		"2016-01-11 13:10:12",
		"2016-01-11 13:10:12",
		time.Now(), time.Now(),
	},
	"2": {
		2,
		"Rice Omelette",
		"30 min",
		"2 people",
		"onion, egg, seasoning, soy sauce",
		700,
		"2016-01-11 13:10:12",
		"2016-01-11 13:10:12",
		time.Now(), time.Now(),
	},
}

type testRepository struct {
	db      map[string]Recipe
	counter int
}

func newTestRepository(db map[string]Recipe) *testRepository {
	return &testRepository{
		db:      db,
		counter: len(db),
	}
}

func (repo *testRepository) InsertRecipe(r *Recipe) (int64, error) {
	recipe := Recipe{
		ID:          repo.counter + 1,
		Title:       r.Title,
		MakingTime:  r.MakingTime,
		Serves:      r.Serves,
		Ingredients: r.Ingredients,
		Cost:        r.Cost,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	repo.db[strconv.Itoa(recipe.ID)] = recipe
	lastID := recipe.ID
	repo.counter++

	return int64(lastID), nil
}

func (repo *testRepository) GetRecipies() ([]Recipe, error) {
	recipes := []Recipe{}
	for _, r := range repo.db {
		recipes = append(recipes, r)
	}
	return recipes, nil
}

func (repo *testRepository) GetRecipieByID(id string) (*Recipe, error) {
	r, ok := repo.db[id]
	if !ok {
		return nil, ErrNotFound
	}
	return &r, nil
}

func (repo *testRepository) UpdateRecipe(id string, r *Recipe) (int64, error) {
	if recipe, ok := repo.db[id]; ok {
		recipe.Title = r.Title
		recipe.MakingTime = r.MakingTime
		recipe.Serves = r.Serves
		recipe.Ingredients = r.Ingredients
		recipe.Cost = r.Cost
		recipe.UpdatedAt = time.Now()
		repo.db[id] = recipe
		return 1, nil
	}
	return 0, ErrNotFound
}

func (repo *testRepository) DeleteRecipe(id string) (int64, error) {
	_, ok := repo.db[id]
	if !ok {
		return 0, ErrNotFound
	}

	delete(repo.db, id)
	return 1, nil
}
