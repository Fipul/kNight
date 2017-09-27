package knight_test

import (
	. "github.com/Fipul/knight"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AvaliableMoves", func() {
	Context("если point=a8", func() {
		It("ожидается правильный результат", func() {
			Expect(AvaliableMoves("a8")).To(Equal([]string{"b6", "c7"}))
		})
	})
	Context("если point=a1", func() {
		It("ожидается правильный результат", func() {
			Expect(AvaliableMoves("a1")).To(Equal([]string{"b3", "c2"}))
		})
	})
	Context("если point=h1", func() {
		It("ожидается правильный результат", func() {
			Expect(AvaliableMoves("h1")).To(Equal([]string{"f2", "g3"}))
		})
	})
	Context("если point=h8", func() {
		It("ожидается правильный результат", func() {
			Expect(AvaliableMoves("h8")).To(Equal([]string{"f7", "g6"}))
		})
	})
	Context("если point=b4", func() {
		It("ожидается правильный результат", func() {
			Expect(AvaliableMoves("b4")).To(Equal([]string{"a2", "a6", "c2", "c6", "d3", "d5"}))
		})
	})
	Context("если point=g5", func() {
		It("ожидается правильный результат", func() {
			Expect(AvaliableMoves("g5")).To(Equal([]string{"e4", "e6", "f3", "f7", "h3", "h7"}))
		})
	})
	Context("если point=f5", func() {
		It("ожидается правильный результат", func() {
			Expect(AvaliableMoves("f5")).To(Equal([]string{"d4", "d6", "e3", "e7", "g3", "g7", "h4", "h6"}))
		})
	})
	Context("если point не верный", func() {
		It("ожидается пустой ответ", func() {
			Expect(AvaliableMoves("a9")).To(Equal([]string{}))
		})
	})
})

var _ = Describe("DecodePoint", func() {
	Context("если point не верный", func() {
		It("ожидается [0 0]", func() {
			Expect(DecodePoint("")).To(Equal([2]int{0, 0}))
		})
	})
})
