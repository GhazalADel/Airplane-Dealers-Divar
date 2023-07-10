package utils

import (
	"Airplane-Divar/models"
	"errors"
	"fmt"
	"math"
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

func ValidateAd(jsonBody map[string]interface{}, cat models.Category) (string, models.Ad, error) {
	msg := "OK"

	var ad models.Ad

	ad.CategoryID = cat.ID

	if model, ok := jsonBody["AirplaneModel"].(string); ok {
		ad.AirplaneModel = model
	} else {
		msg = "Plane Model should be string !"
		return msg, models.Ad{}, errors.New("")
	}

	price, ok := jsonBody["Price"].(float64)
	if !ok {
		msg = "Price should be a number !"
		return msg, models.Ad{}, errors.New("")
	}
	if math.Mod(price, 1) != 0 {
		msg = "Price should be an integer !"
		return msg, models.Ad{}, errors.New("")
	}
	ad.Price = uint64(price)

	fly, ok := jsonBody["FlyTime"].(float64)
	if !ok {
		msg = "fly_time should be a number !"
		return msg, models.Ad{}, errors.New("")
	}
	if math.Mod(fly, 1) != 0 {
		msg = "fly_time should be an integer !"
		return msg, models.Ad{}, errors.New("")
	}
	ad.FlyTime = uint(fly)

	if rc, ok := jsonBody["RepairCheck"].(bool); ok {
		ad.RepairCheck = rc
	} else {
		msg = "Repair Check should be boolean !"
		return msg, models.Ad{}, errors.New("")
	}

	if ec, ok := jsonBody["ExpertCheck"].(bool); ok {
		ad.ExpertCheck = ec
	} else {
		msg = "Expert Check should be boolean !"
		return msg, models.Ad{}, errors.New("")
	}

	age, ok := jsonBody["PlaneAge"].(float64)
	if !ok {
		msg = "Age should be a number !"
		return msg, models.Ad{}, errors.New("")
	}
	if math.Mod(age, 1) != 0 {
		msg = "Age should be an integer !"
		return msg, models.Ad{}, errors.New("")
	}
	if uint(time.Now().Year())-uint(age) < 1903 {
		msg = "The year of the invention of the airplane was 1903 !"
		return msg, models.Ad{}, errors.New("")
	}
	ad.PlaneAge = uint(age)

	//check for optional properties
	if _, ok := jsonBody["Image"]; ok {
		if image, ok := jsonBody["Image"].(string); ok {
			ad.Image = image
		} else {
			msg = "Image should be an url !"
			return msg, models.Ad{}, errors.New("")
		}
	} else {
		ad.Image = ""
	}

	if _, ok := jsonBody["Subject"]; ok {
		if sub, ok := jsonBody["Subject"].(string); ok {
			ad.Subject = sub
		} else {
			msg = "subject should be string !"
			return msg, models.Ad{}, errors.New("")
		}
	} else {
		ad.Subject = fmt.Sprintf("%d years old airplane : %s in the %s category", ad.PlaneAge, ad.AirplaneModel, cat)
	}

	if _, ok := jsonBody["Description"]; ok {
		if desc, ok := jsonBody["Description"].(string); ok {
			ad.Description = desc
		} else {
			msg = "description should be string !"
			return msg, models.Ad{}, errors.New("")
		}
	} else {
		ad.Description = fmt.Sprintf("Model : %s | Age : %d | Category : %s | Price : %d | Fly Time : %d | Has Expert Check : %v | Has Repair Check : %v", ad.AirplaneModel, ad.PlaneAge, cat, ad.Price, ad.FlyTime, ad.ExpertCheck, ad.RepairCheck)
	}

	return msg, ad, nil
}
