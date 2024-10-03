package service

import (
	"context"
	"github.com/punkgoden/song_library/app/internal/dto"
	"github.com/punkgoden/song_library/app/internal/entity"
)

func (s *Service) GetSongs(
	ctx context.Context,
	songDTO *dto.GetSongsRequestDTO,
) ([]*entity.Song, int, error) {
	songs, err := s.st.GetSongs(
		ctx,
		songDTO.Limit,
		songDTO.LastId,
		songDTO.Order,
		songDTO.OrderField,
	)
	if err != nil {
		return nil, 0, err
	}
	newLastId := songs[len(songs)-1].ID
	return songs, newLastId, nil
}

func (s *Service) GetSong(ctx context.Context, name string) (*entity.Song, error) {
	song, err := s.st.GetSong(ctx, name)
	if err != nil {
		return nil, err
	}
	return song, nil
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
) (*entity.Song, error) {
	song, err := s.st.UpdateSong(
		ctx,
		updateSongRequestDTO.Name,
		updateSongRequestDTO.Group,
		updateSongRequestDTO.Text,
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
