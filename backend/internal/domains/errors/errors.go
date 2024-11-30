package errorDomains

// HTTP Status Codes
const (
	StatusNotFound            = 404
	StatusBadRequest          = 400
	StatusInternalServerError = 500
)

// 404
const (
	ErrLanguageNotFound        = "LANGUAGE_NOT_FOUND"
	ErrLanguageDefaultNotFound = "DEFAULT_LANGUAGE_NOT_FOUND"
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

	// General
	ErrInvalidID = "INVALID_ID"
)
