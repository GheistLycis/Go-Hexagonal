package steganography

import (
	app "Go-Hexagonal/src/steganography/app"
	"image/png"
	"log"
	"os"
)

var hiddenMessage = "Secret message!"
var codecService = app.NewCodecService()
var imgCodecService = app.NewImageCodecService(codecService)

func HandleCommands(cmd CommandDTO) {
	encode()
	decode()
}

type CommandDTO struct {
}

func encode() {
	inFile, err := os.Open("input.png")
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	img, err := png.Decode(inFile)
	if err != nil {
		log.Fatal(err)
	}

	encodedImg := imgCodecService.Encode(img, hiddenMessage)

	outFile, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	png.Encode(outFile, encodedImg)
}

func decode() {
	inFile2, err := os.Open("output.png")
	if err != nil {
		log.Fatal(err)
	}
	defer inFile2.Close()

	decodedImg, err := png.Decode(inFile2)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Recovered message:", imgCodecService.Decode(decodedImg))
}
