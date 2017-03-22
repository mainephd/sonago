package main

import (
	"encoding/xml"
	"flag"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sonago", func() {
	var expectedCoverageOutput []byte
	var expectedFileName string
	var expectedCoverageStruct coverage
	BeforeEach(func() {
		expectedFileName = "coverage.xml"
		expectedCoverageOutput, _ = ioutil.ReadFile(expectedFileName)
		xml.Unmarshal(expectedCoverageOutput, expectedCoverageStruct)
	})
	Describe("when outputfile is specified it", func() {
		It("should write to the file specified", func() {
			outputFilename := "testspecifiedfile.xml"
			flag.Set("outputfile", outputFilename)
			main()
			file, err := ioutil.ReadFile(outputFilename)
			Expect(err).ToNot(HaveOccurred())
			var actualCoverageStruct coverage
			xml.Unmarshal(file, actualCoverageStruct)
			Expect(actualCoverageStruct).To(Equal(expectedCoverageStruct))
		})
	})
	Describe("when inputfile is specified", func() {
		It("should read to the file specified", func() {
			filename := "test.coverprofile"
			flag.Set("inputfile", filename)
			main()
			file, err := ioutil.ReadFile(expectedFileName)
			Expect(err).ToNot(HaveOccurred())
			var actualCoverageStruct coverage
			xml.Unmarshal(file, actualCoverageStruct)
			Expect(actualCoverageStruct).To(Equal(expectedCoverageStruct))
		})
	})
	Describe("when outputfile is not specified", func() {
		It("should write contents to coverage.xml", func() {
			main()
			file, err := ioutil.ReadFile(expectedFileName)
			Expect(err).ToNot(HaveOccurred())
			var actualCoverageStruct coverage
			xml.Unmarshal(file, actualCoverageStruct)
			Expect(actualCoverageStruct).To(Equal(expectedCoverageStruct))
		})
	})
	Describe("when inputfile is not specified", func() {
		It("should read contents from gover.coverprofile", func() {
			main()
			file, err := ioutil.ReadFile(expectedFileName)
			Expect(err).ToNot(HaveOccurred())
			var actualCoverageStruct coverage
			xml.Unmarshal(file, actualCoverageStruct)
			Expect(actualCoverageStruct).To(Equal(expectedCoverageStruct))
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
