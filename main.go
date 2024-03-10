// Task - To Implement Inheritance in Go Lang
// Base Type is File Reader
// Derived Types are Text file reader, PDF File Reader, BMP File Reader
package main

import "fmt"

// FileReader is the base type
type FileReader struct {
	FilePath string
}

// NewFileReader creates a new FileReader
func NewFileReader(filePath string) FileReader {
	return FileReader{FilePath: filePath}
}

// Print function for the base FileReader
func (fr *FileReader) Print() {
	fmt.Println("Reading file:", fr.FilePath)
}

// TextFileReader inherits from FileReader
type TextFileReader struct {
	FileReader
}

// NewTextFileReader creates a new TextFileReader
func NewTextFileReader(filePath string) TextFileReader {
	return TextFileReader{FileReader: NewFileReader(filePath)}
}

// Print function for TextFileReader
func (tfr *TextFileReader) Print() {
	tfr.FileReader.Print()
	fmt.Println("Reading text file and displaying content.")
}

// PDFFileReader inherits from FileReader
type PDFFileReader struct {
	FileReader
}

// NewPDFFileReader creates a new PDFFileReader
func NewPDFFileReader(filePath string) PDFFileReader {
	return PDFFileReader{FileReader: NewFileReader(filePath)}
}

// Print function for PDFFileReader
func (pfr *PDFFileReader) Print() {
	pfr.FileReader.Print()
	fmt.Println("Reading PDF file and extracting text.")
}

// BMPFileReader inherits from FileReader
type BMPFileReader struct {
	FileReader
}

// NewBMPFileReader creates a new BMPFileReader
func NewBMPFileReader(filePath string) BMPFileReader {
	return BMPFileReader{FileReader: NewFileReader(filePath)}
}

// Print function for BMPFileReader
func (bfr *BMPFileReader) Print() {
	bfr.FileReader.Print()
	fmt.Println("Reading BMP file and processing image.")
}

func main() {
	// Example usage
	textReader := NewTextFileReader("example.txt")
	pdfReader := NewPDFFileReader("example.pdf")
	bmpReader := NewBMPFileReader("example.bmp")

	textReader.Print()
	fmt.Println()

	pdfReader.Print()
	fmt.Println()

	bmpReader.Print()
}
