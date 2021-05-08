package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ Json Data Struct ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
type Insult struct {
	Active    json.Number `json:"active"`
	Comment   string      `json:"comment"`
	Created   CustomTime  `json:"created"`
	Createdby string      `json:"createdby"`
	Insult    string      `json:"insult"`
	Language  string      `json:"language"`
	Number    json.Number `json:"number"`
	Shown     json.Number `json:"shown"`
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ Time format override ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ Time format override ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// in order to use a differnt time format in json, you need to overwrite the time object in the marshal and unmarshal.
// https://stackoverflow.com/questions/25087960/json-unmarshal-time-that-isnt-in-rfc-3339-format
const ctLayout = "2006-01-02 15:04:05"

var nilTime = (time.Time{}).UnixNano()

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ Time format override ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

func main() {
	url := "https://evilinsult.com/generate_insult.php?lang=en&type=json"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	//req.Header.Set("User-Agent", "Go Harvest API Sample")
	//req.Header.Set("Harvest-Account-ID", os.Getenv("HARVEST_ACCOUNT_ID"))
	//req.Header.Set("Authorization", "Bearer "+os.Getenv("HARVEST_ACCESS_TOKEN"))

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	//var jsonResponse map[string]interface{}
	//json.Unmarshal(body, &jsonResponse)
	//prettyJson, _ := json.MarshalIndent(jsonResponse, "", "  ")
	//fmt.Println("prettyJson", string(prettyJson))
	//////////////////////////////////////////////

	// read values straight from json object as strings.
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	fmt.Println("unstructured:", result["insult"])

	// store json data as struct and read with proper object types.
	jokes := Insult{}
	json.Unmarshal(body, &jokes)
	fmt.Println("structured insult", jokes.Insult)
	fmt.Println("structured created:", jokes.Created.Format("2006-01-02 15:04:05"))

}
