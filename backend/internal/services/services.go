package services

import (
	"database/sql"

	"github.com/C-dexTeam/codex/internal/config/models"
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
	RewardService() *rewardService
	ProgrammingService() *pLanguageService
	CourseService() *courseService
	ChapterService() *chapterService
	AttributeService() *attributeService
}

type Services struct {
	utilService        IUtilService
	userService        *UserService
	roleService        *RoleService
	userProfileService *userProfileService
	languageService    *languageService
	rewardService      *rewardService
	attributeService   *attributeService
	pLanguageService   *pLanguageService
	courseService      *courseService
	chapterService     *chapterService
}

func CreateNewServices(
	validatorService IValidatorService,
	queries *repo.Queries,
	db *sql.DB,
	defaults *models.Defaults,
) *Services {
	utilService := newUtilService(validatorService, defaults)
	userProfileService := newUserProfileService(db, queries, utilService)
	userService := newUserService(db, queries, utilService)
	roleService := newRoleService(db, queries, utilService)
	languageService := newLanguageService(db, queries, utilService)
	rewardService := newRewardService(db, queries, utilService)
	attributeService := NewAttributeService(db, queries, utilService)
	pLanguageService := newPLanguageService(db, queries, utilService)
	courseService := newCourseService(db, queries, utilService)
	chapterService := NewChapterService(db, queries, utilService)

	return &Services{
		utilService:        utilService,
		userProfileService: userProfileService,
		userService:        userService,
		roleService:        roleService,
		languageService:    languageService,
		rewardService:      rewardService,
		attributeService:   attributeService,
		pLanguageService:   pLanguageService,
		courseService:      courseService,
		chapterService:     chapterService,
	}
}

func (s *Services) UtilService() IUtilService {
	return s.utilService
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

func (s *Services) ProgrammingService() *pLanguageService {
	return s.pLanguageService
}

func (s *Services) CourseService() *courseService {
	return s.courseService
}

func (s *Services) ChapterService() *chapterService {
	return s.chapterService
}

func (s *Services) AttributeService() *attributeService {
	return s.attributeService
}

// ------------------------------------------------------

type IValidatorService interface {
	ValidateStruct(s any) error
}

type IUtilService interface {
	Validator() IValidatorService
	D() *models.Defaults
	ParseUUID(id string) (uuid.UUID, error)  // ID can be null
	NParseUUID(id string) (uuid.UUID, error) // ID cannot be null
	ParseString(str string) sql.NullString
	ParseNullUUID(str string) uuid.NullUUID
}

// -------------------

type utilService struct {
	validatorService IValidatorService
	defaults         *models.Defaults
}

func newUtilService(
	validatorService IValidatorService,
	defaults *models.Defaults,
) IUtilService {
	return &utilService{
		validatorService: validatorService,
		defaults:         defaults,
	}
}

func (s *utilService) Validator() IValidatorService {
	return s.validatorService
}

func (s *utilService) D() *models.Defaults {
	return s.defaults
}

func (s *utilService) ParseUUID(id string) (uuid.UUID, error) {
	if id == "" {
		return uuid.UUID{}, nil
	}
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusBadRequest,
			serviceErrors.ErrInvalidID,
			err,
		)
	}
	return parsedUUID, nil
}

func (s *utilService) NParseUUID(id string) (uuid.UUID, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusBadRequest,
			serviceErrors.ErrInvalidID,
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
