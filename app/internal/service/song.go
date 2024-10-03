package service

import (
	"context"
	"github.com/punkgoden/song_library/app/internal/entity"
)

func (s *Service) GetSongs(ctx context.Context, limit, lastId int) ([]*entity.Song, error) {
	songs, err := s.st.GetSongs(ctx, limit, lastId)
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (s *Service) GetSong(ctx context.Context, name string) (*entity.Song, error) {
	song, err := s.st.GetSong(ctx, name)
	if err != nil {
		return nil, err
	}
	return song, nil
}

func (s *Service) DeleteSong(ctx context.Context, name string) error {
	if err := s.st.DeleteSong(ctx, name); err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateSong(ctx context.Context, name, group string) (*entity.Song, error) {
	song, err := s.st.UpdateSong(ctx, name, group)
	if err != nil {
		return nil, err
	}
	return song, err
}

func (s *Service) CreateSong(ctx context.Context, createSong entity.Song) (*entity.Song, error) {
	song, err := s.st.CreateSong(ctx, createSong)
	if err != nil {
		return nil, err
	}
	return song, nil
}
