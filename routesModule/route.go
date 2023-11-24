package routesModule

import (
	"denizkpln7/rest-api/handlers"
	"denizkpln7/rest-api/middlewares"
	"denizkpln7/rest-api/repository"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RouteModule struct {
	r  fiber.Router
	db *gorm.DB
}

func (m *RouteModule) UsersModule() handlers.IUsersHandler {
	usersRespository := repository.UsersRespository(m.db)
	usershandler := handlers.UsersHandler(usersRespository)
	return usershandler
}

func (m *RouteModule) CreditCardsModule() handlers.ICreditCardsHandler {
	creditcardsRepository := repository.CreditCardsRespository(m.db)
	creditcardshandler := handlers.CreditCardsHandler(creditcardsRepository)
	return creditcardshandler
}

func (m *RouteModule) Setup(app *fiber.App) {

	//user
	app.Post("/api/user/register", m.UsersModule().Save)
	app.Post("/api/user/login", m.UsersModule().Login)

	//use middleware
	app.Use(middlewares.IsAuthenticate)

	//user
	app.Get("/api/user/getall", m.UsersModule().GetAllUser)
	app.Get("/api/user/:id", m.UsersModule().GetByIdUser)

	//creditcards
	app.Post("/api/creditcards/save", m.CreditCardsModule().CreateCard)
	app.Get("/api/creditcards/:id", m.CreditCardsModule().GetCreateCardById)
}
