package domains

import (
	"time"

	repo "github.com/C-dexTeam/codex/internal/repos/out"
	"github.com/google/uuid"
)

type Chapter struct {
	ID           uuid.UUID
	CourseID     uuid.UUID
	LanguageID   uuid.UUID
	RewardID     *uuid.UUID
	RewardAmount int32
	Title        string
	Description  string
	Content      string
	FuncName     string
	FrontTmp     string
	DockerTmp    string
	GrantsExp    bool
	Active       bool
	Tests        []Test
	Reward       *Reward
	CreatedAt    time.Time
	DeletedAt    *time.Time
}

func NewChapter(
	chapter *repo.TChapter,
	tests []repo.TTest,
	reward *repo.TReward,
) Chapter {
	var deletedAt *time.Time
	if chapter.DeletedAt.Valid {
		deletedAt = &chapter.DeletedAt.Time
	} else {
		deletedAt = nil
	}
	var rewardID *uuid.UUID
	if chapter.RewardID.Valid {
		r := uuid.MustParse(chapter.RewardID.UUID.String())
		rewardID = &r
	} else {
		rewardID = nil
	}

	return Chapter{
		ID:           chapter.ID,
		CourseID:     chapter.CourseID,
		LanguageID:   chapter.LanguageID,
		RewardID:     rewardID,
		RewardAmount: chapter.RewardAmount,
		Title:        chapter.Title,
		Description:  chapter.Description,
		Content:      chapter.Content,
		FuncName:     chapter.FuncName,
		FrontTmp:     chapter.FrontendTemplate,
		DockerTmp:    chapter.DockerTemplate,
		GrantsExp:    chapter.GrantsExperience,
		Active:       chapter.Active,
		Tests:        NewTests(tests),
		Reward:       NewReward(reward, nil),
		CreatedAt:    chapter.CreatedAt.Time,
		DeletedAt:    deletedAt,
	}
}

func NewChapters(
	chapters []repo.TChapter,
) []Chapter {
	var domainChapters []Chapter
	for _, chapter := range chapters {
		domainChapters = append(domainChapters, NewChapter(&chapter, nil, nil))
	}

	return domainChapters
}
