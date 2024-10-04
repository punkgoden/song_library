package storage

import (
	"context"
	"fmt"
	"github.com/punkgoden/song_library/app/internal/entity"
)

func (o *Storage) GetSongs(
	ctx context.Context,
	limit, offset int,
	order, orderField string,
) ([]*entity.Song, error) {
	var songs []*entity.Song
	tx := o.db.WithContext(ctx)

	var orderStr string
	if order == "ASC" || order == "DESC" {
		orderStr = orderField + " " + order
		tx = tx.Order(orderStr)
	}
	if err := tx.Find(&songs).Offset(offset).Limit(limit).Error; err != nil {
		return nil, err
	}
	return songs, nil
}

func (o *Storage) GetSong(
	ctx context.Context, name string,
) (*entity.Song, error) {
	tx := o.db.WithContext(ctx)
	var song *entity.Song
	if err := tx.Where(`name = ?`, name).
		Find(&song).Debug().Error; err != nil {
		return nil, err
	}
	return song, nil
}

func (o *Storage) GetTextSong(
	ctx context.Context, name string,
) (string, error) {
	tx := o.db.WithContext(ctx)
	var song *entity.Song
	if err := tx.Where(`name = ?`, name).
		Find(&song).Error; err != nil {
		return "", err
	}
	return song.Text, nil
}

func (o *Storage) DeleteSong(ctx context.Context, name string) error {
	tx := o.db.WithContext(ctx)
	var song *entity.Song
	if err := tx.Where(`name = ?`, name).
		Delete(&song).Error; err != nil {
		return err
	}
	return nil
}

func (o *Storage) UpdateSong(
	ctx context.Context, name, group, text, updSongName string,
) (*entity.Song, error) {
	tx := o.db.WithContext(ctx)
	var song entity.Song
	var songUpd entity.Song
	if name != "" {
		songUpd.Name = name
	}
	if group != "" {
		songUpd.Group = group
	}
	if text != "" {
		songUpd.Text = text
	}
	fmt.Println(songUpd)
	if err := tx.Model(&song).Where("name = ?", updSongName).Updates(songUpd).Error; err != nil {
		return nil, err
	}
	if err := tx.Model(&song).Where("name = ?", name).First(&song).Error; err != nil {
		return nil, err
	}
	return &song, nil
}

func (o *Storage) CreateSong(
	ctx context.Context, song entity.Song,
) (*entity.Song, error) {
	tx := o.db.WithContext(ctx)
	if err := tx.Create(&song).Error; err != nil {
		return nil, err
	}
	return &song, nil
}
