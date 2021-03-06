package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

func main() {
	// user, album, photo, todos

	resp, err := http.Get("http://jsonplaceholder.typicode.com/users/1/")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	/*body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}*/

	var t Users
	/*err = json.Unmarshal(body, &t)
	if err != nil {
		fmt.Println(err)
		return
	}*/
	json.NewDecoder(resp.Body).Decode(&t)
	fmt.Printf("% #v\n", t)
}
