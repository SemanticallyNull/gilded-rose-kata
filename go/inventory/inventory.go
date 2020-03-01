package inventory

import "strings"

type Items []Item

type Item struct {
	Name            string
	SellIn, Quality int
}

func (items Items) UpdateItems() {
	for i := 0; i < len(items); i++ {
		items[i].UpdateItem()
	}
}

func (item *Item) UpdateItem() {
	qualityReductionFactor := 1

	if item.Quality >= 50 {
		return
	}

	if item.Quality <= 0 {
		return
	}

	if item.Name == "Aged Brie" {
		item.Quality = item.Quality + 1
		item.SellIn = item.SellIn - 1
		return
	}

	if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		if item.SellIn <= 0 {
			item.Quality = 0
			return
		}

		if item.SellIn <= 5 {
			item.Quality = item.Quality + 3
		} else if item.SellIn <= 10 {
			item.Quality = item.Quality + 2
		} else {
			item.Quality = item.Quality + 1
		}

		item.SellIn = item.SellIn - 1
		return
	}

	if strings.Contains(item.Name, "Conjured") {
		qualityReductionFactor = 2
	}

	item.Quality = item.Quality - (1 * qualityReductionFactor)
	item.SellIn = item.SellIn - 1

	if item.SellIn < 0 {
		item.Quality = item.Quality - (1 * qualityReductionFactor)
	}
}
