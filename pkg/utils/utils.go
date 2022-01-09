package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ParseBody(req *http.Request, x interface{}) {
	defer func() {
		if err := req.Body.Close(); err != nil {
			log.Fatalln("Error closing body: ", err)
		}
	}()
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Fatalln("Error reading body: ", err)
	}

	if err := json.Unmarshal([]byte(body), &x); err != nil{
		log.Fatalln("Error unmarshal: ", err)
	}

}