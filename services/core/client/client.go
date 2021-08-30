package client

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/robatipoor/short-link/services/key-gen/config"
)

func GetNewKey(exp int) (string, error) {
	resp, err := http.Get("http://" + config.ServerAddr + ":" + config.ServerPort + "/getkey?exp=" + strconv.Itoa(exp))
	if err != nil {
		log.Println("get new key failed error details ", err)
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func DeleteKey(key string) error {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://"+config.ServerAddr+":"+config.ServerPort+"/"+key, nil)
	if err != nil {
		log.Println("create new request for delete key failed")
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("call keygen failed")
		return err
	}
	log.Println("response delete key keygen status code ", resp.StatusCode)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("delete key in keygen service un success")
	}
	return nil
}
