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

func main() {
	chest1 := Chest{
		Hats: hats,
	}

	chest1.OpenChest()
}

func (c *Chest) OpenChest() {
	totalPercentage := 0.0

	for _, perc := range c.Hats {
		totalPercentage += float64(perc.Rarity)
	}

	randomNumber := rand.Float64() * totalPercentage

	cumulativePercentage := 0.0

	for _, item := range c.Hats {
		cumulativePercentage += float64(item.Rarity)
		if randomNumber <= cumulativePercentage {
			// fmt.Println("randomNumber: ", randomNumber)
			// fmt.Println("cumulativePercentage: ", cumulativePercentage)
			fmt.Println("Item: ", item, "cumulative: ", cumulativePercentage, "randomNumber:", randomNumber)
			break
		}
	}
}
