package services

import (
	"errors"
	"fmt"
	"strings"

	"kiosk-xephang/dto"
	"kiosk-xephang/model"
	"kiosk-xephang/repositories"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateProvince(province model.Province) error {
	province.Name = strings.TrimSpace(province.Name)
	province.Code = strings.TrimSpace(province.Code)

	if province.Name == "" {
		return errors.New("Không được để trống tên tỉnh thành")
	}

	if province.Code == "" {
		return errors.New("Không được để trống mã tỉnh thành")
	}

	if exists, err := repositories.IsProvinceCodeExists(province.Code); err != nil {
		return err
	} else if exists {
		return fmt.Errorf("code '%s' already exists in database", province.Code)
	}

	return repositories.CreateProvince(province)
}

func GetAllProvinces() ([]model.Province, error) {
	return repositories.GetProvinces()
}

func GetProvinceByID(id string) (model.Province, error) {
	return repositories.GetProvinceByID(id)
}

func GetProvincesByCodeProvince(code string) (model.Province, error) {
	return repositories.GetProvincesByCodeProvince(code)
}

func UpdateProvince(id bson.ObjectID, req dto.UpdateProvinceRequest) error {
	updates := bson.M{}

	if req.Name != nil {
		name := strings.TrimSpace(*req.Name)
		if name == "" {
			return fmt.Errorf("%w: tên tỉnh thành không được để trống", errors.New("Dữ liệu ko hợp lệ"))
		}
		updates["name"] = name
	}

	if req.Code != nil {
		code := strings.TrimSpace(*req.Code)
		if code == "" {
			return fmt.Errorf("%w: mã tỉnh thành không được để trống", errors.New("Dữ liệu ko hợp lệ"))
		}
		updates["code"] = code
	}

	if len(updates) == 0 {
		return errors.New("Phải có nội dung ít nhất được cập nhật")
	}

	return repositories.UpdateProvince(id, updates)
}

func DeleteProvince(id string) error {
	return repositories.DeleteProvince(id)
}

func GetProvince() ([]model.Province, error) {
	return repositories.GetProvince()
}
