package main

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"reflect"
	"runtime"
	"testing"

	"image/jpeg"
	"image/png"

	"golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

var (
	cachedData      = map[string][]byte{}
	cachedImageData = map[string]image.Image{}
)

func GetFileReader(fileType string) io.Reader {
	data, ok := cachedData[fileType]
	if ok {
		return bytes.NewReader(data)
	}
	data, err := ioutil.ReadFile("image." + fileType)
	if err != nil {
		panic(err)
	}
	cachedData[fileType] = data
	return bytes.NewReader(data)
}

func GetImage(fileType string) image.Image {
	img, ok := cachedImageData[fileType]
	if ok {
		return img
	}
	img, _, err := image.Decode(GetFileReader(fileType))
	if err != nil {
		panic(err)
	}
	cachedImageData[fileType] = img
	return img
}

func init() {
	// init read
	for _, ft := range []string{"jpg"} {
		GetFileReader(ft)
	}
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// https://golang.org/pkg/image/jpeg/
func BenchmarkJpegDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rd := GetFileReader("jpg")
		image.Decode(rd)
	}
}

func BenchmarkJpegEncode(b *testing.B) {
	img := GetImage("jpg")
	buf := bytes.NewBuffer(nil)
	for n := 0; n < b.N; n++ {
		buf.Reset()
		jpeg.Encode(buf, img, &jpeg.Options{jpeg.DefaultQuality})
	}
}

// https://golang.org/pkg/image/jpeg/
func BenchmarkPngDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rd := GetFileReader("png")
		image.Decode(rd)
	}
}

func BenchmarkPngEncode(b *testing.B) {
	img := GetImage("png")
	buf := bytes.NewBuffer(nil)
	for n := 0; n < b.N; n++ {
		buf.Reset()
		png.Encode(buf, img)
	}
}

func BenchmarkTiffDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rd := GetFileReader("tiff")
		image.Decode(rd)
	}
}

func BenchmarkTiffEncode(b *testing.B) {
	img := GetImage("tiff")
	buf := bytes.NewBuffer(nil)
	for n := 0; n < b.N; n++ {
		buf.Reset()
		tiff.Encode(buf, img, &tiff.Options{tiff.Deflate, true})
	}
}

func BenchmarkWebpDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rd := GetFileReader("webp")
		image.Decode(rd)
	}
}

func runBench(benchFunc func(*testing.B)) {
	funcName := getFunctionName(benchFunc)
	br := testing.Benchmark(benchFunc)
	fmt.Printf("%-24s\t%d\t%d ms/op\n", funcName, br.N, br.NsPerOp()/1e6)
}

func main() {
	println("-- Decode --")
	runBench(BenchmarkJpegDecode)
	runBench(BenchmarkPngDecode)
	runBench(BenchmarkWebpDecode)
	runBench(BenchmarkTiffDecode)
	println("-- Encode --")
	runBench(BenchmarkJpegEncode)
	runBench(BenchmarkPngEncode)
	runBench(BenchmarkTiffEncode)
}
