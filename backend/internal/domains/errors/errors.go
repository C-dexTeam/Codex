package errorDomains

// HTTP Status Codes
const (
	StatusNotFound            = 404
	StatusBadRequest          = 400
	StatusInternalServerError = 500
)

// 404
const (
	ErrLanguageNotFound            = "LANGUAGE_NOT_FOUND"
	ErrCourseNotFound              = "COURSE_NOT_FOUND"
	ErrProgrammingLanguageNotFound = "PROGRAMMING_LANGUAGE_NOT_FOUND"
	ErrLanguageDefaultNotFound     = "DEFAULT_LANGUAGE_NOT_FOUND"
	ErrUserProfileNotFound         = "USER_PROFILE_NOT_FOUND"
	ErrUserNotFound                = "USER_NOT_FOUND"
	ErrRoleNotFound                = "ROLE_NOT_FOUND"
	ErrRewardNotFound              = "REWARD_NOT_FOUND"
	ErrRewardAttributeNotFound     = "REWARD_ATTRIBUTE_NOT_FOUND"
	ErrChapterNotFound             = "CHAPTER_NOT_FOUND"
)

// 400
const (
	ErrRewardNameCannotBeEmpty = "REWARD_NAME_CANNOT_BE_EMPTY"
	ErrRewardNameTooLong       = "REWARD_NAME_IS_TOO_LONG"

	ErrRewardTypeCannotBeEmpty = "REWARD_TYPE_CANNOT_BE_EMPTY"
	ErrRewardTypeTooLong       = "REWARD_TYPE_IS_TOO_LONG"

	ErrRewardSymbolCannotBeEmpty = "REWARD_SYMBOL_CANNOT_BE_EMPTY"
	ErrRewardSymbolTooLong       = "REWARD_SYMBOL_IS_TOO_LONG"

	ErrRewardImagePathCannotBeEmpty = "REWARD_IMAGE_PATH_CANNOT_BE_EMPTY"
	ErrRewardImagePathTooLong       = "REWARD_IMAGE_PATH_IS_TOO_LONG"

	ErrRewardURICannotBeEmpty = "REWARD_URI_CANNOT_BE_EMPTY"
	ErrRewardURITooLong       = "REWARD_URI_IS_TOO_LONG"

	ErrTraitTypeCannotBeEmpty = "TRAIT_TYPE_CANNOT_BE_EMPTY"
	ErrTraitTypeTooLong       = "TRAIT_TYPE_IS_TOO_LONG"

	ErrValueCannotBeEmpty = "VALUE_CANNOT_BE_EMPTY"
	ErrValueTooLong       = "VALUE_IS_TOO_LONG"

	ErrPLanguageNameCannotBeEmpty = "PROGRAMMING_LANGUAGE_NAME_CANNOT_BE_EMPTY"
	ErrPLanguageNameTooLong       = "PROGRAMMING_LANGUAGE_NAME_IS_TOO_LONG"

	ErrPLanguageDownloadCMDCannotBeEmpty = "PROGRAMMING_LANGUAGE_DOWNLOAD_CMD_CANNOT_BE_EMPTY"
	ErrPLanguageDownloadCMDTooLong       = "PROGRAMMING_LANGUAGE_DOWNLOAD_CMD_IS_TOO_LONG"

	ErrPLanguageCompileCMDCannotBeEmpty = "PROGRAMMING_LANGUAGE_COMPILE_CMD_CANNOT_BE_EMPTY"
	ErrPLanguageCompileCMDTooLong       = "PROGRAMMING_LANGUAGE_COMPILE_CMD_IS_TOO_LONG"

	ErrPLanguageImagePathCannotBeEmpty = "PROGRAMMING_LANGUAGE_IMAGE_PATH_CANNOT_BE_EMPTY"
	ErrPLanguageImagePathTooLong       = "PROGRAMMING_LANGUAGE_IMAGE_PATH_IS_TOO_LONG"

	ErrPLanguageFileExtentionCannotBeEmpty = "PROGRAMMING_LANGUAGE_FILE_EXTENTION_CANNOT_BE_EMPTY"
	ErrPLanguageFileExtentionTooLong       = "PROGRAMMING_LANGUAGE_FILE_EXTENTION_IS_TOO_LONG"

	ErrPLanguageMonacoEditorCannotBeEmpty = "PROGRAMMING_LANGUAGE_MONACO_EDITOR_CANNOT_BE_EMPTY"
	ErrPLanguageMonacoEditorTooLong       = "PROGRAMMING_LANGUAGE_MONACO_EDITOR_IS_TOO_LONG"

	ErrCourseTitleCannotBeEmpty = "COURSE_TITLE_CANNOT_BE_EMPTY"
	ErrCourseTitleTooLong       = "COURSE_TITLE_IS_TOO_LONG"

	ErrCourseImagePathCannotBeEmpty = "COURSE_IMAGE_PATH_CANNOT_BE_EMPTY"
	ErrCourseImagePathTooLong       = "COURSE_IMAGE_PATH_IS_TOO_LONG"

	ErrCourseRewardAmountCannotBeNegative = "COURSE_REWARD_AMOUNT_CANNOT_BE_NEGATIVE"
	ErrCourseRewardAmountEnterInteger     = "COURSE_REWARD_AMAOUNT_HAS_TO_BE_INTEGER_VALUE"

	ErrChapterTitleCannotBeEmpty = "CHAPTER_TITLE_CANNOT_BE_EMPTY"
	ErrChapterTitleTooLong       = "CHAPTER_TITLE_IS_TOO_LONG"

	ErrChapterFuncNameCannotBeEmpty = "CHAPTER_FUNCNAME_CANNOT_BE_EMPTY"
	ErrChapterFuncNameTooLong       = "CHAPTER_FUNCNAME_IS_TOO_LONG"

	ErrUsernameBeingUsed   = "USERNAME_ALREADY_BEING_USED"
	ErrInvalidAuth         = "USERNAME_OR_PASSWORD_WRONG"
	ErrEmailBeingUsed      = "EMAIL_ALREADY_BEING_USED"
	ErrPasswordsDoNotMatch = "PASSWORDS_DO_NOT_MATCH"

	// General
	ErrInvalidID               = "INVALID_ID"
	ErrInvalidBoolean          = "INVALID_BOOLEAN"
	ErrInvalidWalletConnection = "INVALID_WALLET_CONNECTION"
)

// 500
const (
	ErrErrorWhileFilteringRewards              = "ERROR_WHILE_FILTERING_REWARDS"
	ErrErrorWhileFilteringRewardsAttributes    = "ERROR_WHILE_FILTERING_REWARDS_ATTRIBUTES"
	ErrErrorWhileFilteringProgrammingLanguages = "ERROR_WHILE_FILTERING_PROGRAMMING_LANGUAGES"
	ErrErrorWhileFilteringUserProfile          = "ERROR_WHILE_FILTERING_USER_PROFILE"
	ErrErrorWhileFilteringRole                 = "ERROR_WHILE_FILTERING_ROLES"
	ErrErrorWhileFilteringCourse               = "ERROR_WHILE_FILTERING_COURSES"
	ErrErrorWhileFilteringChapter              = "ERROR_WHILE_FILTERING_CHAPTERS"
	ErrErrorWhileAddingExperience              = "ERROR_WHILE_ADDING_EXPERIENCE"
	ErrErrorWhileCreatingUserAuth              = "ERROR_WHILE_CREATING_USER_AUTH"
	ErrErrorWhileCreatingUserProfile           = "ERROR_WHILE_CREATING_USER_PROFILE"
	ErrErrorWhileFilteringUsers                = "ERROR_WHILE_FILTERING_USERS"
	ErrErrorWhileUpdatingUserAuth              = "ERROR_WHILE_UPDATING_USERS"
	ErrErrorWhileFilteringTests                = "ERROR_WHILE_FILTERING_TESTS"
	ErrErrorWhileFilteringInputs               = "ERROR_WHILE_FILTERING_INPUTS"
	ErrErrorWhileFilteringOutputs              = "ERROR_WHILE_FILTERING_OUTPUTS"
	ErrErrorWhileFilteringLanguages            = "ERROR_WHILE_FILTERING_LANGUAGES"
	ErrErrorWhileComparingPassword             = "ERROR_WHILE_COMPARING_PASSWORDS"
	ErrErrorWhileHashing                       = "ERROR_WHILE_HASHING"
	ErrTransactionError                        = "ERROR_WHILE_COMMITING"
	ErrWalletVerificationError                 = "ERROR_WHILE_VERIFICATON_WALLET"
)
