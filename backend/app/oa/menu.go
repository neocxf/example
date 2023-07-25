package oa

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func (ex *Account) GetMenus(c *fiber.Ctx) (err error) {
	_menus, err := ex.officialAccount.GetMenu().GetCurrentSelfMenuInfo()
	if err != nil {
		log.Errorf("ClearQuota error, err=%+v", err)
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}

	c.JSON(_menus)
	return
}

func (ex *Account) NewMenu(c *fiber.Ctx) (err error) {
	_menu := ex.officialAccount.GetMenu()

	var menuJson = `
	{
	 "button":[
           {
           "name":"带我看",
           "sub_button":[
           {
            "type":"view",
            "name":"快速了解",
            "url":"http://baidu.com"
           }]
  
          },
           {
           "name":"即刻体验",
           "sub_button":[
           {
            "type":"view",
            "name":"Android版",
            "url":"http://baidu.com"
           },
           {
            "type":"view",
            "name":"iOS版",
            "url":"http://baidu.com"
           }]
  
          },
          {
           "name":"精彩活动",
           "sub_button":[
           {
            "type":"view",
            "name":"最美母校投票活动",
             "url":"http://baidu.com"
           }]
  
          }
 
 		]
	}
	
	`

	err = _menu.SetMenuByJSON(menuJson)
	if err != nil {
		log.Errorf("SetMenuByJSON error, err=%+v", err)
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
		return
	}
	c.JSON(fiber.Map{
		"msg": "success",
	})
	return
}
