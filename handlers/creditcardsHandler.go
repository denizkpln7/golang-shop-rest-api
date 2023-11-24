package handlers

import (
	"denizkpln7/rest-api/models"
	"denizkpln7/rest-api/repository"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ICreditCardsHandler interface {
	CreateCard(c *fiber.Ctx) error
	GetCreateCardById(c *fiber.Ctx) error
}

type creditcardsHandler struct {
	creditcardsRepository repository.ICreditCardsRepository
}

func CreditCardsHandler(creditcardsRepository repository.ICreditCardsRepository) ICreditCardsHandler {
	return &creditcardsHandler{
		creditcardsRepository: creditcardsRepository,
	}
}

func (cr creditcardsHandler) CreateCard(c *fiber.Ctx) error {
	var req models.CreditCard
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Upss!")
	}
	if err := cr.creditcardsRepository.CreateCard(&req); err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Kaydetmede bi sıkıntı oldu!")
	}

	return c.JSON(fiber.Map{
		"message": "Congratulation!, Your post is live",
	})
}

func (cr creditcardsHandler) GetCreateCardById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	creditCard := cr.creditcardsRepository.GetCreateCardById(id)

	return c.JSON(fiber.Map{
		"data": creditCard,
	})
}
