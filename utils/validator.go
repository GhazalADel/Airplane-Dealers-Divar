package utils

import (
	"Airplane-Divar/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func ValidateJsonFormat(jsonBody map[string]interface{}, fields ...string) (string, error) {
	msg := "OK"
	for _, field := range fields {
		if _, ok := jsonBody[field]; !ok {
			msg = "Input Json doesn't include " + field
			break
		}
	}
	if msg != "OK" {
		return msg, errors.New("")
	}
	return msg, nil
}

func ValidateAd(jsonBody map[string]interface{}, db *gorm.DB) (string, models.Ad, error) {
	msg := "OK"

	var ad models.Ad

	//validate and initialize categoryID in ad object
	cat := ""
	if c, ok := jsonBody["category"].(string); ok {
		cat = c
	} else {
		msg := "Category should be string !"
		return msg, models.Ad{}, errors.New("")
	}

	var categoryObj models.Category
	db.Where("name = ?", cat).First(&categoryObj)
	if categoryObj.ID == 0 {
		msg := "Undefined Category Name !"
		return msg, models.Ad{}, errors.New("")
	}

	ad.CategoryID = categoryObj.ID

	//create ad object and check validation of its data types

	if model, ok := jsonBody["model"].(string); ok {
		ad.AirplaneModel = model
	} else {
		msg := "Plane Model should be string !"
		return msg, models.Ad{}, errors.New("")
	}

	if price, ok := jsonBody["price"].(float64); ok {
		price_int := int64(price)
		if price_int < 0 {
			msg := "Price should be Positive !"
			return msg, models.Ad{}, errors.New("")
		}
		ad.Price = price_int
	} else {
		msg := "Price should be integer !"
		return msg, models.Ad{}, errors.New("")
	}

	if fly, ok := jsonBody["fly_time"].(string); ok {
		flyTime, err := time.Parse("2006-01-02 03:04:05", fly)
		if err != nil {
			msg = "Incorrect Fly Time format"
			return msg, models.Ad{}, errors.New("")
		}
		ad.FlyTime = flyTime
	} else {
		msg := "Fly time should be in format yyyy-mm-dd hh:mm:ss !"
		return msg, models.Ad{}, errors.New("")
	}

	if rc, ok := jsonBody["repair_check"].(bool); ok {
		ad.RepairCheck = rc
	} else {
		msg := "Repair Check should be boolean !"
		return msg, models.Ad{}, errors.New("")
	}

	if ec, ok := jsonBody["expert_check"].(bool); ok {
		ad.ExpertCheck = ec
	} else {
		msg := "Expert Check should be boolean !"
		return msg, models.Ad{}, errors.New("")
	}

	if age, ok := jsonBody["age"].(float64); ok {
		age_int := int(age)
		if age_int < 1903 {
			msg := "The year of the invention of the airplane was 1903 !"
			return msg, models.Ad{}, errors.New("")
		}
		ad.PlaneAge = age_int
	} else {
		msg := "Plane Age should be integer !"
		return msg, models.Ad{}, errors.New("")
	}

	//check for optional properties
	if _, ok := jsonBody["image"]; ok {
		if image, ok := jsonBody["image"].(string); ok {
			ad.Image = image
		} else {
			msg := "Image should be an url !"
			return msg, models.Ad{}, errors.New("")
		}
	} else {
		ad.Image = ""
	}

	if _, ok := jsonBody["subject"]; ok {
		if sub, ok := jsonBody["subject"].(string); ok {
			ad.Subject = sub
		} else {
			msg := "subject should be string !"
			return msg, models.Ad{}, errors.New("")
		}
	} else {
		ad.Subject = fmt.Sprintf("from %d airplane : %s in the %s category", ad.PlaneAge, ad.AirplaneModel, cat)
	}

	if _, ok := jsonBody["description"]; ok {
		if desc, ok := jsonBody["description"].(string); ok {
			ad.Description = desc
		} else {
			msg := "description should be string !"
			return msg, models.Ad{}, errors.New("")
		}
	} else {
		ad.Description = fmt.Sprintf("Model : %s | Year : %d | Category : %s | Price : %d | Fly Time : %v | Has Expert Check : %v | Has Repair Check : %v", ad.AirplaneModel, ad.PlaneAge, cat, ad.Price, ad.FlyTime, ad.ExpertCheck, ad.RepairCheck)
	}

	ad.Status = "initial"

	return msg, ad, nil
}
