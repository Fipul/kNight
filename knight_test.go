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
			Expect(AvailableMoves("a8")).To(Equal([8]string{"b6", "c7"}))
		})
	})
	Context("если point=a1", func() {
		It("ожидается правильный результат", func() {
			Expect(AvailableMoves("a1")).To(Equal([8]string{"b3", "c2"}))
		})
	})
	Context("если point=h1", func() {
		It("ожидается правильный результат", func() {
			Expect(AvailableMoves("h1")).To(Equal([8]string{"f2", "g3"}))
		})
	})
	Context("если point=h8", func() {
		It("ожидается правильный результат", func() {
			Expect(AvailableMoves("h8")).To(Equal([8]string{"f7", "g6"}))
		})
	})
	Context("если point=b4", func() {
		It("ожидается правильный результат", func() {
			Expect(AvailableMoves("b4")).To(Equal([8]string{"a2", "a6", "c2", "c6", "d3", "d5"}))
		})
	})
	Context("если point=g5", func() {
		It("ожидается правильный результат", func() {
			Expect(AvailableMoves("g5")).To(Equal([8]string{"e4", "e6", "f3", "f7", "h3", "h7"}))
		})
	})
	Context("если point=f5", func() {
		It("ожидается правильный результат", func() {
			Expect(AvailableMoves("f5")).To(Equal([8]string{"d4", "d6", "e3", "e7", "g3", "g7", "h4", "h6"}))
		})
	})
	Context("если point не верный", func() {
		It("ожидается пустой ответ", func() {
			Expect(AvailableMoves("a9")).To(Equal([8]string{}))
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
	Context("если point = c2", func() {
		It("ожидается [3 2]", func() {
			Expect(DecodePoint("c2")).To(Equal([2]int{3, 2}))
		})
	})
})

//encode tests
var _ = Describe("EncodePoint", func() {
	Context(`если point слишком короткий`, func() {
		It("ожидается ошибка", func() {
			_, err := DecodePoint("a")
			Expect(err).To(HaveOccurred())
		})
	})
	Context("если point слишком длиный", func() {
		It("ожидается ошибка", func() {
			_, err := DecodePoint("a5r")
			Expect(err).To(HaveOccurred())
		})
	})
	Context("если point пустой", func() {
		It("ожидается ошибка", func() {
			_, err := DecodePoint("")
			Expect(err).To(HaveOccurred())
		})
	})
	Context("если point не верный", func() {
		It("ожидается ошибка", func() {
			_, err := DecodePoint("5s")
			Expect(err).To(HaveOccurred())
		})
	})
	Context("если point = [1,8]", func() {
		It("ожидается [a8]", func() {
			Expect(EncodePoint([2]int{1, 8})).To(Equal("a8"))
		})
	})
})
