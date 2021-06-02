package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var(
url2 = "https://jsonplaceholder.typicode.com/posts"
)

func main()  {
	fmt.Println("hey there")
	for i:=1; i<10; i++ {

		go getUrlmy(i)
	}
	var input string
	fmt.Scanln(&input)

}

func getUrlmy(j int){

	strk := fmt.Sprint(j)
	url2 = "https://jsonplaceholder.typicode.com/posts/" + strk

	res, err := http.Get(url2)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println()
	bodyString := string(body)
	fmt.Print(bodyString)
	//amt := time.Duration(rand.Intn(500))
	//time.Sleep(time.Millisecond * amt)
}


