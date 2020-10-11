package recipe

import (
	"reflect"
	"testing"
)

var testRepo = newTestRepository(db)

func TestService_InsertRecipe(t *testing.T) {
	recipe := Recipe{
		Title:       "Tomato Soup",
		MakingTime:  "15 min",
		Serves:      "5 people",
		Ingredients: "onion, tomato, seasoning, water",
		Cost:        450,
	}
	type args struct {
		r *Recipe
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"insert recipe",
			args{&recipe},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: testRepo,
			}
			got, err := s.InsertRecipe(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.InsertRecipe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			want := testRepo.db["3"]
			want.FormattedCreatedAt = want.CreatedAt.Format(customTimeFormat)
			want.FormattedUpdatedAt = want.UpdatedAt.Format(customTimeFormat)

			if !reflect.DeepEqual(got, &want) {
				t.Errorf("Service.InsertRecipe() = %v, want %v", got, &want)
			}
		})
	}
}

func TestService_GetRecipies(t *testing.T) {
	want := []Recipe{}
	for _, r := range testRepo.db {
		want = append(want, r)
	}

	tests := []struct {
		name    string
		want    []Recipe
		wantErr bool
	}{
		{
			"get all recipes",
			want,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: testRepo,
			}
			got, err := s.GetRecipies()
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetRecipies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetRecipies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetRecipieByID(t *testing.T) {
	want := testRepo.db["1"]
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *Recipe
		wantErr bool
	}{
		{
			"get recipe 1",
			args{id: "1"},
			&want,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: testRepo,
			}
			got, err := s.GetRecipieByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetRecipieByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetRecipieByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UpdateRecipe(t *testing.T) {
	recipe := Recipe{
		Title:       "ssssss",
		MakingTime:  "15 min",
		Serves:      "5 people",
		Ingredients: "onion, tomato, seasoning, water",
		Cost:        99999,
	}
	type args struct {
		id string
		r  *Recipe
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			"update recipe 3",
			args{id: "3", r: &recipe},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: testRepo,
			}
			got, err := s.UpdateRecipe(tt.args.id, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UpdateRecipe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.UpdateRecipe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_DeleteRecipe(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			"delete recipe 3",
			args{id: "3"},
			1,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repo: testRepo,
			}
			got, err := s.DeleteRecipe(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.DeleteRecipe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.DeleteRecipe() = %v, want %v", got, tt.want)
			}
		})
	}
}
