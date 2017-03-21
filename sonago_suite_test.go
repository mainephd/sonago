package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSonago(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sonago Suite")
}
