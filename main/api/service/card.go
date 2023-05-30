package service

import (
	"NiuGame/main/Entity"
	"NiuGame/main/common"
	"errors"
)

type enum Entity.Eumn
type card Entity.Card
type cards Entity.Cards

func (s enum) decorToNum() (n int, err error) {
	switch s {
	case common.Club:
		return 1, nil
	case common.Heart:
		return 3, nil
	case common.Spade:
		return 4, nil
	case common.Diamond:
		return 2, nil
	default:
		return 0, errors.New("未知花色")
	}
}

// 重写sort下方法 sort方法重写要求重写以下三个方法 Len  Less  Swap
func (c cards) Len() int {
	return len(c)
}

//func (c cards) Less(i, j int) bool {
//	if c[i].Number != c[j].Number {
//		return c[i].Number < c[j].Number
//	} else {
//		a, _ := c[i].Decor.decorToNum()
//		b, _ := c[j].Decor.decorToNum()
//		return a < b
//	}
//}
//func (c cards) Swap(i, j int) {
//	c[i], c[j] = c[j], c[i]
//}
//
//// 公平性检测
//func (u User) balanceCheck() error {
//	if u.cards.Len() > 5 {
//		return errors.New("Error,Your cards has omre than other one ")
//	}
//	return nil
//}
//
//func (u User) getResult(send cards) (result string, max card) {
//	err := u.balanceCheck()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	sort.Sort(u.cards)
//	//最大卡
//	max = u.cards[4]
//	niu, flag := send.resultCheck()
//	if !flag {
//		fmt.Println("send cards can not be match a niu ")
//	} else {
//		result = "niu" + strconv.Itoa(niu)
//	}
//	return
//}
//func (c cards) resultCheck() (a int, flag bool) {
//	if (c[0].number+c[1].number+c[2].number)%10 != 0 {
//		return 0, false
//	} else {
//		var niu int
//		niu = (c[3].number + c[4].number) % 10
//		return niu, true
//	}
//}
//
//// 发送搭配
//func (u User) sendMatch(sendChan chan<- cards) {
//	sendChan <- u.cards
//}
//
//// 接收
//func (u *User) getCards(getcards cards) {
//	u.cards = getcards
//}
