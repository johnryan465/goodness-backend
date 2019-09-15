package controllers

import (
	"net/http"
	"encoding/json"
	b64	"encoding/base64"
	u "goodness/utils"
	"os"
	"fmt"
	"bytes"
	"io/ioutil"
	"encoding/xml"
)

var tokenHost = "https://13/apis/v1.0.1/oauth"
var authorizePath = "/authorize"
type Credentials struct {
	user_id	string
	redirect_uri	string
	user_secret	string
}


type Payload struct {
	user_id	string
	redirect_uri	string
	query_hash	string
}

var GetToken = func(w http.ResponseWriter, r *http.Request) {
	credentials := GetCredentials()
	redirect_uri := r.URL.Query()["redirect_uri"][0]
	code := r.URL.Query()["code"][0]
	url := "https://apis-bank-test.apigee.net/apis/v1.0.1/oauth/token"


	var bodyParams, _= json.Marshal(`{
									"redirect_uri" :` + redirect_uri + `,
									"grant_type" : "authorization_code",
									"code" : ` + code + `}`)
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyParams))
  req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", credentials.query_hash)

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      panic(err)
  }
  defer resp.Body.Close()

  fmt.Println("response Status:", resp.Status)
  fmt.Println("response Headers:", resp.Header)
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println("response Body:", string(body))
}

var GetCredentials = func()(*Payload) {
	var data Credentials
	jsonFile, _ := os.Open("data/credentials.json")
	defer jsonFile.Close()

	jsonParser := json.NewDecoder(jsonFile)
	jsonParser.Decode(&data)

	cred := data.user_id + ":" + data.user_secret
	payload := &Payload{
		user_id: data.user_id,
		redirect_uri: data.redirect_uri,
		query_hash: b64.StdEncoding.EncodeToString([]byte(cred)),
	}
	return payload
}

var GetOAuthToken = func(w http.ResponseWriter, r *http.Request) {
	uri := tokenHost + authorizePath
	resp, err := http.Get(uri)
	if err != nil {
		return //errors.Wrap(err, "Berkeley query failed")
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
  buf.ReadFrom(resp.Body)
  res:= buf.Bytes()

	v := Result{}
	err = xml.Unmarshal(res, &v)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	var data Credentials
	jsonFile, err := os.Open("data/credentials.json")
	defer jsonFile.Close()
	if err != nil {
		u.Respond(w, u.Message(false, "Error while parsing JSON credentials file"))
		return
	}
	jsonParser := json.NewDecoder(jsonFile)
	jsonParser.Decode(&data)

	cred := data.user_id + ":" + data.user_secret
	payload := &Payload{
		user_id: data.user_id,
		redirect_uri: data.redirect_uri,
		query_hash: b64.StdEncoding.EncodeToString([]byte(cred)),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}
