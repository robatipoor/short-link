package service

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/robatipoor/short-link/pkg/configs"
	"github.com/robatipoor/short-link/pkg/utils"
	"github.com/robatipoor/short-link/services/key-gen/config"
	"github.com/robatipoor/short-link/services/key-gen/model"
)

var KeysPool chan *model.Key = make(chan *model.Key, config.KeyPoolSize)

func init() {
	s := gocron.NewScheduler()
	s.Every(config.KeyLen).Seconds().Do(generateKeyJob)
	s.Start()
	initKeysPool()
}

func GetKey(exp string) (string, error) {
	ex := configs.ExpireTimeUrl
	var err error
	key := <-KeysPool
	if len(exp) != 0 {
		ex, err = strconv.Atoi(exp)
		if err != nil {
			return "", err
		}
	}
	err = model.ReInsertUsedKey(key, ex)
	if err != nil {
		return "", err
	}
	log.Println("get key : ", key, " exp : ", ex)
	return key.Value, nil
}

func UseKey(k string, exp string) error {
	var err error
	ex := configs.ExpireTimeUrl
	if len(exp) != 0 {
		ex, err = strconv.Atoi(exp)
		if err != nil {
			return err
		}
	}
	if len(k) <= config.KeyLen {
		return errors.New("key invalid")
	}
	key := model.Key{
		Value:      k,
		CreateDate: time.Now(),
	}
	log.Printf("k %s  ex %d", k, ex)
	return model.InsertUsedKey(&key, ex)
}

func DelKey(key string) error {
	k := model.Key{
		Value: key,
	}
	return model.DeleteUsedKey(&k)
}

func generateKeyJob() {
	s := utils.RandomString(config.KeyLen)
	if len(model.FindListUsedKeyByValue(s)) > 1 {
		return
	}
	k := model.NewKey(s)
	err := model.InsertUnusedKey(&k)
	if err != nil {
		log.Println("failed generate key error details : ", err)
	} else {
		log.Println("generate new key : ", k)
	}
}

func initKeysPool() chan bool {
	closed := make(chan bool, 1)
	go func() {
	out:
		for {
			select {
			case <-closed:
				log.Println("init key pool stoped")
				break out
			default:
				key := model.FindListUnUsedKey(1)
				for i := 0; i < len(key); i++ {
					err := model.DeleteUnUsedKey(key[i])
					if err != nil {
						break
					}
					err = model.InsertUsedKey(key[i], configs.ExpireTimeUrl)
					if err != nil {
						break
					}
					KeysPool <- key[i]
					log.Println("consume key to pool : ", key[i].Value)
				}

			}
		}
	}()
	return closed
}
