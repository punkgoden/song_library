package dto

import "github.com/punkgoden/song_library/app/internal/entity"

type GetSongsRequestDTO struct {
	Limit      int    `json:"limit"`
	LastId     int    `json:"last_id"`
	Order      string `json:"order,omitempty"`
	OrderField string `json:"order_field,omitempty"`
}

type GetSongsResponseDTO struct {
	LastId int            `json:"last_id"`
	Songs  []*entity.Song `json:"songs"`
}

type GetSongRequestDTO struct {
	Name string `json:"song"`
}

type GetSongResponseDTO struct {
	Song *entity.Song `json:"song"`
}

type CreateSongRequestDTO struct {
	Name       string `json:"song"`
	Group      string `json:"group"`
	Text       string `json:"text,omitempty"`
	Listenings int    `json:"listenings,omitempty"`
}

type UpdateSongRequestDTO struct {
	Name  string `json:"song,omitempty"`
	Group string `json:"group,omitempty"`
	Text  string `json:"text,omitempty"`
}

type CreateAndUpdateSongResponseDTO struct {
	Song *entity.Song `json:"song"`
}
type DeleteSongRequestDTO struct {
	Name string `json:"song"`
}
