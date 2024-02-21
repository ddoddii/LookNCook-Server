package config

import (
	"github.com/gdsc-ys/lookncook-server/src/handler"
	"github.com/gdsc-ys/lookncook-server/src/service"
)

type Initialization struct {
	UserService    service.UserService
	KitchenService service.KitchenService
	RecipeService  service.RecipeService
	GeminiService  service.GeminiService

	UserHandler    handler.UserHandler
	KitchenHandler handler.KitchenHandler
	RecipeHandler  handler.RecipeHandler
}

func NewInitialization(
	userService service.UserService,
	kitchenService service.KitchenService,
	recipeService service.RecipeService,
	geminiService service.GeminiService,

	userHandler handler.UserHandler,
	kitchenHandler handler.KitchenHandler,
	recipeHandler handler.RecipeHandler,
) *Initialization {
	return &Initialization{
		UserService:    userService,
		KitchenService: kitchenService,
		RecipeService:  recipeService,
		GeminiService:  geminiService,
		UserHandler:    userHandler,
		KitchenHandler: kitchenHandler,
		RecipeHandler:  recipeHandler,
	}
}
