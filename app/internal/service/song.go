package service

import (
	"context"
	"github.com/punkgoden/song_library/app/internal/dto"
	"github.com/punkgoden/song_library/app/internal/entity"
)

func (s *Service) GetSongs(
	ctx context.Context,
	songDTO *dto.GetSongsRequestDTO,
) ([]*entity.Song, error) {
	songs, err := s.st.GetSongs(
		ctx,
		songDTO.Limit,
		songDTO.Offset,
		songDTO.Order,
		songDTO.OrderField,
	)
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

func (s *Service) GetTextSong(ctx context.Context, name string) (string, error) {
	textSong, err := s.st.GetTextSong(ctx, name)
	if err != nil {
		return "", err
	}
	return textSong, nil
}

func (s *Service) DeleteSong(
	ctx context.Context,
	deleteRequestDTO dto.DeleteSongRequestDTO,
) error {
	if err := s.st.DeleteSong(ctx, deleteRequestDTO.Name); err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateSong(
	ctx context.Context,
	updateSongRequestDTO dto.UpdateSongRequestDTO,
	songName string,
) (*entity.Song, error) {
	song, err := s.st.UpdateSong(
		ctx,
		updateSongRequestDTO.Name,
		updateSongRequestDTO.Group,
		updateSongRequestDTO.Text,
		songName,
	)
	if err != nil {
		return nil, err
	}
	return song, err
}

func (s *Service) CreateSong(
	ctx context.Context,
	createSongDTO dto.CreateSongRequestDTO,
) (*entity.Song, error) {
	createSong := entity.Song{
		Name:       createSongDTO.Name,
		Group:      createSongDTO.Group,
		Text:       createSongDTO.Text,
		Listenings: createSongDTO.Listenings,
	}
	song, err := s.st.CreateSong(ctx, createSong)
	if err != nil {
		return nil, err
	}
	return song, nil
}
