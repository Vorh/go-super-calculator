package main

import (
	"net/http"
	"net/url"
	"os"
	"go-super-calculator/src/utils"
	"src/github.com/op/go-logging"
	"io/ioutil"
	"bufio"
)


var log *logging.Logger

const urlService = "http://localhost:8080/calculate"

func main() {

	listenInput(os.Stdin)


}
func init()  {
	utils.InitLogger()
	log = utils.GetLogger()
}

func listenInput(rd *os.File)  {
	reader := bufio.NewReader(rd)

	for {
		text, _ := reader.ReadString('\n')
		handleInput(text)
	}

}

func handleInput(text string)  {
	sendRequest(text)
}

func sendRequest(ex string) {

	resp, err := http.PostForm(urlService, url.Values{"ex": {ex}})

	if err != nil {
		log.Error(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	log.Debugf("Resp body: %s",body)
}
