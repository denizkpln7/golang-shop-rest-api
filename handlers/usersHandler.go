package handlers

import (
	"denizkpln7/rest-api/models"
	"denizkpln7/rest-api/repository"
	"denizkpln7/rest-api/utils"
	"errors"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

type IUsersHandler interface {
	Save(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	GetAllUser(c *fiber.Ctx) error
	GetByIdUser(c *fiber.Ctx) error
}

type usersHandler struct {
	usersRepository repository.IUsersRepository
}

func UsersHandler(usersRepository repository.IUsersRepository) IUsersHandler {
	return &usersHandler{
		usersRepository: usersRepository,
	}
}

func (u usersHandler) Save(c *fiber.Ctx) error {
	var req models.User
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Upss!")
	}
	if err := u.usersRepository.CreateUser(&req); err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Kaydetmede bi sıkıntı oldu!")
	}

	return c.JSON(fiber.Map{
		"message": "Congratulation!, Your post is live",
	})
}

func (u usersHandler) Login(c *fiber.Ctx) error {
	var req models.User
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Upss!")
	}
	user, err := u.usersRepository.Login(&req)

	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, "Böyle bir kullanıcı yok")
	}

	token, err := utils.GenerateJwt(strconv.Itoa(int(user.ID)))
	if err != nil {
		return errors.New("böyle bir kullanıcı yok")
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "giriş işleminiz başarılı",
		"user":    user,
	})
}

func (u usersHandler) GetAllUser(c *fiber.Ctx) error {
	users := u.usersRepository.GetAllUser()
	return c.JSON(fiber.Map{
		"data": users,
	})
}

func (u usersHandler) GetByIdUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := u.usersRepository.GetById(id)
	return c.JSON(fiber.Map{
		"data": user,
	})
}
