package main

import (
	"github.com/hawx/img/blur"
	"github.com/hawx/img/utils"
	"os"
)

var cmdBlur = &Command{
	UsageLine: "blur [options]",
	Short:     "blur an image",
Long: `
  Blur takes an image from STDIN, and prints a blurred version to STDOUT.

    --radius <r>                 # Set radius of blur (default: 2.0)
    --size <height>x<width>      # Set size of blur
    --style <option>             # Either clamp, ignore or wrap (default: ignore)

    --box                        # Perform box blur
    --gaussian <sigma>           # Perform gaussian blur (default: 5.0)
`,
}

var blurRadius int
var blurSize utils.Dimension
var blurStyle string

var blurBox bool
var blurGaussian float64

var styleNames map[string] blur.Style = map[string] blur.Style {
	"clamp":  blur.CLAMP,
	"ignore": blur.IGNORE,
	"wrap":   blur.WRAP,
}

func init() {
	cmdBlur.Run = runBlur

	cmdBlur.Flag.IntVar(&blurRadius, "radius", 2.0, "")
	cmdBlur.Flag.Var(&blurSize, "size", "")
	cmdBlur.Flag.StringVar(&blurStyle, "style", "ignore", "")

	cmdBlur.Flag.BoolVar(&blurBox, "box", false, "")
	cmdBlur.Flag.Float64Var(&blurGaussian, "gaussian", 5.0, "")
}

func runBlur(cmd *Command, args []string) {
	if _, ok := styleNames[blurStyle]; !ok {
		utils.Warn("--style must be one of 'clamp', 'ignore' or 'wrap'")
		os.Exit(2)
	}

	style, _ := styleNames[blurStyle]

	i := utils.ReadStdin()

	if !utils.FlagVisited("size", cmd.Flag) {
		diameter := blurRadius * 2 + 1
		blurSize = utils.Dimension{diameter, diameter}
	}

	if blurBox {
		i = blur.Box(i, blurSize, style)
	} else {
		i = blur.Gaussian(i, blurSize, blurGaussian, style)
	}

	utils.WriteStdout(i)
}
