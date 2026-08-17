package services

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"kiosk-xephang/dto"
	"kiosk-xephang/model"
	"kiosk-xephang/repositories"
	"kiosk-xephang/utils"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const UploadDir = "./uploads/icons"

func CreateService(service model.Service) error {
	if service.CodeService != "" {
		exists, err := repositories.IsServiceCodeExists(service.CodeService)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("code_service '%s' already exists in database", service.CodeService)
		}
	}

	return repositories.CreateService(service)
}

func GetAllServices() ([]model.Service, error) {
	return repositories.GetServices()
}

func GetServicesByCodeService(code string) (model.Service, error) {
	return repositories.GetServicesByCodeService(code)
}

func GetServiceByID(id string) (model.Service, error) {
	return repositories.GetServiceByID(id)
}

func UpdateService(id bson.ObjectID, req dto.UpdateServiceRequest) error {
	updates := bson.M{}

	if req.NameServiceVN != nil {
		nameVN := strings.TrimSpace(*req.NameServiceVN)
		if nameVN == "" {
			return fmt.Errorf("%w: tên dịch vụ (VN) không được để trống", errors.New("Dữ liệu ko hợp lệ"))
		}
		updates["name_service_VN"] = nameVN
	}

	if req.NameServiceEN != nil {
		nameEN := strings.TrimSpace(*req.NameServiceEN)
		if nameEN == "" {
			return fmt.Errorf("%w: tên dịch vụ (EN) không được để trống", errors.New("Dữ liệu ko hợp lệ"))
		}
		updates["name_service_EN"] = nameEN
	}

	if req.CodeService != nil {
		code := strings.TrimSpace(*req.CodeService)
		if code == "" {
			return fmt.Errorf("%w: mã dịch vụ không được để trống", errors.New("Dữ liệu ko hợp lệ"))
		}
		existingService, err := repositories.GetServiceByID(id.Hex())
		if err != nil {
			return err
		}
		if code != existingService.CodeService {
			exists, err := repositories.IsServiceCodeExists(code)
			if err != nil {
				return err
			}
			if exists {
				return fmt.Errorf("code_service '%s' already exists in database", code)
			}
		}
		updates["code_service"] = code
	}

	var oldIcon string
	if req.IconService != nil {
		existingService, err := repositories.GetServiceByID(id.Hex())
		if err != nil {
			return err
		}
		oldIcon = existingService.IconService

		if !utils.IsAllowedImage(req.IconService.Filename) {
			return errors.New("invalid file type, allowed: jpg, jpeg, png, gif, svg, webp")
		}

		newFilename, err := utils.SaveUploadedFile(req.IconService, UploadDir)
		if err != nil {
			return fmt.Errorf("error saving new icon file: %w", err)
		}

		updates["icon_service"] = newFilename
	}

	if len(updates) == 0 {
		return errors.New("Phải có nội dung ít nhất được cập nhật")
	}

	if err := repositories.UpdateService(id, updates); err != nil {
		if newIcon, ok := updates["icon_service"].(string); ok && newIcon != "" {
			_ = utils.DeleteFile(filepath.Join(UploadDir, newIcon))
		}
		return err
	}

	// Delete old icon file if a new one was saved successfully
	if newIcon, ok := updates["icon_service"].(string); ok && oldIcon != "" && oldIcon != newIcon {
		_ = utils.DeleteFile(filepath.Join(UploadDir, oldIcon))
	}

	return nil
}

func DeleteService(id string) error {
	existingService, err := repositories.GetServiceByID(id)
	if err != nil {
		return err
	}

	if err := repositories.DeleteService(id); err != nil {
		return err
	}

	if existingService.IconService != "" {
		_ = utils.DeleteFile(filepath.Join(UploadDir, existingService.IconService))
	}

	return nil
}
