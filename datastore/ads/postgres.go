package ads

import (
	"Airplane-Divar/filter"
	"Airplane-Divar/models"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type AdDatastorer struct {
	db *gorm.DB
}

func New(db *gorm.DB) AdDatastorer {
	return AdDatastorer{db: db}
}

func (a AdDatastorer) Get(id int) ([]models.Ad, error) {
	var ads []models.Ad
	var result *gorm.DB

	if id != 0 {
		result = a.db.Where("status = ? AND id = ?", "Active", id).Find(&ads)
	} else {
		result = a.db.Where("status = ?", "Active").Find(&ads)
	}

	if result.Error != nil {
		return []models.Ad{}, fmt.Errorf("database error: Get ads from database")
	}
	return ads, nil
}

func (a AdDatastorer) ListFilterByColumn(f *filter.AdsFilter) (ads []models.Ad, err error) {
	builder := a.db.Select([]string{"id", "user_id", "image", "description", "subject", "price", "category_id", "status", "fly_time", "airplane_model", "repair_check", "expert_check", "plane_age"}).
		Offset(f.Base.Offset).
		Limit(f.Base.Limit).
		Order("id")

	if f.PlaneAge != 0 {
		builder = builder.Where("plane_age = ?", f.PlaneAge)
	}
	if f.Price != 0 {
		builder = builder.Where("price = ?", f.Price)
	}
	if f.FlyTime != 0 {
		builder = builder.Where("fly_time = ?", f.FlyTime)
	}
	if f.CategoryID != 0 {
		builder = builder.Where("category_id = ?", f.CategoryID)
	}
	if f.AirplaneModel != "" {
		builder = builder.Where("airplane_model = ?", f.AirplaneModel)
	}

	err = builder.Find(&ads).Error
	if err != nil {
		return nil, err
	}

	return ads, nil
}

func (a AdDatastorer) ListFilterSort(f *filter.Filter) (ads []models.Ad, err error) {
	var orderClause []string
	for col, order := range f.Sort {
		orderClause = append(orderClause, fmt.Sprintf("%s %s", col, order))
	}
	err = a.db.Limit(f.Limit).
		Order(strings.Join(orderClause, ",")).
		Find(&ads).
		Error
	if err != nil {
		return nil, err
	}

	return ads, nil
}

func (a AdDatastorer) GetCategoryByName(name string) (models.Category, error) {
	var categoryObj models.Category
	a.db.Where("name = ?", name).First(&categoryObj)
	if categoryObj.ID == 0 {
		msg := "Undefined Category Name !"
		return models.Category{}, fmt.Errorf(msg)
	}
	return categoryObj, nil
}

func (a AdDatastorer) CreateAdminAd(ad *models.AdminAds) (models.AdminAds, error) {
	var ad_ad *models.AdminAds
	ad_ad = ad
	createdAd := a.db.Create(&ad_ad)
	if createdAd.Error != nil {
		return models.AdminAds{}, fmt.Errorf("Admin Ad Creation Failed")
	}
	return *ad_ad, nil
}
