# goresize

Simple image resizing tool written in Golang

## Usage

```shell
$ goresize -i input.png -h 600 -w 400 -o output.png
$ goresize -i input.png -h 600 -w 400 -o output.png
$ goresize -i input.png -h 600 -o output.png
$ goresize -i input.png -w 600 -o output.png
$ goresize -i input.png -w 600 # -> output.png
```

 ```text
[General options]
 --height, -h: image height (defaults to: 100)
  --width, -w: image width (defaults to: 0)
  --input, -i: input filename
 --output, -o: output filename (defaults to: output.png)

```

