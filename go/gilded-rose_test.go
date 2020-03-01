package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	It("should have items", func() {
		Expect(len(items)).To(BeNumerically(">", 0))
	})

	It("should lower the quality each day", func() {
		testItems := []Item{
			Item{"Test Item", 10, 20},
		}
		GildedRose(testItems)
		Expect(testItems[0].quality).To(Equal(19))
	})
})
