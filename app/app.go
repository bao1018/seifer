package app

import (
	"bitbucket.org/kyicy/seifer/app/router"
	"github.com/labstack/echo/v4"
)

// RegisterRoute function
func RegisterRoute(e *echo.Echo) {
	e.POST("/user_story", router.CreateUserStory)
	e.POST("/user_story/similar", router.SimilarUserStories)
}
