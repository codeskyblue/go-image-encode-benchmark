# go-image-encode-benchmark
Golang Image encode bench mark

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

```
-- Decode --
main.BenchmarkJpegDecode        50      32 ms/op
main.BenchmarkPngDecode         20      94 ms/op
main.BenchmarkWebpDecode        20      59 ms/op
main.BenchmarkTiffDecode        100     11 ms/op
-- Encode --
main.BenchmarkJpegEncode        10      180 ms/op
main.BenchmarkPngEncode         3       424 ms/op
main.BenchmarkTiffEncode        5       242 ms/op
```

## LICENSE
Under [MIT](LICENSE)