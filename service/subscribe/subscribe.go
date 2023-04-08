package subscribe

import (
	"context"
	"fmt"
	"youtube-clone/pkg/repository/dto"
	"youtube-clone/pkg/repository/gormadp"
)

var _ Interface = (*Subscribe)(nil)

type Subscribe struct {
	repo *gormadp.Repository
}

func New(r *gormadp.Repository) (*Subscribe, error) {
	if r == nil {
		return nil, fmt.Errorf("new subscribe service : %w", ErrRepositoryShouldBeSetUp)
	}

	return &Subscribe{
		repo: r,
	}, nil
}

func (s *Subscribe) Subscribe(ctx context.Context, userID, channelID uint) error {
	var err error

	if _, err = s.repo.UserRepo.GetUserByID(ctx, channelID); err != nil {
		return fmt.Errorf("subscribe : %w", err)
	}

	if err = s.repo.SubscribeRepo.Subscribe(ctx, userID, channelID); err != nil {
		return fmt.Errorf("subscribe : %w", err)
	}

	return nil
}

func (s *Subscribe) Unsubscribe(ctx context.Context, userID, channelID uint) error {
	var err error

	if _, err = s.repo.UserRepo.GetUserByID(ctx, channelID); err != nil {
		return fmt.Errorf("unsubscribe : %w", err)
	}

	if err = s.repo.SubscribeRepo.UnSubscribe(ctx, userID, channelID); err != nil {
		return fmt.Errorf("unsubscribe : %w", err)
	}

	return nil
}

func (s *Subscribe) GetSubscribers(ctx context.Context, channelID uint, limit, offset int) (*SubList, error) {
	var (
		subs *dto.SubscribeList
		err  error
	)

	if subs, err = s.repo.SubscribeRepo.GetSubscribers(ctx, channelID, limit, offset); err != nil {
		return nil, fmt.Errorf("get subscribers : %w", err)
	}

	subList := make([]*Sub, 0, len(subs.Subscribes))

	for i := range subs.Subscribes {
		var (
			user *dto.User
			err  error
		)

		if user, err = s.repo.UserRepo.GetUserByID(ctx, subs.Subscribes[i].UserID); err != nil {
			return nil, fmt.Errorf("get subscribers : %w", err)
		}

		subList = append(subList, &Sub{
			ID:       user.ID,
			FullName: user.FullName,
		})
	}

	return &SubList{
		Limit:  subs.Limit,
		Offset: subs.Offset,
		Count:  subs.Count,
		Subs:   subList,
	}, nil
}

func (s *Subscribe) GetSubscribeList(ctx context.Context, userID uint, limit, offset int) (*SubList, error) {
	var (
		subs *dto.SubscribeList
		err  error
	)

	if subs, err = s.repo.SubscribeRepo.GetSubscribeList(ctx, userID, limit, offset); err != nil {
		return nil, fmt.Errorf("get subscribe list : %w", err)
	}

	subList := make([]*Sub, 0, len(subs.Subscribes))

	for i := range subs.Subscribes {
		var (
			channel *dto.User
			err     error
		)

		if channel, err = s.repo.UserRepo.GetUserByID(ctx, subs.Subscribes[i].ChannelID); err != nil {
			return nil, fmt.Errorf("get subscribe list : %w", err)
		}

		subList = append(subList, &Sub{
			ID:       channel.ID,
			FullName: channel.FullName,
		})
	}

	return &SubList{
		Limit:  subs.Limit,
		Offset: subs.Offset,
		Count:  subs.Count,
		Subs:   subList,
	}, nil
}
