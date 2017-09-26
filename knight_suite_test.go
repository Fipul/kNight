package knight_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestKnight(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Knight Suite")
}
