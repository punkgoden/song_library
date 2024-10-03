package storage

import (
	"context"
	"github.com/punkgoden/song_library/app/internal/entity"
)

func (o *Storage) GetSongs(ctx context.Context, limit, lastId int) ([]*entity.Song, error) {
	tx := o.db.WithContext(ctx)
	var songs []*entity.Song
	if err := tx.Model(songs).Where("id > ?", lastId).Limit(limit).
		Find(&songs).Error; err != nil {
		return nil, err
	}
	return songs, nil
}

func (o *Storage) GetSong(
	ctx context.Context, name string,
) (*entity.Song, error) {
	tx := o.db.WithContext(ctx)
	var song *entity.Song
	if err := tx.Model(song).Where(`name = ?`, name).
		Find(&song).Error; err != nil {
		return nil, err
	}
	return song, nil
}

func (o *Storage) DeleteSong(ctx context.Context, name string) error {
	tx := o.db.WithContext(ctx)
	var song *entity.Song
	if err := tx.Model(song).Where(`name = ?`, name).
		Delete(&song).Error; err != nil {
		return err
	}
	return nil
}

func (o *Storage) UpdateSong(
	ctx context.Context, name, group string,
) (*entity.Song, error) {
	tx := o.db.WithContext(ctx)
	var song *entity.Song
	var songUpd entity.Song
	if name != "" {
		songUpd.Name = name
	}
	if group != "" {
		songUpd.Group = group
	}
	if err := tx.Model(&song).Where("name = ?", name).
		Updates(songUpd).Find(&song).Error; err != nil {
		return nil, err
	}
	return song, nil
}

func (o *Storage) CreateSong(
	ctx context.Context, song entity.Song,
) (*entity.Song, error) {
	tx := o.db.WithContext(ctx)
	if err := tx.Model(&song).Create(&song).Error; err != nil {
		return nil, err
	}
	return &song, nil
}
