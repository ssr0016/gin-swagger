package repository

import (
	"errors"
	"gin-swagger/data/request"
	"gin-swagger/helper"
	"gin-swagger/model"

	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

func (t *TagsRepositoryImpl) Delete(tagsId int) {
	var tags model.Tags
	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
	helper.ErrPanic(result.Error)
}

func (t *TagsRepositoryImpl) FindAll() []model.Tags {
	var tags []model.Tags
	result := t.Db.Find(&tags)
	helper.ErrPanic(result.Error)
	return tags
}

func (t *TagsRepositoryImpl) FindById(tagsId int) (tags model.Tags, err error) {
	var tag model.Tags
	result := t.Db.Find(&tag, tagsId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

func (t *TagsRepositoryImpl) Save(tags model.Tags) {
	result := t.Db.Create(&tags)
	helper.ErrPanic(result.Error)
}

func (t *TagsRepositoryImpl) Update(tags model.Tags) {
	var updateTag = request.UpdateTagRequest{
		Id:   tags.Id,
		Name: tags.Name,
	}

	result := t.Db.Model(&tags).Updates(updateTag)
	helper.ErrPanic(result.Error)
}
