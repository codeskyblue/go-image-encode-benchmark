package main

import (
	"bytes"
	"image"
	"io/ioutil"
	"log"
	"testing"

	turboJpeg "github.com/pixiv/go-libjpeg/jpeg"
)

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
