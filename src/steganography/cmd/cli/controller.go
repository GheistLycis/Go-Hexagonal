package steganography

import (
	codec_infra "Go-Hexagonal/src/codec/infra"
	app "Go-Hexagonal/src/steganography/app"
	infra "Go-Hexagonal/src/steganography/infra"
	"log"
)

func HandleCommands(cmd CommandDTO) {
	codecService := codec_infra.NewCodecService()
	imgCodecService := infra.NewImageCodecService(codecService)
	codecImagService := app.NewCodecImageService(imgCodecService)

	// TODO: auto detect file type and encode more than just images
	switch cmd.Operation {
	case "encode":
		codecImagService.Encode(cmd.FilePath, cmd.Message)
	case "decode":
		codecImagService.Decode(cmd.FilePath)
	default:
		log.Fatalf("Unexpected operation \"%s\". Should be either \"encode\" or \"decode\"", cmd.Operation)
	}
}

type CommandDTO struct {
	Operation string
	FilePath  string
	Message   string
}
