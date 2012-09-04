# img

A selection of image manipulation tools.

Requires Go 1. Build all tools into `built/` by running,

``` bash
$ ./build
```

All tools respond to the `--help` flag, use it to get information on options
that are supported.

## Conversion with 'from' and 'to'

`from` takes an input image (jpeg, gif or png) and outputs a `.png` file.

``` bash
$ ./from image.jpg > image.png
```

`to` takes a png image and outputs some other file (jpeg or png).

``` bash
$ image.png < ./to image.jpg
```

These allow you to use other filetypes as input, and get different output:

``` bash
$ ./from input.jpg | ./greyscale | ./pxl | ./to output.jpg
```

## shuffle

Randomly shuffles pixels around the image. Use `-v` or `-h` to constrain it to
vertical or horizontal shuffling, respectively.

``` bash
$ ./shuffle --vertical < input.png > output.png
```

![Shuffle](http://github.com/hawx/img/raw/master/examples/shuffle.jpg)

## pixelate

Pixelates an image. Use `--size HxW` to set pixel size used.

``` bash
$ ./pixelate --size 10x50 < input.png > output.png
```

![Pixelate](http://github.com/hawx/img/raw/master/examples/pixelate.jpg)

## pxl

Implementation of the triangle filter from [pxl app][pxlapp], using the
algorithm loosely described by [revdancatt][rev].

``` bash
$ ./pxl --size 30x30 < input.png > output.png
```

![pxl](http://github.com/hawx/img/raw/master/examples/pxl.jpg)

## hxl

An (almost; that is I'm not sure this is exactly the same) implementation of the
equilateral triangle filter from [pxl app][pxlapp].

``` bash
$ ./hxl --width 50 < input.png > output.png
```

![hxl](http://github.com/hawx/img/raw/master/examples/hxl.jpg)

## greyscale

Creates a greyscale version of an image.

``` bash
$ ./greyscale --average < input.png > output.png
```

![Greyscale](http://github.com/hawx/img/raw/master/examples/greyscale.jpg)

## contrast

Adjusts the contrast of the given image.

``` bash
$ ./contrast --by -25 < input.png > output.png
```

![contrast](http://github.com/hawx/img/raw/master/examples/contrast.jpg)

## brightness

Adjusts the brightness of the given image.

``` bash
$ ./brightness --by -25 < input.png > output.png
```

![brightness](http://github.com/hawx/img/raw/master/examples/brightness.jpg)

## hue, saturation and lightness

Adjust the hue, saturation and lightness of the an image.

``` bash
$ ./hue --by -30 < input.png > output.png
$ ./saturation --by 0.3 < input.png > output.png
$ ./lightness --by -0.07 < input.png > output.png
```

![hsl](http://github.com/hawx/img/raw/master/examples/hsl.jpg)

## blend

Allows you to blend two images together using a blend mode. Takes one image from
STDIN (the base image, imagine the bottom layer in photoshop) and one image as
an argument (the blend image, the layer above).

``` bash
$ < input.png ./blend --screen blend.png --opacity 0.3 > output.png
```

![blend](http://github.com/hawx/img/raw/master/examples/blend.jpg)

# Composition

These tools have been created to do one task each, and to use standard
input/output so that they can be easily composed. For example;

``` bash
$ < input.png ./shuffle --horizontal | ./hxl | ./hue --by -20 > output.png
```

![Composed](http://github.com/hawx/img/raw/master/examples/composed.jpg)


[pxlapp]: http://kohlberger.net/apps/pxl
[rev]:    http://revdancatt.com/2012/03/31/the-pxl-effect-with-javascript-and-canvas-and-maths/
