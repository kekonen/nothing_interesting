package main

import "github.com/tidwall/gjson"

// const container = `
// {
// 	"container":{
// 		"first":"Janet",
// 		"last":"Prichard"
// 	},
// 	"age":47
// }`

const container = `
{
	"airbnb": [
		{
			"name": "DE_BE_001",
			"email": "aaa@aaa.com",
			"password": "1234",
			"session": "qwerty"
		},
		{
			"name": "DE_BE_002",
			"email": "bbb@bbb.com",
			"password": "5678",
			"session": "zxcvb"
		}
	],
	"booking": {
		"email": "kek@booking.com",
		"password": "iamthepass",
		"session": "howehfioefwe"
	},
}`

func main() {

	println(gjson.Get(container, "airbnb.#.name").String())
	println(gjson.Get(container, "airbnb.#.session").String())
}
