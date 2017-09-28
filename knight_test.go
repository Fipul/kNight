package knight_test

import (
	. "github.com/Fipul/knight"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//available moves test
var _ = Describe("AvailableMoves", func() {
	Context("если point=a8", func() {
		It("ожидается правильный результат", func() {
			Expect(AvailableMoves("a8")).To(Equal([]string{"b6", "c7"}))
		})
	})
	Context("если point=a1", func() {
		It("ожидается правильный результат", func() {
			Expect(AvailableMoves("a1")).To(Equal([]string{"b3", "c2"}))
		})
	})
	Context("если point=h1", func() {
		It("ожидается правильный результат", func() {
			Expect(AvailableMoves("h1")).To(Equal([]string{"f2", "g3"}))
		})
	})
	Context("если point=h8", func() {
		It("ожидается правильный результат", func() {
			Expect(AvailableMoves("h8")).To(Equal([]string{"f7", "g6"}))
		})
	})
	Context("если point=b4", func() {
		It("ожидается правильный результат", func() {
			Expect(AvailableMoves("b4")).To(Equal([]string{"a2", "a6", "c2", "c6", "d3", "d5"}))
		})
	})
	Context("если point=g5", func() {
		It("ожидается правильный результат", func() {
			Expect(AvailableMoves("g5")).To(Equal([]string{"e4", "e6", "f3", "f7", "h3", "h7"}))
		})
	})
	Context("если point=f5", func() {
		It("ожидается правильный результат", func() {
			Expect(AvailableMoves("f5")).To(Equal([]string{"d4", "d6", "e3", "e7", "g3", "g7", "h4", "h6"}))
		})
	})
	Context("если point не верный", func() {
		It("ожидается пустой ответ", func() {
			_, err := AvailableMoves("a9")
			Expect(err).To(HaveOccurred())
		})
	})
})

//decode tests
var _ = Describe("DecodePoint", func() {
	Context("если point не верный", func() {
		It("ожидается ошибка", func() {
			_, err := DecodePoint("5s")
			Expect(err).To(HaveOccurred())
		})
	})
	Context("если point пустой", func() {
		It("ожидается ошибка", func() {
			_, err := DecodePoint("")
			Expect(err).To(HaveOccurred())
		})
	})
	Context("если point слишком длиный", func() {
		It("ожидается ошибка", func() {
			_, err := DecodePoint("a5r")
			Expect(err).To(HaveOccurred())
		})
	})
	Context(`если point слишком короткий`, func() {
		It("ожидается ошибка", func() {
			_, err := DecodePoint("a")
			Expect(err).To(HaveOccurred())
		})
	})
	Context("если point верный", func() {
		It("ожидается правильный результат", func() {
			Expect(DecodePoint("a1")).To(Equal([2]int{1, 1}))
			Expect(DecodePoint("a8")).To(Equal([2]int{1, 8}))
			Expect(DecodePoint("h1")).To(Equal([2]int{8, 1}))
			Expect(DecodePoint("h8")).To(Equal([2]int{8, 8}))
			Expect(DecodePoint("b2")).To(Equal([2]int{2, 2}))
			Expect(DecodePoint("b7")).To(Equal([2]int{2, 7}))
			Expect(DecodePoint("g2")).To(Equal([2]int{7, 2}))
			Expect(DecodePoint("g7")).To(Equal([2]int{7, 7}))
			Expect(DecodePoint("e4")).To(Equal([2]int{5, 4}))
		})
	})
})

//encode tests
var _ = Describe("EncodePoint", func() {
	Context("если point не верный", func() {
		It("ожидается ошибка", func() {
			_, err := EncodePoint([2]int{0, 8})
			Expect(err).To(HaveOccurred())
			_, err = EncodePoint([2]int{1, 9})
			Expect(err).To(HaveOccurred())
			_, err = EncodePoint([2]int{-2, 14})
			Expect(err).To(HaveOccurred())
		})
	})
	Context("если point верный", func() {
		It("ожидается правильный результат", func() {
			Expect(EncodePoint([2]int{1, 8})).To(Equal("a8"))
			Expect(EncodePoint([2]int{8, 8})).To(Equal("h8"))
			Expect(EncodePoint([2]int{8, 1})).To(Equal("h1"))
			Expect(EncodePoint([2]int{1, 1})).To(Equal("a1"))
			Expect(EncodePoint([2]int{5, 4})).To(Equal("e4"))
		})
	})
})
