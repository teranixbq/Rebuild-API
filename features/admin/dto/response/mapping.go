package response

import "recything/features/admin/entity"

func AdminCoreToAdminResponse(admin entity.AdminCore) AdminRespon {
	return AdminRespon{
		ID:        admin.Id,
		Fullname:  admin.Fullname,
		Image:     admin.Image,
		Email:     admin.Email,
		Status:    admin.Status,
		CreatedAt: admin.CreatedAt,
	}
}

func ListAdminCoreToAdminResponse(admins []entity.AdminCore) []AdminRespon {
	listAdmin := []AdminRespon{}
	for _, admin := range admins {
		adminResp := AdminCoreToAdminResponse(admin)
		listAdmin = append(listAdmin, adminResp)
	}
	return listAdmin
}

func AdminCoreToAdminResponseLogin(admin entity.AdminCore, token string) AdminResponseLogin {
	return AdminResponseLogin{
		ID:       admin.Id,
		Fullname: admin.Fullname,
		Email:    admin.Email,
		Image:    admin.Image,
		Token:    token,
	}
}
