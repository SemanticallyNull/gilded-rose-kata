package inventory_test

import (
	. "code.katiechapman.ie/gilded-rose-kata/go/inventory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	It("should lower the quality each day", func() {
		testItem := Item{"Test Item", 10, 20}
		testItem.UpdateItem()
		Expect(testItem.Quality).To(Equal(19))
	})

	It("should lower the sell in each day", func() {
		testItem := Item{"Test Item", 10, 20}
		testItem.UpdateItem()
		Expect(testItem.SellIn).To(Equal(9))
	})

	It("should lower the quality by 2 when sellin day passed", func() {
		testItem := Item{"Test Item", 0, 20}
		testItem.UpdateItem()
		Expect(testItem.Quality).To(Equal(18))
	})

	It("should not give items a negative quality", func() {
		testItem := Item{"Test Item", 0, 0}
		testItem.UpdateItem()
		Expect(testItem.Quality).To(Equal(0))
	})

	Context("Aged Brie", func() {
		It("should increase the quality of Aged Brie as it ages", func() {
			testItem := Item{"Aged Brie", 2, 20}
			testItem.UpdateItem()
			Expect(testItem.Quality).To(Equal(21))
		})

		It("should decrease the sellIn time of Aged Brie as it ages", func() {
			testItem := Item{"Aged Brie", 2, 20}
			testItem.UpdateItem()
			Expect(testItem.SellIn).To(Equal(1))
		})
	})

	It("should not increase the quality of an item above 50", func() {
		testItem := Item{"Aged Brie", 2, 50}
		testItem.UpdateItem()
		Expect(testItem.Quality).To(Equal(50))
	})

	It("does not have to sell Sulfuras or decrease in quality", func() {
		testItem := Item{"Sulfuras, Hand of Ragnaros", 0, 80}
		testItem.UpdateItem()
		Expect(testItem.SellIn).To(Equal(0))
		Expect(testItem.Quality).To(Equal(80))
	})

	Context("Backstage Passes", func() {
		It("increases in quality as sellin day approaches", func() {
			testItem := Item{"Backstage passes to a TAFKAL80ETC concert", 20, 5}
			testItem.UpdateItem()
			Expect(testItem.Quality).To(Equal(6))
		})

		It("increases in quality by 2 when as sellin day is 10 or less", func() {
			testItem := Item{"Backstage passes to a TAFKAL80ETC concert", 6, 5}
			testItem.UpdateItem()
			Expect(testItem.Quality).To(Equal(7))
		})

		It("increases in quality by 3 when as sellin day is 5 or less", func() {
			testItem := Item{"Backstage passes to a TAFKAL80ETC concert", 4, 10}
			testItem.UpdateItem()
			Expect(testItem.Quality).To(Equal(13))
		})

		It("drops quality to 0 after the concert", func() {
			testItem := Item{"Backstage passes to a TAFKAL80ETC concert", 0, 10}
			testItem.UpdateItem()
			Expect(testItem.Quality).To(Equal(0))
		})
	})

	It("should degrade quality of conjured items twice as fast as regular items", func() {
		testItem := Item{"Conjured Mana Cake", 4, 10}
		testItem.UpdateItem()
		Expect(testItem.Quality).To(Equal(8))
	})

	It("can update multiple items", func() {
		testItems := Items{
			{"Conjured Mana Cake", 4, 10},
			{"Test Item", 10, 20},
		}
		testItems.UpdateItems()
		Expect(testItems[0].SellIn).To(Equal(3))
		Expect(testItems[0].Quality).To(Equal(8))
		Expect(testItems[1].SellIn).To(Equal(9))
		Expect(testItems[1].Quality).To(Equal(19))
	})
})
