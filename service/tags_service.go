package service

import (
	"gin-swagger/data/request"
	"gin-swagger/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
