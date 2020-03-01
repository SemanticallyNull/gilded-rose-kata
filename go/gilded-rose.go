package main

import (
	"fmt"
	"strings"
)

type Item struct {
	name            string
	sellIn, quality int
}

var items = []Item{
	{"+5 Dexterity Vest", 10, 20},
	{"Aged Brie", 2, 0},
	{"Elixir of the Mongoose", 5, 7},
	{"Sulfuras, Hand of Ragnaros", 0, 80},
	{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	{"Conjured Mana Cake", 3, 6},
}

func main() {
	fmt.Println("OMGHAI!")
	// fmt.Print(items)
	UpdateItems(items)
}

func UpdateItems(items []Item) {
	for i := 0; i < len(items); i++ {
		UpdateItem(&items[i])
	}

}

func UpdateItem(item *Item) {
	qualityReductionFactor := 1

	if item.quality >= 50 {
		return
	}

	if item.quality <= 0 {
		return
	}

	if item.name == "Aged Brie" {
		item.quality = item.quality + 1
		item.sellIn = item.sellIn - 1
		return
	}

	if item.name == "Backstage passes to a TAFKAL80ETC concert" {
		if item.sellIn <= 0 {
			item.quality = 0
			return
		}

		if item.sellIn <= 5 {
			item.quality = item.quality + 3
		} else if item.sellIn <= 10 {
			item.quality = item.quality + 2
		} else {
			item.quality = item.quality + 1
		}

		item.sellIn = item.sellIn - 1
		return
	}

	if strings.Contains(item.name, "Conjured") {
		qualityReductionFactor = 2
	}

	item.quality = item.quality - (1 * qualityReductionFactor)
	item.sellIn = item.sellIn - 1

	if item.sellIn < 0 {
		item.quality = item.quality - (1 * qualityReductionFactor)
	}
}
