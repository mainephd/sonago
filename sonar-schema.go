// generated by goxsd; DO NOT EDIT
package main

// Coverage is generated from an XSD element
type coverage struct {
	//Version positiveInteger `xml:"version,attr"`
	File []File `xml:"file"`
}

// File is generated from an XSD element
type File struct {
	Path        string        `xml:"path,attr"`
	LineToCover []LineToCover `xml:"lineToCover"`
}

// LineToCover is generated from an XSD element
type LineToCover struct {
	LineNumber      int  `xml:"lineNumber,attr"`
	Covered         bool `xml:"covered,attr"`
	BranchesToCover int  `xml:"branchesToCover,attr,omitempty"`
	CoveredBranches int  `xml:"coveredBranches,attr,omitempty"`
}