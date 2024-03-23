package entity

import (
	"recything/features/admin/model"
)

func AdminModelToAdminCore(data model.Admin) AdminCore {
	return AdminCore{
		Id:              data.Id,
		Fullname:        data.Fullname,
		Image:           data.Image,
		Role:            data.Role,
		Email:           data.Email,
		Password:        data.Password,
		ConfirmPassword: data.ConfirmPassword,
		Status:          data.Status,
		CreatedAt:       data.CreatedAt,
		UpdatedAt:       data.UpdatedAt,
	}

}

func ListAdminModelToAdminCore(admins []model.Admin) []AdminCore {
	listAdmin := []AdminCore{}
	for _, admin := range admins {
		adminCore := AdminModelToAdminCore(admin)
		listAdmin = append(listAdmin, adminCore)
	}
	return listAdmin
}

func AdminCoreToAdminModel(admin AdminCore) model.Admin {
	return model.Admin{
		Id:        admin.Id,
		Fullname:  admin.Fullname,
		Image:     admin.Image,
		Role:      admin.Role,
		Email:     admin.Email,
		Password:  admin.Password,
		Status:    admin.Status,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}
}

func ListAdminCoreToAdminModel(admins []AdminCore) []model.Admin {
	listAdmin := []model.Admin{}
	for _, admin := range admins {
		adminModel := AdminCoreToAdminModel(admin)
		listAdmin = append(listAdmin, adminModel)
	}
	return listAdmin
}
