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
	CourseService() domains.ICourseService
	ChapterService() domains.IChapterService
	AttributeService() domains.IAttributeService
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
	courseService      domains.ICourseService
	chapterService     domains.IChapterService
	attributeService   domains.IAttributeService
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
	courseRepository domains.ICourseRepository,
	chapterRepository domains.IChapterRepository,

) *Services {
	utilsService := newUtilService(validatorService)
	userProfileService := newUserProfileService(userProfileRepository, utilsService)
	userService := newUserService(userRepository, userProfileRepository, transactionRepository, utilsService)
	adminService := newAdminService(userRepository, userProfileRepository, transactionRepository, utilsService)
	roleService := newRoleService(roleRepository)
	languageService := newLanguageService(languageRepository)
	rewardService := newRewardService(rewardRepository, attributeRepository)
	pLanguageService := newPLanguageService(pLanguageRepository)
	courseService := newCourseService(courseRepository)
	chapterService := NewChapterService(chapterRepository)
	attributeService := NewAttributeService(attributeRepository)

	return &Services{
		utilService:        utilsService,
		userService:        userService,
		adminService:       adminService,
		userProfileService: userProfileService,
		roleService:        roleService,
		languageService:    languageService,
		rewardService:      rewardService,
		pLanguageService:   pLanguageService,
		courseService:      courseService,
		chapterService:     chapterService,
		attributeService:   attributeService,
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

func (s *Services) CourseService() domains.ICourseService {
	return s.courseService
}

func (s *Services) ChapterService() domains.IChapterService {
	return s.chapterService
}

func (s *Services) AttributeService() domains.IAttributeService {
	return s.attributeService
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
