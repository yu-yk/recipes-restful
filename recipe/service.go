package recipe

import (
	"errors"
	"strconv"
)

const customTimeFormat = "2006-01-02 15:04:05"

var ErrCreateRecipe = errors.New("Recipe creation failed!")
var ErrUpdateRecipe = errors.New("Recipe update failed!")
var ErrNotFound = errors.New("No recipe found")

type RecipeRepository interface {
	InsertRecipe(r Recipe) (int64, error)
	GetRecipies() ([]Recipe, error)
	GetRecipieByID(id string) (*Recipe, error)
	UpdateRecipe(id string, r Recipe) (int64, error)
	DeleteRecipe(id string) (int64, error)
}

type Service struct {
	repo RecipeRepository
}

func NewService(repo RecipeRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) InsertRecipe(r Recipe) (*Recipe, error) {
	if r.Title == "" || r.MakingTime == "" || r.Serves == "" || r.Ingredients == "" || r.Cost == 0 {
		return nil, ErrCreateRecipe
	}

	lastID, err := s.repo.InsertRecipe(r)
	if err != nil {
		return nil, err
	}

	lastRecipe, err := s.repo.GetRecipieByID(strconv.Itoa(int(lastID)))
	if err != nil {
		return nil, err
	}

	// format the time ouput
	lastRecipe.FormattedCreatedAt = lastRecipe.CreatedAt.Format(customTimeFormat)
	lastRecipe.FormattedUpdatedAt = lastRecipe.UpdatedAt.Format(customTimeFormat)

	return lastRecipe, nil
}

func (s *Service) GetRecipies() ([]Recipe, error) {
	return s.repo.GetRecipies()
}

func (s *Service) GetRecipieByID(id string) (*Recipe, error) {
	return s.repo.GetRecipieByID(id)
}

func (s *Service) UpdateRecipe(id string, r Recipe) (int64, error) {
	if r.Title == "" || r.MakingTime == "" || r.Serves == "" || r.Ingredients == "" || r.Cost == 0 {
		return 0, ErrUpdateRecipe
	}

	affected, err := s.repo.UpdateRecipe(id, r)
	if err != nil {
		return 0, err
	}

	if affected == 0 {
		return 0, ErrNotFound
	}

	return affected, nil
}

func (s *Service) DeleteRecipe(id string) (int64, error) {
	affected, err := s.repo.DeleteRecipe(id)
	if err != nil {
		return 0, err
	}

	if affected == 0 {
		return 0, ErrNotFound
	}

	return affected, nil
}
