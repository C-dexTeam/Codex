package models

type Defaults struct {
	Roles    RoleDefaults
	Language LanguageDefaults
	Limits   LimitDefaults
	Secret   string
}

type RoleDefaults struct {
	DefaultRole    string
	RoleWalletUser string
	RoleAdmin      string
	RolePublic     string
}

type LanguageDefaults struct {
	DefaultLanguage string
}

type LimitDefaults struct {
	DefaultLanguageLimit            int
	DefaultCourseLimit              int
	DefaultPopulerCourseLimit       int
	DefaultChapterLimit             int
	DefaultTestLimit                int
	DefaultRewardLimit              int
	DefaultUserLimit                int
	DefaultRoleLimit                int
	DefaultAttributeLimit           int
	DefaultProgrammingLanguageLimit int
}

// Initialize Defaults
func NewDefaults(secret string) Defaults {
	return Defaults{
		Roles: RoleDefaults{
			DefaultRole:    "user",
			RoleWalletUser: "wallet-user",
			RoleAdmin:      "admin",
			RolePublic:     "public",
		},
		Language: LanguageDefaults{
			DefaultLanguage: "EN",
		},
		Limits: LimitDefaults{
			DefaultLanguageLimit:            10,
			DefaultCourseLimit:              10,
			DefaultPopulerCourseLimit:       6,
			DefaultChapterLimit:             10,
			DefaultTestLimit:                10,
			DefaultRewardLimit:              10,
			DefaultUserLimit:                10,
			DefaultRoleLimit:                10,
			DefaultAttributeLimit:           10,
			DefaultProgrammingLanguageLimit: 10,
		},
		Secret: secret,
	}
}
