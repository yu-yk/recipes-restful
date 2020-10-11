package api

// Serve serves the api server on the given address:port
func (s *Server) Serve(addr string) {

	s.router.POST("/recipes", s.postRecipeHandler)
	s.router.GET("/recipes", s.getAllRecipesHandler)
	s.router.GET("/recipes/:id", s.getRecipeHandler)
	s.router.PATCH("/recipes/:id", s.updateRecipeHandler)
	s.router.DELETE("/recipes/:id", s.deleteRecipeHandler)

	s.router.Run(addr)
}
