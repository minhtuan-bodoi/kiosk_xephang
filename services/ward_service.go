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

func CreateWard(ward model.Ward) error {
	ward.Name = strings.TrimSpace(ward.Name)
	ward.Code = strings.TrimSpace(ward.Code)
	if ward.Name == "" {
		return errors.New("Không được để trống tên phường")
	}
	if ward.Code == "" {
		return errors.New("Không được để trống mã phường")
	}
	if ward.ProvinceCode == "" {
		return errors.New("province_id không được để trống")
	}

	if exists, err := repositories.IsProvinceCodeExists(ward.ProvinceCode); err != nil {
		return err
	} else if !exists {
		return fmt.Errorf("province_id '%s' không tồn tại", ward.ProvinceCode)
	}

	if exists, err := repositories.IsWardCodeExists(ward.Code); err != nil {
		return err
	} else if exists {
		return fmt.Errorf("code '%s' already exists in database", ward.Code)
	}

	return repositories.CreateWard(ward)
}

func GetAllWards() ([]model.Ward, error) {
	wards, err := repositories.GetWards()
	if err != nil {
		return wards, err
	}

	for i, ward := range wards {
		if province, err := repositories.GetProvincesByCodeProvince(ward.ProvinceCode); err == nil {
			wards[i].ProvinceName = province.Name
		}
	}

	return wards, nil
}

func GetWardByID(id string) (model.Ward, error) {
	ward, err := repositories.GetWardByID(id)
	if err != nil {
		return ward, err
	}

	if province, err := repositories.GetProvincesByCodeProvince(ward.ProvinceCode); err == nil {
		ward.ProvinceName = province.Name
	}

	return ward, nil
}

func GetWardsByCodeWard(code string) (model.Ward, error) {
	ward, err := repositories.GetWardsByCodeWard(code)
	if err != nil {
		return ward, err
	}

	if province, err := repositories.GetProvincesByCodeProvince(ward.ProvinceCode); err == nil {
		ward.ProvinceName = province.Name
	}

	return ward, nil
}

func UpdateWard(id bson.ObjectID, req dto.UpdateWardRequest) error {
	updates := bson.M{}

	if req.Name != nil {
		name := strings.TrimSpace(*req.Name)
		if name == "" {
			return fmt.Errorf("%w: tên phường không được để trống", errors.New("Dữ liệu ko hợp lệ"))
		}
		updates["name"] = name
	}

	if req.Code != nil {
		code := strings.TrimSpace(*req.Code)
		if code == "" {
			return fmt.Errorf("%w: mã phường không được để trống", errors.New("Dữ liệu ko hợp lệ"))
		}
		updates["code"] = code
	}

	if req.ProvinceID != nil {
		provinceID := strings.TrimSpace(*req.ProvinceID)
		if provinceID == "" {
			return fmt.Errorf("%w: province_id không được để trống", errors.New("Dữ liệu ko hợp lệ"))
		}

		if exists, err := repositories.IsProvinceCodeExists(provinceID); err != nil {
			return err
		} else if !exists {
			return fmt.Errorf("%w: province_id '%s' không tồn tại", errors.New("Dữ liệu ko hợp lệ"), provinceID)
		}

		updates["province_id"] = provinceID
	}

	if len(updates) == 0 {
		return errors.New("Phải có nội dung ít nhất được cập nhật")
	}

	return repositories.UpdateWard(id, updates)
}

func DeleteWard(id string) error {
	return repositories.DeleteWard(id)
}
