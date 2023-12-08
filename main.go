package main

import (
	"fmt"
	"math/rand"
)

type Chest struct {
	Hats []Hat
}

var hats = []Hat{
	{
		Name:   "Flame Crest",
		Rarity: Red,
	},
	{
		Name:   "Sunbeam Halo",
		Rarity: Yellow,
	},
	{
		Name:   "Azure Skyliner",
		Rarity: Blue,
	},
	{
		Name:   "Blush Blossom",
		Rarity: Pink,
	},
	{
		Name:   "Royal Velvet Crown",
		Rarity: Pink,
	},
	{
		Name:   "Ember Blaze Fedora",
		Rarity: Blue,
	},
	{
		Name:   "Gilded Sunshade",
		Rarity: Blue,
	},
	{
		Name:   "Sapphire Twilight Topper",
		Rarity: Blue,
	},
	{
		Name:   "Petal Pink Plume",
		Rarity: Purple,
	},
	{
		Name:   "Majestic Amethyst Helm",
		Rarity: Purple,
	},
}

var caseItems = []CaseItem{
	{Name: "------------------AWP | Dragon Lore", Rarity: 0.0025},
	{Name: "AK-47 | Fire Serpent", Rarity: 0.249375},
	{Name: "M4A4 | Howl", Rarity: 0.249375},
	{Name: "Glock-18 | Fade", Rarity: 0.249375},
	{Name: "USP-S | Kill Confirmed", Rarity: 0.249375},
	{Name: "Desert Eagle | Golden Koi", Rarity: 0.249375},
}

type CaseItem struct {
	Name   string
	Rarity float64
}

func main() {
	// caseHardened := Chest{
	// 	Hats: hats,
	// }
	// totalPercentage := 0.0025 + 0.249375 + 0.249375 + 0.249375 + 0.249375
	totalPercentage := 0.0

	for _, perc := range caseItems {
		totalPercentage += perc.Rarity
	}

	fmt.Println("totalPercentage: ", totalPercentage)
	// for i := 0; i < 10; i++ {
	// }
	// fmt.Println(randomNumber)
	cumulativePercentage := 100.0
	// for i := 0; i < 100; i++ {
	tries := 0
	for cumulativePercentage > 0.249375 {
		tries += 1
		randomNumber := rand.Float64() * totalPercentage
		cumulativePercentage = 0.0
		for _, item := range caseItems {
			cumulativePercentage += item.Rarity
			if randomNumber <= cumulativePercentage {
				// fmt.Println("randomNumber: ", randomNumber)
				// fmt.Println("cumulativePercentage: ", cumulativePercentage)
				fmt.Println("tries: ", tries, "Item: ", item, "cumulative: ", cumulativePercentage, "randomNumber:", randomNumber)
				break
			}
		}
	}
}

// func (c *Chest) OpenChest() {
// 	totalProbability := 0.0

// 	for _, hat := range c.Hats {
// 		totalProbability += hat.Rarity
// 	}
// }
