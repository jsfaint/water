package main

import (
	"github.com/songtianyi/wechat-go/wxweb"
)

func getGroupLists(origLists []*wxweb.User, name string) (result []*wxweb.User) {
	users := getLineFromFile(name)
	for _, group := range origLists {
		for _, nick := range users {
			if group.NickName == nick {
				result = append(result, group)
			}
		}
	}

	return result
}
