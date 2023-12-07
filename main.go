package main

import "fmt"

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
		Name:   " Azure Skyliner",
		Rarity: Blue,
	},
	{
		Name:   "Blush Blossom",
		Rarity: Pink,
	},
	{
		Name:   "Royal Velvet Crown",
		Rarity: Purple,
	},
	{
		Name:   "Ember Blaze Fedora",
		Rarity: Red,
	},
	{
		Name:   "Gilded Sunshade",
		Rarity: Yellow,
	},
	{
		Name:   "Sapphire Twilight Topper",
		Rarity: Blue,
	},
	{
		Name:   "Petal Pink Plume",
		Rarity: Pink,
	},
	{
		Name:   "Majestic Amethyst Helm",
		Rarity: Purple,
	},
}

func main() {
	caseHardened := Chest{
		Hats: hats,
	}

	fmt.Println(caseHardened)
}
