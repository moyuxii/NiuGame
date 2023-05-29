package main

import (
	"NiuGame/main/Auth"
	"NiuGame/main/Config"
	"github.com/gin-gonic/gin"
)

func main() {
	//var username string
	//
	//u := User{
	//	userId: 1,
	//	cards: cards{{1, "club"},
	//		{2, "Diamond"},
	//		{3, "Heart"},
	//		{4, "Spade"},
	//		{5, "Diamond"}},
	//}
	//fmt.Print("Hello,Please input your name: ")
	//_, err := fmt.Scanln(&username)
	//if err != nil {
	//	return
	//}
	//fmt.Println("Hello,", username)
	//fmt.Println("Your cards is  \n", u)
	//send := cards{{1, "club"},
	//	{5, "Diamond"},
	//	{4, "Heart"},
	//	{2, "Spade"}
	//	{3, "Diamond"}}
	//result, max := u.getResult(send)
	//if result != "" {
	//	fmt.Println(result, max)
	//}
	//chan1 := make(chan cards, 1)
	//chan2 := make(chan cards)
	//card1 := cards{{1, "club"},
	//	{2, "Diamond"},
	//	{3, "Heart"},
	//	{4, "Spade"},
	//	{5, "Diamond"}}
	//chan1 <- cards{{1, "club"},
	//	{2, "Diamond"},
	//	{3, "Heart"},
	//	{4, "Spade"},
	//	{5, "Diamond"}}
	//go func() {
	//	var user1 User
	//	user1.userId = 1
	//	user1.getCards(<-chan1)
	//	user1.cards = user1.cards[:4]
	//	user1.cards = append(user1.cards, card{2, "Diamond"})
	//	user1.sendMatch(chan2)
	//}()
	//close(chan1)
	//card2 := <-chan2
	//close(chan2)
	//fmt.Println(card2)
	//if card1.Len() == card2.Len() {
	//	sort.Sort(card1)
	//	sort.Sort(card2)
	//	for i := 0; i < card1.Len(); i++ {
	//		if card1[i] != card2[i] {
	//			fmt.Println("!!!!")
	//			return
	//		}
	//	}
	//	fmt.Println("Yes!!")
	//} else {
	//	fmt.Println("!!!!")
	//}
	_, err := Config.ParseConfig("E:\\Project\\NiuGame\\main\\Config\\app.json")
	if err != nil {
		panic("读取配置文件失败，" + err.Error())
	}

	r := gin.Default()
	r.Use(Auth.JWTAuth())
	// 用户登录接口
	r.POST("/user/login", Login)
	r.GET("/getCards", getCards)
	r.Run()
}
