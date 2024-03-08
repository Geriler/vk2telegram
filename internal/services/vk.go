package services

import (
	g "gopkg.in/h2non/gentleman.v2"
	"strconv"
	"vk2telegram/internal/model/messenger"
)

type VKAPI struct {
	token   string
	url     string
	version string
}

func NewVKAPI(token string, url string, version string) *VKAPI {
	return &VKAPI{
		token:   token,
		url:     url,
		version: version,
	}
}

func (vk *VKAPI) GroupsSearch(search string) ([]messenger.Group, error) {
	resp, err := g.New().
		URL(vk.url + "/groups.search?q=" + search + "&access_token=" + vk.token + "&v=" + vk.version).
		Request().
		Send()

	if err != nil {
		return nil, err
	}

	var wrapper messenger.WrapperGroup
	err = resp.JSON(&wrapper)
	if err != nil {
		return nil, err
	}

	return wrapper.Response.Items, nil
}

func (vk *VKAPI) WallGet(group messenger.Group, count int) ([]messenger.Post, error) {
	resp, err := g.New().
		URL(vk.url + "/wall.get?domain=" + group.ScreenName + "&access_token=" + vk.token + "&v=" + vk.version + "&count=" + strconv.Itoa(count)).
		Request().
		Send()

	if err != nil {
		return nil, err
	}

	var wrapper messenger.WrapperPost
	err = resp.JSON(&wrapper)
	if err != nil {
		return nil, err
	}

	return wrapper.Response.Items, nil
}
