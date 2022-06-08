package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinpikaFE/go_fiber/models"
	"github.com/jinpikaFE/go_fiber/pkg/app"
	"github.com/jinpikaFE/go_fiber/pkg/logging"
	"github.com/jinpikaFE/go_fiber/pkg/untils"
	"github.com/jinpikaFE/go_fiber/pkg/valodates"
)

// 获取User
func GetUser(c *fiber.Ctx) error {
	// maps := make(map[string]interface{})
	// // 获取get query参数 或者使用queryparser
	// id := c.Query("id")
	// maps["id"] = id
	appF := app.Fiber{C: c}
	maps := &models.User{}

	if err := c.QueryParser(maps); err != nil {
		logging.Error(err)
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", err)
	}

	res, errs := models.GetUser(maps)

	if errs != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "查询失败", errs)
	}

	if !(res.ID > 0) {
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "用户不存在", nil)
	}

	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", res)
}

// 添加test
func AddUser(c *fiber.Ctx) error {
	appF := app.Fiber{C: c}
	user := &models.User{}

	if err := c.BodyParser(user); err != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "参数解析错误", err)
	}

	// 入参验证
	errors := valodates.ValidateStruct(*user)

	if errors != nil {
		return appF.Response(fiber.StatusBadRequest, fiber.StatusBadRequest, "检验参数错误", errors)
	}

	user.Password = untils.GetSha256(user.Password)

	err := models.AddUser(*user)

	if err != nil {
		return appF.Response(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "添加失败", err)
	}

	return appF.Response(fiber.StatusOK, fiber.StatusOK, "SUCCESS", user)
}
