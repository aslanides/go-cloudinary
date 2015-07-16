package main

import(	"github.com/go-cloudinary"
		"encoding/json"
		"fmt"
		"io/ioutil"
		"log"
)

type CloudinaryConfig struct {
	Key string `json:"cloudinary_api_key"`
	Secret string `json:"cloudinary_secret"`
	Preset string `json:"cloudinary_upload_preset"`
	CloudName string `json:"cloudinary_cloudname"`
}

func main() {
	config := "../config.json"
	databuf, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatalf("While reading config file: %+v", err)
		return
	}
	
	var c CloudinaryConfig
	if err := json.Unmarshal(databuf, &c); err != nil {
		log.Fatalf("While unmarshal of config: %+v", err)
		return
	}

	uri := fmt.Sprintf("cloudinary://%s:%s@%s", c.Key, c.Secret, c.CloudName)
	s, err := cloudinary.Dial(uri)

	s.SetFolder("")
	if _, err := s.Upload("wat.jpg", nil, "",true,cloudinary.ImageType); err != nil {
		log.Fatalf("While uploading: %+v",err)
		return
	}
	
	// JWT shit
	/* 

	myKey := "efoAl6Nge5U68Kz0Kpj4osiX6BE"
	byteKey := []byte(myKey)
	token := jwt.New(jwt.SigningMethodHS256)
	
	token.Claims["foo"] = "bar"
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	
	tokenString, err := token.SignedString(byteKey)
	if err != nil {
		fmt.Println("Error:",err)
	} else {
		fmt.Println(tokenString)	
	}
	*/
}