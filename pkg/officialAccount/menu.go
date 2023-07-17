package officialAccount

import (
	"github.com/gin-gonic/gin"
	"github.com/gowechat/example/pkg/util"
	log "github.com/sirupsen/logrus"
)

func (ex *ExampleOfficialAccount) GetMenus(c *gin.Context) {
	_menus, err := ex.officialAccount.GetMenu().GetCurrentSelfMenuInfo()
	if err != nil {
		log.Errorf("ClearQuota error, err=%+v", err)
		util.RenderError(c, err)
		return
	}
	util.RenderSuccess(c, _menus)
}

func (ex *ExampleOfficialAccount) NewMenu(c *gin.Context) {
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

	err := _menu.SetMenuByJSON(menuJson)
	if err != nil {
		log.Errorf("SetMenuByJSON error, err=%+v", err)
		util.RenderError(c, err)
		return
	}
	util.RenderSuccess(c, "success")
}
