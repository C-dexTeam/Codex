package services

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/C-dexTeam/codex/internal/domains"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	dto "github.com/C-dexTeam/codex/internal/http/dtos"
	"github.com/C-dexTeam/codex/internal/http/response"
	repo "github.com/C-dexTeam/codex/internal/repos/out"
	"github.com/google/uuid"
)

type chapterService struct {
	db          *sql.DB
	queries     *repo.Queries
	utilService IUtilService
}

func NewChapterService(
	db *sql.DB,
	queries *repo.Queries,
	utilService IUtilService,
) *chapterService {
	return &chapterService{
		db:          db,
		queries:     queries,
		utilService: utilService,
	}
}

func (s *chapterService) GetChapters(
	ctx context.Context,
	id, langugeID, courseID, rewardID, title, grantsExperience, active, page, limit string,
) ([]domains.Chapter, error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = s.utilService.D().Limits.DefaultChapterLimit
	}

	// Hata var ise dönsün diye
	if _, err := s.utilService.ParseUUID(id); err != nil {
		return nil, err
	}
	if _, err := s.utilService.ParseUUID(langugeID); err != nil {
		return nil, err
	}
	if _, err := s.utilService.ParseUUID(courseID); err != nil {
		return nil, err
	}
	if _, err := s.utilService.ParseUUID(rewardID); err != nil {
		return nil, err
	}

	chapters, err := s.queries.GetChapters(ctx, repo.GetChaptersParams{
		ID:         s.utilService.ParseNullUUID(id),
		LanguageID: s.utilService.ParseNullUUID(langugeID),
		RewardID:   s.utilService.ParseNullUUID(rewardID),
		CourseID:   s.utilService.ParseNullUUID(courseID),
		Title:      s.utilService.ParseString(title),
		Lim:        int32(limitNum),
		Off:        (int32(pageNum) - 1) * int32(limitNum),
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringChapter,
			err,
		)
	}
	domainChapters := domains.NewChapters(chapters)

	return domainChapters, nil
}

func (s *chapterService) GetChapter(
	ctx context.Context,
	id, page, limit string,
) (*domains.Chapter, error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = s.utilService.D().Limits.DefaultTestLimit
	}

	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return nil, err
	}

	chapter, err := s.queries.GetChapter(ctx, idUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrChapterNotFound)
		}
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringChapter, err)
	}

	// TODO: Return tests with input and output by chapter id
	chapterTests, err := s.queries.GetTests(ctx, repo.GetTestsParams{
		ChapterID: s.utilService.ParseNullUUID(id),
		Lim:       int32(limitNum),
		Off:       (int32(pageNum) - 1) * int32(limitNum),
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringTests,
			err,
		)
	}

	var chapterReward repo.TReward
	if chapter.RewardID.Valid {
		chapterReward, err = s.queries.GetReward(ctx, chapter.RewardID.UUID)
		if err != nil {
			if strings.Contains(err.Error(), "sql: no rows in result set") {
				return nil, serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrRewardNotFound)
			}
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringRewards, err)
		}

	}
	domainChapter := domains.NewChapter(&chapter, chapterTests, &chapterReward)

	return &domainChapter, nil
}

func (s *chapterService) AddChapter(
	ctx context.Context,
	courseID, languageID, rewardID, title, description, content, funcName string,
	frontendTmp, dockerTmp, checkTmp string,
	grantsExperience, active bool,
	rewardAmount int,
) (uuid.UUID, error) {
	languageUUID, err := s.utilService.NParseUUID(languageID)
	if err != nil {
		return uuid.Nil, err
	}
	courseUUID, err := s.utilService.ParseUUID(courseID)
	if err != nil {
		return uuid.Nil, err
	}
	if _, err := s.utilService.ParseUUID(rewardID); err != nil {
		return uuid.Nil, err
	}

	id, err := s.queries.CreateChapter(ctx, repo.CreateChapterParams{
		LanguageID:       languageUUID,
		CourseID:         courseUUID,
		RewardID:         s.utilService.ParseNullUUID(rewardID),
		Title:            title,
		Description:      description,
		FuncName:         funcName,
		FrontendTemplate: frontendTmp,
		DockerTemplate:   dockerTmp,
		CheckTemplate:    checkTmp,
		RewardAmount:     int32(rewardAmount),
		GrantsExperience: grantsExperience,
		Active:           active,
	})
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (s *chapterService) UpdateChapter(
	ctx context.Context,
	id, courseID, languageID, rewardID, title, description, content, funcName string,
	frontendTmp, dockerTmp, checkTmp string,
	grantsExperience, active bool,
	rewardAmount int,
) error {
	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return err
	}

	if ok, err := s.queries.CheckChapterByID(ctx, idUUID); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringUsers, err)
	} else if !ok {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrUserNotFound)
	}

	var rewAmountNullInt sql.NullInt32
	if rewardAmount == 0 {
		rewAmountNullInt.Valid = false
	} else {
		rewAmountNullInt.Valid = true
		rewAmountNullInt.Int32 = int32(rewardAmount)
	}

	var grantsExpNullBool sql.NullBool
	if grantsExperience {
		grantsExpNullBool.Valid = true
		grantsExpNullBool.Bool = true
	} else {
		grantsExpNullBool.Valid = false
	}

	var validNulBool sql.NullBool
	if active {
		validNulBool.Valid = true
		validNulBool.Bool = true
	} else {
		validNulBool.Valid = false
	}

	if err := s.queries.UpdateChapter(ctx, repo.UpdateChapterParams{
		ChapterID:        idUUID,
		LanguageID:       s.utilService.ParseNullUUID(languageID),
		CourseID:         s.utilService.ParseNullUUID(courseID),
		RewardID:         s.utilService.ParseNullUUID(rewardID),
		Title:            s.utilService.ParseString(title),
		Description:      s.utilService.ParseString(description),
		Content:          s.utilService.ParseString(content),
		FuncName:         s.utilService.ParseString(funcName),
		FrontendTemplate: s.utilService.ParseString(frontendTmp),
		DockerTemplate:   s.utilService.ParseString(dockerTmp),
		CheckTemplate:    s.utilService.ParseString(checkTmp),
		RewardAmount:     rewAmountNullInt,
		GrantsExperience: grantsExpNullBool,
		Active:           validNulBool,
	}); err != nil {
		return err
	}

	return nil
}

func (s *chapterService) DeleteChapter(
	ctx context.Context,
	id string,
) (err error) {
	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return err
	}

	if ok, err := s.queries.CheckChapterByID(ctx, idUUID); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringChapter, err)
	} else if !ok {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrChapterNotFound)
	}

	if err = s.queries.SoftDeleteChapter(ctx, idUUID); err != nil {
		return
	}
	return
}

func (s *chapterService) Run(ctx context.Context, sessionID string, questView dto.QuestView) error {
	data, err := s.requestCompiler(sessionID, questView)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusInternalServerError, serviceErrors.ErrCompilerRunError)
	}

	return data
}

func (s *chapterService) requestCompiler(sessionID string, questView dto.QuestView) (*response.BaseResponse, error) {
	// nginx domain because we are inside of docker & i'm going to do load balancer.
	url := "http://nginx/compiler-api/v1/private/run"

	// Serialize questDTO to JSON
	requestBody, err := json.Marshal(questView)
	if err != nil {
		return nil, response.Response(500, "Error marshalling questDTO", err)
	}

	// Create a new POST request with the JSON body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, response.Response(500, "Error creating POST request", err)
	}

	// Set the Content-Type header to application/json
	req.Header.Add("Content-Type", "application/json")

	// Add the Codex-Compiler header
	req.Header.Add("Codex-Compiler", "b77759141fc85bf31e75b1d9c48bbe67")

	// Add the session_id cookie to the request
	req.AddCookie(&http.Cookie{
		Name:  "session_id",
		Value: sessionID,
	})

	// Create an HTTP client and execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, response.Response(500, "Error making POST request", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, response.Response(500, "Error reading response body", nil)
	}

	var data response.BaseResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessage(500, "Error decoding session data")
	}

	return &data, nil
}

func (s *chapterService) StartChapter(
	ctx context.Context,
	id, courseID, userAuthID string,
) error {
	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return err
	}
	courseUUID, err := s.utilService.NParseUUID(courseID)
	if err != nil {
		return err
	}

	// Comes From Session
	userAuthUUID := uuid.MustParse(userAuthID)

	if ok, err := s.queries.CheckChapterByID(ctx, idUUID); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringChapter, err)
	} else if !ok {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrChapterNotFound)
	}
	if ok, err := s.queries.CheckCourseByID(ctx, courseUUID); err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(serviceErrors.StatusInternalServerError, serviceErrors.ErrErrorWhileFilteringCourse, err)
	} else if !ok {
		return serviceErrors.NewServiceErrorWithMessage(serviceErrors.StatusBadRequest, serviceErrors.ErrCourseNotFound)
	}

	if err := s.queries.AddChapterToUser(ctx, repo.AddChapterToUserParams{
		ChapterID:  idUUID,
		CourseID:   courseUUID,
		UserAuthID: userAuthUUID,
	}); err != nil {
		return err
	}

	return nil
}
