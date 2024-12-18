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
	UserService() *UserService
	UserProfileService() *userProfileService
	LanguageService() *languageService
	RoleService() *RoleService
	AdminService() domains.IAdminService
	RewardService() *rewardService
	ProgrammingService() domains.IPLanguagesService
	CourseService() domains.ICourseService
	ChapterService() domains.IChapterService
	AttributeService() *attributeService
	TestService() domains.ITestService
}

type Services struct {
	utilService        IUtilService
	userService        *UserService
	roleService        *RoleService
	userProfileService *userProfileService
	languageService    *languageService
	rewardService      *rewardService
	attributeService   *attributeService
	adminService       domains.IAdminService
	pLanguageService   domains.IPLanguagesService
	courseService      domains.ICourseService
	chapterService     domains.IChapterService
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
	roleService := newRoleService(db, queries, utilService)
	languageService := newLanguageService(db, queries, utilService)
	rewardService := newRewardService(db, queries, utilService)
	attributeService := NewAttributeService(db, queries, utilService)
	// adminService := newAdminService(userRepository, userProfileRepository, transactionRepository, utilsService)
	// pLanguageService := newPLanguageService(pLanguageRepository)
	// courseService := newCourseService(courseRepository, chapterRepository)
	// chapterService := NewChapterService(chapterRepository)
	// testService := newTestService(testRepository)

	return &Services{
		utilService:        utilService,
		userProfileService: userProfileService,
		userService:        userService,
		roleService:        roleService,
		languageService:    languageService,
		rewardService:      rewardService,
		attributeService:   attributeService,
	}
}

func (s *Services) UtilService() IUtilService {
	return s.utilService
}

func (s *Services) AdminService() domains.IAdminService {
	return s.adminService
}

func (s *Services) UserService() *UserService {
	return s.userService
}

func (s *Services) UserProfileService() *userProfileService {
	return s.userProfileService
}

func (s *Services) RoleService() *RoleService {
	return s.roleService
}

func (s *Services) LanguageService() *languageService {
	return s.languageService
}

func (s *Services) RewardService() *rewardService {
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

func (s *Services) AttributeService() *attributeService {
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
	NParseUUID(id string) (uuid.UUID, error)
	ParseString(str string) sql.NullString
	ParseNullUUID(str string) uuid.NullUUID
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

func (s *utilService) NParseUUID(id string) (uuid.UUID, error) {
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

func (s *utilService) ParseString(str string) sql.NullString {
	var value string
	var valid bool

	if str == "" {
		value = ""
		valid = false
	} else {
		value = str
		valid = true
	}

	return sql.NullString{String: value, Valid: valid}
}

func (s *utilService) ParseNullUUID(str string) uuid.NullUUID {
	var value uuid.UUID
	var valid bool

	if str == "" {
		valid = false
	} else {
		parsedUUID, err := uuid.Parse(str)
		if err != nil {
			valid = false
		} else {
			value = parsedUUID
			valid = true
		}
	}

	return uuid.NullUUID{UUID: value, Valid: valid}
}
