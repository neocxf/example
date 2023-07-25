package oa

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// GetAccessToken 获取ak
func (ex *Account) GetAccessToken(c *fiber.Ctx) (err error) {
	ak, err := ex.officialAccount.GetAccessToken()
	if err != nil {
		log.Errorf("get ak error, err=%+v", err)
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(ak)
	return
}

// GetCallbackIP 获取微信callback IP地址
func (ex *Account) GetCallbackIP(c *fiber.Ctx) (err error) {
	ipList, err := ex.officialAccount.GetBasic().GetCallbackIP()
	if err != nil {
		log.Errorf("GetCallbackIP error, err=%+v", err)
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(ipList)
	return
}

// GetAPIDomainIP 获取微信callback IP地址
func (ex *Account) GetAPIDomainIP(c *fiber.Ctx) (err error) {
	ipList, err := ex.officialAccount.GetBasic().GetAPIDomainIP()
	if err != nil {
		log.Errorf("GetAPIDomainIP error, err=%+v", err)
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(ipList)
	return
}

// GetAPIDomainIP  清理接口调用次数
func (ex *Account) ClearQuota(c *fiber.Ctx) (err error) {
	err = ex.officialAccount.GetBasic().ClearQuota()
	if err != nil {
		log.Errorf("ClearQuota error, err=%+v", err)
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}
	c.JSON("success")
	return
}
