package http_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/punkgoden/song_library/app/internal/dto"
	"net/http"
	"strconv"
)

// GetSongs godoc
//
//	@Summary		Get Songs
//	@Description	Returns songs
//	@Tags			song
//	@Accept			json
//	@Produce		json
//	@Param   		limit     		query   int 	true    "limit"
//	@Param   		offset     		query   int 	true    "offset"
//	@Param   		order     		query   string 	false   "order" 		Enums(ASC, DESC)
//	@Param   		order_field     query   string 	false   "order_field"
//	@Success		200		{object}		dto.GetSongsResponseDTO
//	@Router			/songs [get]
func (c *Controller) GetSongs(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	offset, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	songRequestDTO := dto.GetSongsRequestDTO{
		Limit:      limit,
		Offset:     offset,
		Order:      ctx.Query("order"),
		OrderField: ctx.Query("order_field"),
	}
	if err := ctx.ShouldBind(&songRequestDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	songs, err := c.sc.GetSongs(ctx, &songRequestDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	songResponseDTO := dto.GetSongsResponseDTO{
		Songs: songs,
	}
	ctx.JSON(http.StatusOK, &songResponseDTO)
}

// GetSong godoc
//
//	@Summary		Get Song
//	@Description	Returns song by name
//	@Tags			song
//	@Accept			json
//	@Produce		json
//	@Param   		name     query   string 	true   "Get song by name"
//	@Success		200		{object}		dto.GetSongResponseDTO
//	@Router			/song [get]
func (c *Controller) GetSong(ctx *gin.Context) {
	getSongRequestDTO := dto.GetSongRequestDTO{
		Name: ctx.Query("name"),
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

// GetTextSong godoc
//
//	@Summary		Get Text Song
//	@Description	Returns song by name
//	@Tags			song
//	@Accept			json
//	@Produce		json
//	@Param   		name     query   string 	true   "Get text song by name"
//	@Success		200		{object}		string
//	@Router			/text-song [get]
func (c *Controller) GetTextSong(ctx *gin.Context) {
	getSongRequestDTO := dto.GetSongRequestDTO{
		Name: ctx.Query("name"),
	}
	textSong, err := c.sc.GetTextSong(ctx, getSongRequestDTO.Name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"text": textSong})
}

// CreateSong godoc
//
//	@Summary		Create song
//	@Description	Returns song
//	@Tags			song
//	@Accept  json
//	@Produce  json
//	@Param   		song     	body   dto.CreateSongRequestDTO 	true    "Create Song"
//	@Success		201		{object}		dto.CreateAndUpdateSongResponseDTO
//	@Router			/song [post]
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

// UpdateSong godoc
//
//	@Summary		Update song
//	@Description	Returns song
//	@Tags			song
//	@Accept  json
//	@Produce  json
//	@Param   		song     	body   dto.UpdateSongRequestDTO 	true    "Update Song"
//	@Param   		name     	query   string 						true   	"Update song name"
//	@Success		200		{object}		dto.CreateAndUpdateSongResponseDTO
//	@Router			/song [patch]
func (c *Controller) UpdateSong(ctx *gin.Context) {
	var updateSongRequestDTO dto.UpdateSongRequestDTO
	if err := ctx.ShouldBind(&updateSongRequestDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	songName := ctx.Query("name")
	fmt.Println(songName)
	song, err := c.sc.UpdateSong(ctx, updateSongRequestDTO, songName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	songResponseDTO := dto.CreateAndUpdateSongResponseDTO{
		Song: song,
	}
	ctx.JSON(http.StatusOK, &songResponseDTO)
}

// DeleteSong godoc
//
//	@Summary		Delete song
//	@Description	Delete song by song name
//	@Tags			song
//	@Accept			json
//	@Produce		json
//	@Param   		song     	body   dto.DeleteSongRequestDTO 	true    "Delete Song"
//	@Success		204
//	@Router			/song [delete]
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
