package decisiontree

import (
	"bytes"
	"log"
	"net/http"
)

func tweet(text string) bool {
	endPoint := "https://api.twitter.com/2/tweets"
	auth := "T3ZvMlZOUUllVjVVWHdEM0puM1NnWm1WR0R6RE9lckt6SnVnSm5vNzRyazVUOjE2ODAyNDY3NjU4Mjk6MToxOmF0OjE"
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
