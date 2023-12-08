package main

type Rarity string

const (
	Red    Rarity = "Red"
	Yellow Rarity = "Yellow"
	Blue   Rarity = "Blue"
	Pink   Rarity = "Pink"
	Purple Rarity = "Purple"
)

type RarityType struct {
	Rarity     Rarity
	Percentage float64
}

type Hat struct {
	Name   string
	Rarity RarityType
}

/*
Yellow  0.25575
Red     0.63940
Pink	3.19693
Purple	15.98465
Blue 	79.92327
*/
