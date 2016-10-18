# go-image-encode-benchmark
Golang Image encode bench mark

## Setup
Before you import library, you need to install libjpeg-turbo.

On Ubuntu: `sudo apt-get install libjpeg-turbo8-dev`

On Mac OS X: `brew install libjpeg-turbo`

## Support types
<https://godoc.org/golang.org/x/image>

**Encode and Decode**

* jpg
* png
* bmp
* tiff

**Only Deocde**

* webp
* vp8 (todo)
* vp8l (todo)

## Result
2016-10-14

> System: Win 10
> Memory: 8G
> CPU: Core(TM) i5-4570 3.20GHz

```
-- Decode --
main.BenchmarkJpegDecode        50      32 ms/op
main.BenchmarkPngDecode         20      94 ms/op
main.BenchmarkWebpDecode        20      58 ms/op
main.BenchmarkTiffDecode        100     11 ms/op
-- Encode --
main.BenchmarkJpegEncode        20      68 ms/op
main.BenchmarkPngEncode         3       438 ms/op
main.BenchmarkTiffEncode        5       239 ms/op
```

> System: Macmini OS X 10.11.6
> Memory: 4GB
> CPU: 1.4GHz Core i5

```
-- Decode --
main.BenchmarkJpegDecode        30      40 ms/op
main.BenchmarkPngDecode         10      112 ms/op
main.BenchmarkWebpDecode        20      69 ms/op
main.BenchmarkTiffDecode        100     13 ms/op
-- Encode --
image type: *image.YCbCr
main.BenchmarkJpegEncode        5       207 ms/op
main.BenchmarkPngEncode         1       1183 ms/op
main.BenchmarkTiffEncode        3       411 ms/op
main.BenchmarkTurboJpegEncode   100     10 ms/op
-- Encode --
image type: *image.RGBA
main.BenchmarkJpegEncode        20      82 ms/op
main.BenchmarkPngEncode         2       501 ms/op
main.BenchmarkTiffEncode        5       275 ms/op
main.BenchmarkTurboJpegEncode   0       0 ms/op
```

## turbo-jpeg
turbo-jpeg use <https://github.com/pixiv/go-libjpeg>

## LICENSE
Under [MIT](LICENSE)
