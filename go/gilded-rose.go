package main

import "fmt"

type Item struct {
	name            string
	sellIn, quality int
}

var items = []Item{
	Item{"+5 Dexterity Vest", 10, 20},
	Item{"Aged Brie", 2, 0},
	Item{"Elixir of the Mongoose", 5, 7},
	Item{"Sulfuras, Hand of Ragnaros", 0, 80},
	Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	Item{"Conjured Mana Cake", 3, 6},
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

		item.quality = item.quality + 1
		if item.name == "Backstage passes to a TAFKAL80ETC concert" {
			if item.sellIn <= 10 {
				item.quality = item.quality + 1
			}
			if item.sellIn <= 5 {
				item.quality = item.quality + 1
			}
		}
		item.sellIn = item.sellIn - 1
		return
	}

	item.quality = item.quality - 1
	item.sellIn = item.sellIn - 1

	if item.sellIn < 0 {
		item.quality = item.quality - 1
	}
}
