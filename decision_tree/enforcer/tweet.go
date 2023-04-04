package enforcer

import (
	"bytes"
	"log"
	"net/http"
)

func Tweet(text string) bool {
	endPoint := "https://api.twitter.com/2/tweets"
	auth := "RHNFZnRtaEtGS2JxN3BuNVNad2dfSEduZHQ3MTl4TVpfdjVEOWE4czc0YVJ1OjE2ODA1MjMyOTA2MjU6MToxOmFjOjE"
	var jsonStr = []byte(`{"text":"` + text + `"}`)
	request, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("Error creating request: ", err)
	}
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Authorization", "Bearer "+auth)
	request.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, e := client.Do(request)
	if e != nil {
		log.Println("Error sending request: ", e)
	} else {
		log.Println(resp)
		log.Println("Tweet sent: " + text)
		return true
	}
	defer resp.Body.Close()
	return false
}
