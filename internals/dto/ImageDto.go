package dto

import "truthly/internals/model"

type ImageReqDto struct {
	UserId   string `json:"userId"`
	ImageUrl string `json:"imageUrl"`
}

func ToImageModel(i *ImageReqDto) *model.Image {
	return &model.Image{
		ImageUrl: i.ImageUrl,
		UserId:   i.UserId,
	}
}
