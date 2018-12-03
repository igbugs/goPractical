package models

import "github.com/jinzhu/gorm"

type Article struct {
	Model

	TagId int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
}

func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, err
	}

	return false, nil
}

func GetArticleTotal(maps interface{}) (count int, err error) {
	err = db.Model(&Article{}).Where(maps).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	return count, nil
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []*Article, err error) {
	err = db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articles, nil
}

func GetArticle(id int) (article *Article, err error) {
	//db.Where("id = ?", id).First(&article)
	//db.Model(&article).Related(&article.Tag)
	err = db.Where("id = ? AND delete_on = ?", id, 0).First(&article).Related(&article.Tag, "TagID").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

func EditArticle(id int, data interface{}) error {
	db.Model(&Article{}).Where("id = ?", id).Update(data)
	return nil
}

func AddArticle(data map[string]interface{}) (err error) {
	err = db.Create(&Article{
		TagId:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CoverImageUrl: data["cover_image_url"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
	}).Error

	if err != nil {
		return err
	}

	return
}

func DeleteArticle(id int) error {
	if err := db.Where("id = ?", id).Delete(Article{}).Error; err != nil {
		return err
	}
	return nil
}

//func (article *Article) BeforeCreate(scope *gorm.Scope) error {
//	scope.SetColumn("CreateOn", time.Now().Unix())
//	return nil
//}
//
//func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
//	scope.SetColumn("ModifiedOn", time.Now().Unix())
//	return nil
//}

func CleanAllArticle() error {
	if err := db.Unscoped().Where("deleted_on != ?", 0).Delete(&Article{}).Error; err != nil {
		return err
	}
	return nil
}
