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
		log.Printf("get new key failed %s \n", err)
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
		log.Println(err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return errors.New("unsuccess")
	}
	return nil
}
