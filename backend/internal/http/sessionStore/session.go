package sessionStore

import (
	"encoding/gob"
	"strconv"

	"github.com/C-dexTeam/codex/internal/config/models"
	"github.com/C-dexTeam/codex/internal/domains"

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
}

func (s *SessionData) ParseFromUser(user *domains.User, userProfile *domains.UserProfile, userRole *domains.Role) {
	s.UserID = user.GetID().String()
	s.UserProfileID = userProfile.GetID().String()
	s.PublicKey = user.GetPublicKey()
	s.RoleID = userProfile.GetRoleID().String()
	s.Role = userRole.GetName()
	s.Username = user.GetUsername()
	s.Email = user.GetEmail()
	s.Name = userProfile.GetName()
	s.Surname = userProfile.GetSurname()
	s.Level = userProfile.GetLevel()
	s.Experience = userProfile.GetExperience()
	s.NextLevelExp = userProfile.GetNextLevelExperience()
}

func (s *SessionData) SetPublicKey(publicKey string) {
	s.PublicKey = publicKey
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
