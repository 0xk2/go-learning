package protoexample

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"

	"google.golang.org/protobuf/proto"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStringBytesMaskImpr(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// Create a new Person and save file; then return object and binary data
func NewPerson(name string, age int32, twitterFollower int32, facebookFollower int32) (*Person, []byte) {
	newId := randStringBytesMaskImpr(6)
	p := &Person{
		Id:   newId,
		Name: name,
		Age:  age,
		SocialFollower: &SocialFollowers{
			Twitter:  twitterFollower,
			Facebook: facebookFollower,
		},
	}
	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal("Marshalling error: ", err)
	}
	fmt.Println(data)
	saveFile(newId, data)
	return p, data
}

func saveFile(filename string, data []byte) error {
	location := "exampledata/" + filename
	err := ioutil.WriteFile(location, data, 0755)
	if err != nil {
		log.Fatal("Cannot save file ", err)
	}
	return err
}

func LoadPerson(b []byte) (*Person, error) {
	person := &Person{}
	err := proto.Unmarshal(b, person)
	if err != nil {
		log.Fatal("Unmarshalling error: ", err)
	}
	return person, err
}

func LoadPersonFromFile(filename string) (*Person, error) {
	var p *Person
	data, err := ioutil.ReadFile("exampledata/" + filename)
	if err != nil {
		log.Fatal("Cannot load file exampledata/", filename, "; ", err)
	} else {
		p, err = LoadPerson(data)
	}
	return p, err
}
