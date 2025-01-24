package sessionStore

import (
	"encoding/gob"
	"strconv"

	"github.com/C-dexTeam/codex/internal/config/models"
	repo "github.com/C-dexTeam/codex/internal/repos/out"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis/v2"
)

type SessionData struct {
	UserID        string
	UserProfileID string
	PublicKey     string
	RoleID        string
	Role          string
	Username      string
	Email         string
	Name          string
	Surname       string
	Level         int
	Experience    int
	NextLevelExp  int
	Streak        int
}

func (s *SessionData) ParseFromUser(user *repo.TUsersAuth, userProfile *repo.TUsersProfile, userRole *repo.TRole) {
	s.UserID = user.ID.String()
	s.UserProfileID = userProfile.ID.String()
	s.PublicKey = user.PublicKey.String
	s.RoleID = userProfile.RoleID.String()
	s.Role = userRole.Name
	s.Username = user.Username.String
	s.Email = user.Email.String
	s.Name = userProfile.Name.String
	s.Surname = userProfile.Surname.String
	s.Level = int(userProfile.Level.Int32)
	s.Experience = int(userProfile.Experience.Int32)
	s.NextLevelExp = int(userProfile.NextLevelExp.Int32)
	s.Streak = int(userProfile.Streak.Int32)
}

func (s *SessionData) SetPublicKey(publicKey string) {
	s.PublicKey = publicKey
}

func (s *SessionData) SetStreak(streak int) {
	s.Streak = streak
}

func (s *SessionData) SetLevel(level, experience, nextLevelExp int32) {
	s.Level = int(level)
	s.Experience = int(experience)
	s.NextLevelExp = int(nextLevelExp)
}

func (s *SessionData) SetNameSurname(name, surname string) {
	if name != "" {
		s.Name = name
	}
	if surname != "" {
		s.Surname = surname
	}
}

func GetSessionData(c *fiber.Ctx) *SessionData {
	user := c.Locals("user")
	if user == nil {
		return nil
	}
	sessionData, ok := user.(SessionData)
	if !ok {
		return nil
	}
	return &sessionData
}

func NewSessionStore(cfg *models.RedisConfig) *session.Store {
	// Redis storage configuration

	port, _ := strconv.Atoi(cfg.Port)
	redisStorage := redis.New(redis.Config{
		Host:     cfg.Driver,
		Port:     port,
		Username: "",
		Password: "",
		Database: 0,
	})

	// Session verilerini seri hale getirmek için struct'u kayıt ediyoruz
	gob.Register(SessionData{})

	// Fiber session store oluşturuyoruz
	return session.New(session.Config{
		CookieSecure:   true,
		CookieHTTPOnly: true,
		Storage:        redisStorage,
	})
}
