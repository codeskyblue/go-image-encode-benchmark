package main

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"log"
	"reflect"
	"runtime"
	"testing"

	"image/jpeg"
	"image/png"

	turboJpeg "github.com/pixiv/go-libjpeg/jpeg"
	"golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

var (
	cachedData = map[string][]byte{}
	img        image.Image
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

func init() {
	// init read
	for _, ft := range []string{"jpg"} {
		GetFileReader(ft)
	}
}

func setupImage(ft string) {
	var err error
	img, _, err = image.Decode(GetFileReader(ft))
	if err != nil {
		panic(err)
	}
	fmt.Printf("image type: %T\n", img)
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
	buf := bytes.NewBuffer(nil)
	for n := 0; n < b.N; n++ {
		buf.Reset()
		jpeg.Encode(buf, img, &jpeg.Options{jpeg.DefaultQuality})
	}
}

func BenchmarkTurboJpegEncode(b *testing.B) {
	buf := bytes.NewBuffer(nil)
	for n := 0; n < b.N; n++ {
		buf.Reset()
		err := turboJpeg.Encode(buf, img, &turboJpeg.EncoderOptions{Quality: 70})
		if err != nil {
			b.Fatal(err)
		}
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

func testTurboEncodeRGBA() {
	setupImage("png")
	buf := bytes.NewBuffer(nil)
	buf.Reset()
	err := turboJpeg.Encode(buf, img, &turboJpeg.EncoderOptions{Quality: 70})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ioutil.WriteFile("out.jpg", buf.Bytes(), 0644))
}

func testTurboEncodeGray() {
	setupImage("png")
	data, _ := ioutil.ReadFile("image-gray.png")
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.NewBuffer(nil)
	err = turboJpeg.Encode(buf, img, &turboJpeg.EncoderOptions{})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//testTurboEncodeRGBA()
	//testTurboEncodeGray()
	//return

	println("-- Decode --")
	//runBench(BenchmarkTurboJpegDecode) // commit it on windows
	runBench(BenchmarkJpegDecode)
	runBench(BenchmarkPngDecode)
	runBench(BenchmarkWebpDecode)
	runBench(BenchmarkTiffDecode)
	println("-- Encode --")
	setupImage("jpg")
	runBench(BenchmarkJpegEncode)
	runBench(BenchmarkPngEncode)
	runBench(BenchmarkTiffEncode)
	runBench(BenchmarkTurboJpegEncode) // commit it on windows
	println("-- Encode --")
	setupImage("png")
	runBench(BenchmarkJpegEncode)
	runBench(BenchmarkPngEncode)
	runBench(BenchmarkTiffEncode)
	runBench(BenchmarkTurboJpegEncode) // commit it on windows
}
