package main

import (
	"flag"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sonago", func() {
	var expectedCoverageOutput []byte
	var expectedFileName string
	BeforeEach(func() {
		expectedFileName = "coverage.xml"
		expectedCoverageOutput, _ = ioutil.ReadFile(expectedFileName)
	})
	Describe("when outputfile is specified it", func() {
		It("should write to the file specified", func() {
			outputFilename := "testspecifiedfile.xml"
			flag.Set("outputfile", outputFilename)
			main()
			file, err := ioutil.ReadFile(outputFilename)
			Expect(err).ToNot(HaveOccurred())
			Expect(file).To(Equal(expectedCoverageOutput))
		})
	})
	Describe("when inputfile is specified", func() {
		It("should read to the file specified", func() {
			filename := "test.coverprofile"
			flag.Set("inputfile", filename)
			main()
			file, err := ioutil.ReadFile(expectedFileName)
			Expect(err).ToNot(HaveOccurred())
			Expect(file).To(Equal(expectedCoverageOutput))
		})
	})
	Describe("when outputfile is not specified", func() {
		It("should write contents to coverage.xml", func() {
			main()
			file, err := ioutil.ReadFile(expectedFileName)
			Expect(err).ToNot(HaveOccurred())
			Expect(file).To(Equal(expectedCoverageOutput))
		})
	})
	Describe("when inputfile is not specified", func() {
		It("should read contents from gover.coverprofile", func() {
			main()
			file, err := ioutil.ReadFile(expectedFileName)
			Expect(err).ToNot(HaveOccurred())
			Expect(file).To(Equal(expectedCoverageOutput))

		})
	})

	Describe("When errors", func() {
		Describe("occur opening inputfile it", func() {
			It("cause the utility to panic", func() {
				flag.Set("inputfile", "iamnotpresent.coverprofile")
				Expect(main).To(Panic())
			})
		})
		Describe("occur parsing coverprofile it", func() {
			It("cause the utility to panic", func() {
				flag.Set("inputfile", "bad-gover.coverprofile")
				Expect(main).To(Panic())
			})
		})
	})
})
