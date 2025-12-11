package dto

import (
	"truthly/internals/model"
)

type AnalyticReqDto struct {
	ImageId       string `json:"imageId"`
	DescriptionId string `json:"descriptionId"`
	UserId        string `json:"userId"`

	Like    int `json:"like"`
	Share   int `json:"share"`
	Comment int `json:"comment"`
}

type AnalyticResDto struct {
	AnalyticId   string          `json:"analyticId"`
	AnalyticData *model.Analytic `json:"analyticData"`
}

func ToAnalyticModel(a *AnalyticReqDto) *model.Analytic {
	return &model.Analytic{
		ImageId:       a.ImageId,
		DescriptionId: a.DescriptionId,
		UserId:        a.UserId,
		Like:          a.Like,
		Share:         a.Share,
		Comment:       a.Comment,
	}
}
