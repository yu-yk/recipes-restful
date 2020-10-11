package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yu-yk/recipes-restful/recipe"
)

func (s *Server) postRecipeHandler(c *gin.Context) {
	repository := recipe.NewMySQLRepository(s.db)
	recipeService := recipe.NewService(repository)

	var r recipe.Recipe
	c.Bind(&r)

	lastRecipe, err := recipeService.InsertRecipe(r)
	if err == recipe.ErrCreateRecipe {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":  err.Error(),
			"required": "title, preparation_time, serves, ingredients, cost",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Recipe successfully created!",
		"recipes": []*recipe.Recipe{lastRecipe},
	})
}

func (s *Server) getAllRecipesHandler(c *gin.Context) {
	repository := recipe.NewMySQLRepository(s.db)
	recipeService := recipe.NewService(repository)

	recipes, err := recipeService.GetRecipies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"recipes": recipes,
	})
}

func (s *Server) getRecipeHandler(c *gin.Context) {
	repository := recipe.NewMySQLRepository(s.db)
	recipeService := recipe.NewService(repository)

	id := c.Param("id")
	r, err := recipeService.GetRecipieByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe details by id",
		"recipes": []*recipe.Recipe{r},
	})
}

func (s *Server) updateRecipeHandler(c *gin.Context) {
	repository := recipe.NewMySQLRepository(s.db)
	recipeService := recipe.NewService(repository)

	id := c.Param("id")
	var r recipe.Recipe
	c.Bind(&r)

	_, err := recipeService.UpdateRecipe(id, r)
	if err == recipe.ErrUpdateRecipe {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":  err.Error(),
			"required": "title, preparation_time, serves, ingredients, cost",
		})
		return
	} else if err == recipe.ErrNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Recipe successfully updated!",
		"recipes": []recipe.Recipe{r},
	})
}

func (s *Server) deleteRecipeHandler(c *gin.Context) {
	repository := recipe.NewMySQLRepository(s.db)
	recipeService := recipe.NewService(repository)

	id := c.Param("id")

	_, err := recipeService.DeleteRecipe(id)
	if err == recipe.ErrNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Recipe successfully removed!",
	})
}
