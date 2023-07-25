package oa

import (
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/neocxf/oa/pkg/config"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/officialaccount"

	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	log "github.com/sirupsen/logrus"
)

// Account 公众号操作样例
type Account struct {
	wc              *wechat.Wechat
	officialAccount *officialaccount.OfficialAccount
}

// ExampleOfficialAccount new
func NewAccount(wc *wechat.Wechat) *Account {
	//init config
	wechatCfg := config.GetConfig()
	offCfg := &offConfig.Config{
		AppID:          wechatCfg.AppID,
		AppSecret:      wechatCfg.AppSecret,
		Token:          wechatCfg.Token,
		EncodingAESKey: wechatCfg.EncodingAESKey,
	}
	log.Debugf("offCfg=%+v", offCfg)
	officialAccount := wc.GetOfficialAccount(offCfg)
	return &Account{
		wc:              wc,
		officialAccount: officialAccount,
	}
}

// Serve 处理消息
func (ex *Account) Serve(c *fiber.Ctx) (err error) {
	r, _ := adaptor.ConvertRequest(c, false)

	server := ex.officialAccount.GetServer(r.WithContext(c.Context()), &netHTTPResponseWriter{w: c.Response().BodyWriter()})
	server.SkipValidate(true)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		//TODO
		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		//article1 := message.NewArticle("测试图文1", "图文描述", "", "")
		//articles := []*message.Article{article1}
		//news := message.NewNews(articles)
		//return &message.Reply{MsgType: message.MsgTypeNews, MsgData: news}

		//voice := message.NewVoice(mediaID)
		//return &message.Reply{MsgType: message.MsgTypeVoice, MsgData: voice}

		//
		//video := message.NewVideo(mediaID, "标题", "描述")
		//return &message.Reply{MsgType: message.MsgTypeVideo, MsgData: video}

		//music := message.NewMusic("标题", "描述", "音乐链接", "HQMusicUrl", "缩略图的媒体id")
		//return &message.Reply{MsgType: message.MsgTypeMusic, MsgData: music}

		//多客服消息转发
		//transferCustomer := message.NewTransferCustomer("")
		//return &message.Reply{MsgType: message.MsgTypeTransfer, MsgData: transferCustomer}
	})

	//处理消息接收以及回复
	err = server.Serve()
	if err != nil {
		log.Errorf("Serve Error, err=%+v", err)
		return
	}
	//发送回复的消息
	err = server.Send()
	if err != nil {
		log.Errorf("Send Error, err=%+v", err)
		return
	}
	return nil
}

type netHTTPResponseWriter struct {
	statusCode int
	h          http.Header
	w          io.Writer
}

func (w *netHTTPResponseWriter) StatusCode() int {
	if w.statusCode == 0 {
		return http.StatusOK
	}
	return w.statusCode
}

func (w *netHTTPResponseWriter) Header() http.Header {
	if w.h == nil {
		w.h = make(http.Header)
	}
	return w.h
}

func (w *netHTTPResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (w *netHTTPResponseWriter) Write(p []byte) (int, error) {
	return w.w.Write(p)
}

func (w *netHTTPResponseWriter) Flush() {}
