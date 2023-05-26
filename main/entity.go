package main

// 卡牌
type eumn string

type card struct {
	number int
	decor  eumn
}
type cards []card

type User struct {
	userId int
	cards  cards
}

type Result struct {
	message string
	max     card
}
