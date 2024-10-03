package http_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/punkgoden/song_library/app/internal/dto"
	"net/http"
)

func (c *Controller) GetSongs(ctx *gin.Context) {
	var songRequestDTO dto.GetSongsRequestDTO
	if err := ctx.ShouldBind(&songRequestDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	songs, lastId, err := c.sc.GetSongs(ctx, &songRequestDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	songResponseDTO := dto.GetSongsResponseDTO{
		Songs:  songs,
		LastId: lastId,
	}
	ctx.JSON(http.StatusOK, &songResponseDTO)
}

func (c *Controller) GetSong(ctx *gin.Context) {
	var getSongRequestDTO dto.GetSongRequestDTO
	if err := ctx.ShouldBind(&getSongRequestDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	song, err := c.sc.GetSong(ctx, getSongRequestDTO.Name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	songResponseDTO := dto.GetSongResponseDTO{
		Song: song,
	}
	ctx.JSON(http.StatusOK, &songResponseDTO)
}

func (c *Controller) CreateSong(ctx *gin.Context) {
	var createSongRequestDTO dto.CreateSongRequestDTO
	if err := ctx.ShouldBind(&createSongRequestDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	song, err := c.sc.CreateSong(ctx, createSongRequestDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	songResponseDTO := dto.CreateAndUpdateSongResponseDTO{
		Song: song,
	}
	ctx.JSON(http.StatusCreated, &songResponseDTO)
}

func (c *Controller) UpdateSong(ctx *gin.Context) {
	var updateSongRequestDTO dto.UpdateSongRequestDTO
	if err := ctx.ShouldBind(&updateSongRequestDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	song, err := c.sc.UpdateSong(ctx, updateSongRequestDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	songResponseDTO := dto.CreateAndUpdateSongResponseDTO{
		Song: song,
	}
	ctx.JSON(http.StatusOK, &songResponseDTO)
}

func (c *Controller) DeleteSong(ctx *gin.Context) {
	var deleteSongRequestDTO dto.DeleteSongRequestDTO
	if err := ctx.ShouldBind(&deleteSongRequestDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.sc.DeleteSong(ctx, deleteSongRequestDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
