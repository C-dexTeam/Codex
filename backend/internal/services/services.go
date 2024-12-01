package services

import (
	"github.com/C-dexTeam/codex/internal/domains"
)

type IService interface {
	UtilService() IUtilService
	UserService() domains.IUserService
	UserProfileService() domains.IUserProfileService
	RoleService() domains.IRoleService
	AdminService() domains.IAdminService
	RewardService() domains.IRewardService
	ProgrammingLService() domains.IPLanguagesService
}

type Services struct {
	utilService        IUtilService
	userService        domains.IUserService
	adminService       domains.IAdminService
	userProfileService domains.IUserProfileService
	roleService        domains.IRoleService
	languageService    domains.ILanguagesService
	rewardService      domains.IRewardService
	pLanguageService   domains.IPLanguagesService
}

func CreateNewServices(
	validatorService IValidatorService,
	userRepository domains.IUserRepository,
	userProfileRepository domains.IUserProfileRepository,
	transactionRepository domains.ITransactionRepository,
	roleRepository domains.IRoleRepository,
	languageRepository domains.ILanguagesRepository,
	rewardRepository domains.IRewardRepository,
	attributeRepository domains.IAttributeRepository,
	pLanguageRepository domains.IPLanguagesRepository,

) *Services {
	utilsService := newUtilService(validatorService)
	userProfileService := newUserProfileService(userProfileRepository, roleRepository, utilsService)
	userService := newUserService(userRepository, userProfileRepository, transactionRepository, utilsService)
	adminService := newAdminService(userRepository, userProfileRepository, transactionRepository, utilsService)
	roleService := newRoleService(roleRepository)
	languageService := newLanguageService(languageRepository)
	rewardService := newRewardService(rewardRepository, attributeRepository)
	pLanguageService := newPLanguageService(pLanguageRepository)

	return &Services{
		utilService:        utilsService,
		userService:        userService,
		adminService:       adminService,
		userProfileService: userProfileService,
		roleService:        roleService,
		languageService:    languageService,
		rewardService:      rewardService,
		pLanguageService:   pLanguageService,
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

func (s *Services) LanguageService() domains.ILanguagesService {
	return s.languageService
}

func (s *Services) RewardService() domains.IRewardService {
	return s.rewardService
}

func (s *Services) ProgrammingLService() domains.IPLanguagesService {
	return s.pLanguageService
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
