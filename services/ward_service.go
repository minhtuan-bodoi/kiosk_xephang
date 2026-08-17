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
	if ward.ProvinceID.IsZero() {
		return errors.New("province_id không được để trống")
	}

	if exists, err := repositories.IsWardCodeExists(ward.Code); err != nil {
		return err
	} else if exists {
		return fmt.Errorf("code '%s' already exists in database", ward.Code)
	}

	return repositories.CreateWard(ward)
}

func GetAllWards() ([]model.Ward, error) {
	return repositories.GetWards()
}

func GetWardByID(id string) (model.Ward, error) {
	return repositories.GetWardByID(id)
}

func GetWardsByCodeWard(code string) (model.Ward, error) {
	return repositories.GetWardsByCodeWard(code)
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
		objectID, err := bson.ObjectIDFromHex(provinceID)
		if err != nil {
			return errors.New("invalid province_id format")
		}
		updates["province_id"] = objectID
	}

	if len(updates) == 0 {
		return errors.New("Phải có nội dung ít nhất được cập nhật")
	}

	return repositories.UpdateWard(id, updates)
}

func DeleteWard(id string) error {
	return repositories.DeleteWard(id)
}
