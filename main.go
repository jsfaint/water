package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/wechat-go/wxweb"
)

type wechat struct {
	session *wxweb.Session
	users   []*wxweb.User
}

var wx wechat

func drinkWater() {
	for _, v := range wx.users {
		wx.session.SendImg("water.jpg", wx.session.Bot.UserName, v.UserName)
	}
}

func afterLoginHandler() error {
	session := wx.session
	opt := parseFlags()

	wx.users = getGroupLists(session.Cm.GetGroupContacts(), opt.group)

	if len(wx.users) == 0 {
		fmt.Println("Get user fail")
		return fmt.Errorf("Invalid user")
	}

	spew.Dump(wx.users)

	cronTable(opt.time)

	return nil
}

func main() {
	var err error
	wx.session, err = wxweb.CreateSession(nil, nil, wxweb.TERMINAL_MODE)
	if err != nil {
		logs.Error(err)
		return
	}

	wx.session.SetAfterLogin(afterLoginHandler)

	if err := wx.session.LoginAndServe(false); err != nil {
		logs.Error(err)
		return
	}
}
