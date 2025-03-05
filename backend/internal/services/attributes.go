package services

import (
	"context"
	"database/sql"
	"strconv"
	"strings"

	"github.com/C-dexTeam/codex/internal/domains"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	repo "github.com/C-dexTeam/codex/internal/repos/out"
	"github.com/google/uuid"
)

type attributeService struct {
	db          *sql.DB
	queries     *repo.Queries
	utilService IUtilService
}

func NewAttributeService(
	db *sql.DB,
	queries *repo.Queries,
	utilService IUtilService,
) *attributeService {
	return &attributeService{
		db:          db,
		queries:     queries,
		utilService: utilService,
	}
}

func (s *attributeService) GetAttributes(
	ctx context.Context,
	id, rewardID, traitType, page, limit string,
) ([]domains.Attribute, error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = s.utilService.D().Limits.DefaultAttributeLimit
	}

	if _, err := s.utilService.ParseUUID(id); err != nil {
		return nil, err
	}
	if _, err := s.utilService.ParseUUID(id); err != nil {
		return nil, err
	}

	attributes, err := s.queries.GetAttributes(ctx, repo.GetAttributesParams{
		ID:        s.utilService.ParseNullUUID(id),
		RewardID:  s.utilService.ParseNullUUID(rewardID),
		TraitType: s.utilService.ParseString(traitType),
		Lim:       int32(limitNum),
		Off:       (int32(pageNum) - 1) * int32(limitNum),
	})
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringRewardsAttributes,
			err,
		)
	}
	domainsAttr := domains.NewAttributes(attributes)

	return domainsAttr, nil
}

func (s *attributeService) AddAttribute(
	ctx context.Context,
	rewardID, traitType, value string,
) (uuid.UUID, error) {
	rewardUUID, err := s.utilService.ParseUUID(rewardID)
	if err != nil {
		return uuid.Nil, err
	}

	id, err := s.queries.CreateAttribute(ctx, repo.CreateAttributeParams{
		RewardID:  rewardUUID,
		TraitType: traitType,
		Value:     value,
	})
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (s *attributeService) UpdateAttribute(
	ctx context.Context,
	id, rewardID, traitType, value string,
) error {
	idUUID, err := s.utilService.ParseUUID(id)
	if err != nil {
		return err
	}
	if _, err := s.utilService.ParseUUID(rewardID); err != nil {
		return err
	}

	_, err = s.queries.GetAttributeByID(ctx, idUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return serviceErrors.NewServiceErrorWithMessage(
				serviceErrors.StatusBadRequest,
				serviceErrors.ErrRewardAttributeNotFound,
			)
		}
		return serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringRewardsAttributes,
			err,
		)
	}

	if err := s.queries.UpdateAttribute(ctx, repo.UpdateAttributeParams{
		AttributeID: idUUID,
		RewardID:    s.utilService.ParseNullUUID(rewardID),
		TraitType:   s.utilService.ParseString(traitType),
		Value:       s.utilService.ParseString(value),
	}); err != nil {
		return err
	}

	return nil
}

func (s *attributeService) DeleteAttribute(ctx context.Context, id string) (err error) {

	idUUID, err := s.utilService.NParseUUID(id)
	if err != nil {
		return err
	}

	_, err = s.queries.GetAttributeByID(ctx, idUUID)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return serviceErrors.NewServiceErrorWithMessage(
				serviceErrors.StatusBadRequest,
				serviceErrors.ErrRewardAttributeNotFound,
			)
		}
		return serviceErrors.NewServiceErrorWithMessageAndError(
			serviceErrors.StatusInternalServerError,
			serviceErrors.ErrErrorWhileFilteringRewardsAttributes,
			err,
		)
	}

	if err := s.queries.DeleteAttribute(ctx, idUUID); err != nil {
		return err
	}

	return
}
