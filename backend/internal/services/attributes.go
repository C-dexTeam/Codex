package services

import (
	"context"
	"strconv"

	"github.com/C-dexTeam/codex/internal/domains"
	errorDomains "github.com/C-dexTeam/codex/internal/domains/errors"
	serviceErrors "github.com/C-dexTeam/codex/internal/errors"
	"github.com/google/uuid"
)

type attributeService struct {
	attributeRepository domains.IAttributeRepository
}

func NewAttributeService(attributeRepository domains.IAttributeRepository) domains.IAttributeService {
	return &attributeService{
		attributeRepository: attributeRepository,
	}
}

func (s *attributeService) GetAttributes(ctx context.Context, id, rewardID, traitType, page, limit string) (attributes []domains.Attribute, err error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil || page == "" {
		pageNum = 1
	}

	limitNum, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limitNum = domains.DefaultAttributeLimit
	}

	var (
		attributeUUID uuid.UUID
		rewardUUID    uuid.UUID
	)
	if id != "" {
		attributeUUID, err = uuid.Parse(id)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
		}
	}
	if rewardID != "" {
		rewardUUID, err = uuid.Parse(rewardID)
		if err != nil {
			return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
		}
	}

	attributes, _, err = s.attributeRepository.Filter(ctx, domains.AttributeFilter{
		ID:        attributeUUID,
		RewardID:  rewardUUID,
		TraitType: traitType,
	}, int64(limitNum), int64(pageNum))
	if err != nil {
		return nil, serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringRewardsAttributes, err)
	}

	return attributes, nil
}

func (s *attributeService) AddAttribute(
	ctx context.Context,
	rewardID, traitType, value string,
) (uuid.UUID, error) {
	newAttribute, err := domains.NewAttribute(
		"",
		rewardID,
		traitType,
		value,
	)
	if err != nil {
		return uuid.Nil, err
	}

	id, err := s.attributeRepository.Add(ctx, newAttribute)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (s *attributeService) UpdateAttribute(
	ctx context.Context,
	id, rewardID, traitType, value string,
) error {
	var idUUID uuid.UUID
	idUUID, err := uuid.Parse(id)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
	}

	attributes, _, err := s.attributeRepository.Filter(ctx, domains.AttributeFilter{
		ID: idUUID,
	}, 1, 1)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringRewardsAttributes, err)
	}
	if len(attributes) != 1 {
		return serviceErrors.NewServiceErrorWithMessage(errorDomains.StatusNotFound, errorDomains.ErrRewardAttributeNotFound)
	}

	updateAttribute, err := domains.NewAttribute(
		id,
		rewardID,
		traitType,
		value,
	)
	if err != nil {
		return err
	}

	if err := s.attributeRepository.Update(ctx, updateAttribute); err != nil {
		return err
	}

	return nil
}

func (s *attributeService) DeleteAttribute(ctx context.Context, attributeID string) (err error) {
	var attributeUUID uuid.UUID
	attributeUUID, err = uuid.Parse(attributeID)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrInvalidID, err)
	}

	attributes, _, err := s.attributeRepository.Filter(ctx, domains.AttributeFilter{
		ID: attributeUUID,
	}, 1, 1)
	if err != nil {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusInternalServerError, errorDomains.ErrErrorWhileFilteringRewardsAttributes, err)
	}
	if len(attributes) != 1 {
		return serviceErrors.NewServiceErrorWithMessageAndError(errorDomains.StatusBadRequest, errorDomains.ErrRewardAttributeNotFound, err)
	}

	if err = s.attributeRepository.Delete(ctx, attributeUUID); err != nil {
		return
	}
	return
}
