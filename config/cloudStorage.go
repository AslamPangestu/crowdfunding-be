package config

import (
	"log"
	"os"

	cloudinary "github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func NewCloudStorage() *cloudinary.Cloudinary {
	//GET CONFIG
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var CLOUD_KEY = os.Getenv("CLOUD_API_KEY")
	var CLOUD_SECRET = os.Getenv("CLOUD_API_SECRET")
	config, err := cloudinary.CreateFromParams(CLOUD_NAME, CLOUD_KEY, CLOUD_SECRET)
	if err != nil {
		log.Fatalf("Cloudinary Failure, %v\n", err.Error())
	}
	return config
}

func ConfigCloudStorage(folder string) uploader.UploadParams {
	return uploader.UploadParams{
		Folder: folder,
	}
}
