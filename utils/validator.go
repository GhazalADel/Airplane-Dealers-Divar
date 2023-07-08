package utils

import (
	"Airplane-Divar/models"
	"errors"
	"fmt"
	"time"
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

func ValidateAd(jsonBody map[string]interface{}, cat models.Category) (string, models.AdminAds, error) {
	msg := "OK"

	var ad models.AdminAds

	ad.CategoryID = cat.ID

	if model, ok := jsonBody["model"].(string); ok {
		ad.AirplaneModel = model
	} else {
		msg := "Plane Model should be string !"
		return msg, models.AdminAds{}, errors.New("")
	}

	if price, ok := jsonBody["price"].(uint64); ok {
		ad.Price = price
	} else {
		msg := "Price should be integer !"
		return msg, models.AdminAds{}, errors.New("")
	}

	if fly, ok := jsonBody["fly_time"].(uint); ok {
		ad.FlyTime = fly
	} else {
		msg := "Fly time should be integer !"
		return msg, models.AdminAds{}, errors.New("")
	}

	if rc, ok := jsonBody["repair_check"].(bool); ok {
		ad.RepairCheck = rc
	} else {
		msg := "Repair Check should be boolean !"
		return msg, models.AdminAds{}, errors.New("")
	}

	if ec, ok := jsonBody["expert_check"].(bool); ok {
		ad.ExpertCheck = ec
	} else {
		msg := "Expert Check should be boolean !"
		return msg, models.AdminAds{}, errors.New("")
	}

	if age, ok := jsonBody["age"].(uint); ok {
		if uint(time.Now().Year())-age < 1903 {
			msg := "The year of the invention of the airplane was 1903 !"
			return msg, models.AdminAds{}, errors.New("")
		}
		ad.PlaneAge = age
	} else {
		msg := "Plane Age should be integer !"
		return msg, models.AdminAds{}, errors.New("")
	}

	//check for optional properties
	if _, ok := jsonBody["image"]; ok {
		if image, ok := jsonBody["image"].(string); ok {
			ad.Image = image
		} else {
			msg := "Image should be an url !"
			return msg, models.AdminAds{}, errors.New("")
		}
	} else {
		ad.Image = ""
	}

	if _, ok := jsonBody["subject"]; ok {
		if sub, ok := jsonBody["subject"].(string); ok {
			ad.Subject = sub
		} else {
			msg := "subject should be string !"
			return msg, models.AdminAds{}, errors.New("")
		}
	} else {
		ad.Subject = fmt.Sprintf("%d years old airplane : %s in the %s category", ad.PlaneAge, ad.AirplaneModel, cat)
	}

	if _, ok := jsonBody["description"]; ok {
		if desc, ok := jsonBody["description"].(string); ok {
			ad.Description = desc
		} else {
			msg := "description should be string !"
			return msg, models.AdminAds{}, errors.New("")
		}
	} else {
		ad.Description = fmt.Sprintf("Model : %s | Age : %d | Category : %s | Price : %d | Fly Time : %d | Has Expert Check : %v | Has Repair Check : %v", ad.AirplaneModel, ad.PlaneAge, cat, ad.Price, ad.FlyTime, ad.ExpertCheck, ad.RepairCheck)
	}

	return msg, ad, nil
}
