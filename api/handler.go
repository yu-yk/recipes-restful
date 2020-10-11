package api

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/yu-yk/recipes-restful/recipe"
)

func requestLogMiddleware(c *gin.Context) {
	// Read the Body content
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}
	// Restore the io.ReadCloser to its original state
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	body, _ := ioutil.ReadAll(bytes.NewBuffer(bodyBytes))
	log.Println("---body/--- \r\n" + string(body))
	c.Next()
}

func (s *Server) postRecipeHandler(c *gin.Context) {
	repository := recipe.NewMySQLRepository(s.db)
	recipeService := recipe.NewService(repository)

	var r recipe.Recipe
	c.ShouldBindBodyWith(&r, binding.JSON)

	lastRecipe, err := recipeService.InsertRecipe(&r)
	if err == recipe.ErrCreateRecipe {
		c.JSON(http.StatusOK, gin.H{
			"message":  err.Error(),
			"required": "title, preparation_time, serves, ingredients, cost",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
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
		"recipe":  []*recipe.Recipe{r},
	})
}

func (s *Server) updateRecipeHandler(c *gin.Context) {
	repository := recipe.NewMySQLRepository(s.db)
	recipeService := recipe.NewService(repository)

	id := c.Param("id")
	var r recipe.Recipe
	c.ShouldBindBodyWith(&r, binding.JSON)

	_, err := recipeService.UpdateRecipe(id, &r)
	if err == recipe.ErrUpdateRecipe {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":  err.Error(),
			"required": "title, preparation_time, serves, ingredients, cost",
		})
		return
	} else if err == recipe.ErrNotFound {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
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
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe successfully removed!",
	})
}
