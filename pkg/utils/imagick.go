package utils

import (
	"fmt"
	"gopkg.in/gographics/imagick.v2/imagick"
	"strings"
)

const densityCommand string = "mogrify -density %d -colorspace RGB -units PixelsPerInch -trim %s %s"

func Density(density uint, inFile, outFIle string) error {
	_, err := imagick.ConvertImageCommand(strings.Split(fmt.Sprintf(densityCommand, density, inFile, outFIle), " "))
	if err != nil {
		return err
	}

	return nil
}
