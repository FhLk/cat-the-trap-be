package controller

//import (
//	"cat-the-trap-back-end/service"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func RewardInfo(c *gin.Context) {
//	albums, err := service.ListAllAlbum()
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	c.IndentedJSON(http.StatusOK, albums)
//}
//
//func Play(c *gin.Context) {
//	id := c.Param("id")
//	album, err := service.AlbumByID(id)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	c.IndentedJSON(http.StatusOK, album)
//}
