package service

import (
	"gin-swagger/data/request"
	"gin-swagger/data/response"
	"gin-swagger/helper"
	"gin-swagger/model"
	"gin-swagger/repository"

	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func NewTagsServiceImpl(TagsRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: TagsRepository,
		Validate:       validate,
	}
}

func (t *TagsServiceImpl) Create(tags request.CreateTagsRequest) {
	err := t.Validate.Struct(tags)
	helper.ErrPanic(err)

	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

func (t *TagsServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

func (t *TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()

	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}

func (t *TagsServiceImpl) FindById(tagsId int) response.TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagsId)
	helper.ErrPanic(err)

	tagResponse := response.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}
	return tagResponse
}

func (t *TagsServiceImpl) Update(tags request.UpdateTagRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helper.ErrPanic(err)
	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}
