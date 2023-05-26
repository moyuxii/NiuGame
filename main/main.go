package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

const (
	club    = "club"
	Heart   = "Heart"
	Spade   = "Spade"
	Diamond = "Diamond"
)

func (s eumn) decorToNum() (n int, err error) {
	switch s {
	case club:
		return 1, nil
	case Heart:
		return 3, nil
	case Spade:
		return 4, nil
	case Diamond:
		return 2, nil
	default:
		return 0, errors.New("未知花色")
	}
}

// 重写sort下方法 sort方法重写要求重写以下三个方法 Len  Less  Swap
func (c cards) Len() int {
	return len(c)
}
func (c cards) Less(i, j int) bool {
	if c[i].number != c[j].number {
		return c[i].number < c[j].number
	} else {
		a, _ := c[i].decor.decorToNum()
		b, _ := c[j].decor.decorToNum()
		return a < b
	}
}
func (c cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// 公平性检测
func (u User) balanceCheck() error {
	if u.cards.Len() > 5 {
		return errors.New("Error,Your cards has omre than other one ")
	}
	return nil
}

func (u User) getResult(send cards) (result string, max card) {
	err := u.balanceCheck()
	if err != nil {
		fmt.Println(err)
		return
	}
	sort.Sort(u.cards)
	//最大卡
	max = u.cards[4]
	niu, flag := send.resultCheck()
	if !flag {
		fmt.Println("send cards can not be match a niu ")
	} else {
		result = "niu" + strconv.Itoa(niu)
	}
	return
}
func (c cards) resultCheck() (a int, flag bool) {
	if (c[0].number+c[1].number+c[2].number)%10 != 0 {
		return 0, false
	} else {
		var niu int
		niu = (c[3].number + c[4].number) % 10
		return niu, true
	}
}

// 发送搭配
func (u User) sendMatch(sendChan chan<- cards) {
	sendChan <- u.cards
}

// 接收
func (u *User) getCards(getcards cards) {
	u.cards = getcards
}

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
	chan1 := make(chan cards, 1)
	chan2 := make(chan cards)
	card1 := cards{{1, "club"},
		{2, "Diamond"},
		{3, "Heart"},
		{4, "Spade"},
		{5, "Diamond"}}
	chan1 <- cards{{1, "club"},
		{2, "Diamond"},
		{3, "Heart"},
		{4, "Spade"},
		{5, "Diamond"}}
	go func() {
		var user1 User
		user1.userId = 1
		user1.getCards(<-chan1)
		user1.cards = user1.cards[:4]
		user1.cards = append(user1.cards, card{2, "Diamond"})
		user1.sendMatch(chan2)
	}()
	close(chan1)
	card2 := <-chan2
	close(chan2)
	fmt.Println(card2)
	if card1.Len() == card2.Len() {
		sort.Sort(card1)
		sort.Sort(card2)
		for i := 0; i < card1.Len(); i++ {
			if card1[i] != card2[i] {
				fmt.Println("!!!!")
				return
			}
		}
		fmt.Println("Yes!!")
	} else {
		fmt.Println("!!!!")
	}
}
