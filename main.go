package main

import (
	"fmt"
	"image/png"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

const (
	path   = "./characters"
	output = "./output/"
	width  = 200
	height = 200
)

func main() {

	createFolder()

	list, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}

	for _, characters := range list {
		resizeImage(characters.Name())
	}

}

func resizeImage(character string) {
	file, err := os.Open(path + "/" + character)
	if err != nil {
		fmt.Println(err)
	}

	img, err := png.Decode(file)
	if err != nil {
		fmt.Println(err)
	}
	file.Close()

	m := resize.Resize(width, height, img, resize.Lanczos3)

	out, err := os.Create(output + character)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("O tamanho da imagem: %s foi alterado com sucesso!\n", strings.Title(strings.Replace(character, ".png", "", -1)))
	defer out.Close()
	png.Encode(out, m)

}

func createFolder() {
	if _, err := os.Stat(output); os.IsNotExist(err) {
		os.Mkdir(output, 0755)
	}
}
