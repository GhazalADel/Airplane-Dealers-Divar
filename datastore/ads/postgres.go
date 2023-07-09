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

func (a AdDatastorer) Get(id int, userRole string) ([]models.Ad, error) {
	var ads []models.Ad
	var result *gorm.DB
	result = a.db.Select([]string{"id", "user_id", "image", "description", "subject", "price", "category_id", "status", "fly_time", "airplane_model", "repair_check", "expert_check", "plane_age"})
	if id != 0 {
		result.Where("id = ?", id)
	}

	result, err := checkUserRole(userRole, result)
	if err != nil {
		return nil, err
	}

	if result.Find(&ads).Error != nil {
		return []models.Ad{}, fmt.Errorf("database error: Get ads from database")
	}
	return ads, nil
}

func (a AdDatastorer) ListFilterByColumn(f *filter.AdsFilter) (ads []models.Ad, err error) {
	builder := a.db.Select([]string{"id", "user_id", "image", "description", "subject", "price", "category_id", "status", "fly_time", "airplane_model", "repair_check", "expert_check", "plane_age"}).
		Offset(f.Base.Offset).
		Limit(f.Base.Limit).
		Order("id")

	builder, err = checkUserRole(f.Base.UserRole, builder)
	if err != nil {
		return nil, err
	}
	// if user is admin, can filter by status to check ads is active or not
	if f.Status != "" && f.Base.UserRole == "Admin" {
		builder = builder.Where("status = ?", f.Status)
	}

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

	if builder.Find(&ads).Error != nil {
		return nil, err
	}

	return ads, nil
}

func (a AdDatastorer) ListFilterSort(f *filter.Filter) (ads []models.Ad, err error) {
	var orderClause []string
	for col, order := range f.Sort {
		orderClause = append(orderClause, fmt.Sprintf("%s %s", col, order))
	}

	builder := a.db.Limit(f.Limit).
		Order(strings.Join(orderClause, ","))

	builder, err = checkUserRole(f.UserRole, builder)
	if err != nil {
		return nil, err
	}

	if builder.Find(&ads).Error != nil {
		return nil, err
	}

	return ads, nil
}

func checkUserRole(role string, builder *gorm.DB) (*gorm.DB, error) {
	switch role {
	case "Airline": // Airline
		builder = builder.Where("status = ?", "Active")
	case "Expert": // Expert
		builder = builder.Where("status = ? AND expert_check = ?", "Active", true)
	case "Matin": // Matin
		builder = builder.Where("status = ? AND repair_check = ?", "Active", true)
	case "Admin": // Admin
	default:
		return nil, fmt.Errorf("user role is not exist")
	}
	return builder, nil
}
