package dto

import "truthly/internals/model"

type ImageReqDto struct {
	UserId   string `json:"userId"`
	ImageUrl string `json:"imageUrl"`
}

type ImageResDto struct {
	ImageId   string       `json:"imageId"`
	ImageData *model.Image `json:"imageData"`
}

func ToImageModel(i *ImageReqDto) *model.Image {
	return &model.Image{
		ImageUrl: i.ImageUrl,
		UserId:   i.UserId,
	}
}
