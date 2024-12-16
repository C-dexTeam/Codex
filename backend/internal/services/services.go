package services

import (
	"database/sql"

	"github.com/C-dexTeam/codex/internal/domains"
	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	repo "github.com/C-dexTeam/codex/internal/repos/out"
	"github.com/google/uuid"
)

type IService interface {
	UtilService() IUtilService
	UserService() *userService
	UserProfileService() *userProfileService
	RoleService() domains.IRoleService
	AdminService() domains.IAdminService
	RewardService() domains.IRewardService
	ProgrammingService() domains.IPLanguagesService
	CourseService() domains.ICourseService
	ChapterService() domains.IChapterService
	AttributeService() domains.IAttributeService
	LanguageService() domains.ILanguagesService
	TestService() domains.ITestService
}

type Services struct {
	utilService        IUtilService
	userService        *userService
	adminService       domains.IAdminService
	userProfileService *userProfileService
	roleService        domains.IRoleService
	languageService    domains.ILanguagesService
	rewardService      domains.IRewardService
	pLanguageService   domains.IPLanguagesService
	courseService      domains.ICourseService
	chapterService     domains.IChapterService
	attributeService   domains.IAttributeService
	testService        domains.ITestService
}

func CreateNewServices(
	validatorService IValidatorService,
	queries *repo.Queries,
	db *sql.DB,
) *Services {
	utilService := newUtilService(validatorService)
	userProfileService := newUserProfileService(db, queries, utilService)
	userService := newUserService(db, queries, utilService)
	// adminService := newAdminService(userRepository, userProfileRepository, transactionRepository, utilsService)
	// roleService := newRoleService(roleRepository)
	// languageService := newLanguageService(languageRepository)
	// rewardService := newRewardService(rewardRepository, attributeRepository)
	// pLanguageService := newPLanguageService(pLanguageRepository)
	// courseService := newCourseService(courseRepository, chapterRepository)
	// chapterService := NewChapterService(chapterRepository)
	// attributeService := NewAttributeService(attributeRepository)
	// testService := newTestService(testRepository)

	return &Services{
		utilService:        utilService,
		userService:        userService,
		userProfileService: userProfileService,
	}
}

func (s *Services) UtilService() IUtilService {
	return s.utilService
}

func (s *Services) AdminService() domains.IAdminService {
	return s.adminService
}

func (s *Services) UserService() *userService {
	return s.userService
}

func (s *Services) UserProfileService() *userProfileService {
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

func (s *Services) ProgrammingService() domains.IPLanguagesService {
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

func (s *Services) TestService() domains.ITestService {
	return s.testService
}

// ------------------------------------------------------

type IValidatorService interface {
	ValidateStruct(s any) error
}

type IUtilService interface {
	Validator() IValidatorService
	ParseUUID(id string) (uuid.UUID, error)
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

func (s *utilService) ParseUUID(id string) (uuid.UUID, error) {
	if id == "" {
		return uuid.UUID{}, nil
	}
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, serviceErrors.NewServiceErrorWithMessageAndError(
			errorDomains.StatusBadRequest,
			errorDomains.ErrInvalidID,
			err,
		)
	}
	return parsedUUID, nil
}
