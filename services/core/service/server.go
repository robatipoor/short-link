package service

import (
	"errors"
	"log"
	"time"

	lru "github.com/hashicorp/golang-lru"
	"github.com/robatipoor/short-link/pkg/configs"
	"github.com/robatipoor/short-link/services/core/client"
	"github.com/robatipoor/short-link/services/core/config"
	"github.com/robatipoor/short-link/services/core/domain/request"
	"github.com/robatipoor/short-link/services/core/domain/response"
	"github.com/robatipoor/short-link/services/core/model"
)

var cache *lru.Cache

func init() {
	var err error
	cache, err = lru.New(100)
	if err != nil {
		log.Println(err)
	}
}

func Redirect(req *request.Redirect) (*response.Redirect, error) {
	if url, ok := cache.Get(req.ShortUrl); ok {
		log.Println("find in cache key ", req.ShortUrl)
		r := &response.Redirect{
			Link: url.(string),
		}
		return r, nil
	}
	url, err := model.FindUrlByShortUrl(req.ShortUrl)
	if err != nil {
		return nil, err
	}
	cache.Add(url.ShortUrl, url.OriginalUrl)
	resp := response.Redirect{
		Link: url.OriginalUrl,
	}
	return &resp, nil
}

func CreateNewLink(req *request.CreateNewUrl) (*response.CreateNewUrl, error) {
	key, err := client.GetNewKey(req.ExpireTime)
	if err != nil {
		return nil, err
	}
	var exp int
	if req.ExpireTime != 0 {
		exp = req.ExpireTime
	} else {
		exp = configs.ExpireTimeUrl
	}
	url := model.Url{
		OriginalUrl: req.OriginalUrl,
		CreateDate:  time.Now(),
		ShortUrl:    key,
	}
	err = model.InsertUrl(&url, exp)
	if err != nil {
		log.Println("failed insert new url ", err.Error())
		return nil, err
	}
	resp := response.CreateNewUrl{}
	resp.Link = config.Http + config.ServerAddr + ":" + config.ServerPort + "/" + key
	return &resp, nil
}

func DeleteLink(key string) error {
	err := model.DeleteUrlByShortUrl(key)
	if err != nil {
		log.Println("detele link failed details : ", err.Error())
		return err
	}
	err = client.DeleteKey(key)
	if err != nil {
		log.Println("detele link keygen failed details : ", err.Error())
		return err
	}
	if !cache.Remove(key) {
		return errors.New("failed delete key " + key)
	}
	return nil
}
