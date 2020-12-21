package services

import (
	"fmt"
	"github.com/GlebSolncev/http-rest-api/app/models"
	"github.com/GlebSolncev/http-rest-api/configs"
	_ "github.com/go-sql-driver/mysql"
)

//GetAllImages Fetch all image data
func GetAllImages(image *[]models.Image) (err error) {
	if err := configs.DB.Find(image).Error; err != nil {
		return err
	}
	return nil
}

//CreateImages ... Insert New data
func CreateImages(image *models.Image) (err error) {
	if err = configs.DB.Create(image).Error; err != nil {
		return err
	}
	return nil
}

//GetImagesByID ... Fetch only one image by Id
func GetImagesByID(image *models.Image, id string) (err error) {
	if err = configs.DB.Where("id = ?", id).First(image).Error; err != nil {
		return err
	}
	return nil
}

//UpdateImages ... Update image
func UpdateImages(image *models.Image, id string) (err error) {
	fmt.Println(image)
	configs.DB.Save(image)
	return nil
}

//DeleteImages ... Delete image
func DeleteImages(image *models.Image, id string) (err error) {
	configs.DB.Where("id = ?", id).Delete(image)
	return nil
}
