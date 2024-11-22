package services

import "github.com/C-dexTeam/codex/internal/domains"

type IService interface {
	UtilService() IUtilService
	UserService() domains.IUserService
	UserProfileService() domains.IUserProfileService
	RoleService() domains.IRoleService
	AdminService() domains.IAdminService
}

type Services struct {
	utilService        IUtilService
	userService        domains.IUserService
	adminService       domains.IAdminService
	userProfileService domains.IUserProfileService
	roleService        domains.IRoleService
}

func CreateNewServices(
	validatorService IValidatorService,
	userRepository domains.IUserRepository,
	userProfileRepository domains.IUserProfileRepository,
	transactionRepository domains.ITransactionRepository,
	roleRepository domains.IRoleRepository,

) *Services {
	utilsService := newUtilService(validatorService)
	userProfileService := newUserProfileService(userProfileRepository, utilsService)
	userService := newUserService(userRepository, userProfileRepository, transactionRepository, utilsService)
	adminService := newAdminService(userRepository, userProfileRepository, transactionRepository, utilsService)
	roleService := newRoleService(roleRepository)

	return &Services{
		utilService:        utilsService,
		userService:        userService,
		adminService:       adminService,
		userProfileService: userProfileService,
		roleService:        roleService,
	}
}

func (s *Services) UtilService() IUtilService {
	return s.utilService
}

func (s *Services) AdminService() domains.IAdminService {
	return s.adminService
}

func (s *Services) UserService() domains.IUserService {
	return s.userService
}

func (s *Services) UserProfileService() domains.IUserProfileService {
	return s.userProfileService
}

func (s *Services) RoleService() domains.IRoleService {
	return s.roleService
}

// ------------------------------------------------------

type IValidatorService interface {
	ValidateStruct(s any) error
}

type IUtilService interface {
	Validator() IValidatorService
}

// -------------------

type utilService struct {
	validatorService IValidatorService
}

func newUtilService(
	validatorService IValidatorService,
) IUtilService {
	return &utilService{
		validatorService: validatorService,
	}
}

func (s *utilService) Validator() IValidatorService {
	return s.validatorService
}
