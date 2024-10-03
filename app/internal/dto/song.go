package dto

import "github.com/punkgoden/song_library/app/internal/entity"

type GetSongsRequestDTO struct {
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
	Order      string `json:"order,omitempty"`
	OrderField string `json:"order_field,omitempty"`
}

type GetSongsResponseDTO struct {
	Songs []*entity.Song `json:"songs"`
}

type GetSongRequestDTO struct {
	Name string `json:"name"`
}

type GetSongResponseDTO struct {
	Song *entity.Song `json:"song"`
}

type CreateSongRequestDTO struct {
	Name       string `json:"name"`
	Group      string `json:"group"`
	Text       string `json:"text,omitempty"`
	Listenings int    `json:"listenings,omitempty"`
}

type UpdateSongRequestDTO struct {
	Name  string `json:"name"`
	Group string `json:"group,omitempty"`
	Text  string `json:"text,omitempty"`
}

type CreateAndUpdateSongResponseDTO struct {
	Song *entity.Song `json:"song"`
}
type DeleteSongRequestDTO struct {
	Name string `json:"name"`
}
