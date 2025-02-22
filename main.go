package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SongDetail — структура, соответствующая схеме ответа внешнего API
type SongDetail struct {
	ReleaseDate string `json:"releaseDate"` // например, "16.07.2006"
	Text        string `json:"text"`
	Link        string `json:"link"`
}

// Для упрощения используем "mock-базу" — словарь, где ключом будет название группы (group),
// а значением — карта из названия песни (song) в SongDetail.
// Вы можете изменить логику, чтобы читать эти данные из файла или иной БД.
var mockDB = map[string]map[string]SongDetail{
	"Muse": {
		"Supermassive Black Hole": {
			ReleaseDate: "16.07.2006",
			Text: `Ooh baby, don't you know I suffer?
Ooh baby, can you hear me moan?
You caught me under false pretenses
How long before you let me go?

Ooh
You set my soul alight
Ooh
You set my soul alight`,
			Link: "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
		},
	},
	"SomeGroup": {
		"SomeSong": {
			ReleaseDate: "01.01.2025",
			Text:        "Some text here...",
			Link:        "https://www.youtube.com/watch?v=example",
		},
	},
}

func main() {
	router := gin.Default()

	// GET /info?group=...&song=...
	router.GET("/info", func(c *gin.Context) {
		group := c.Query("group")
		song := c.Query("song")

		// Проверяем, что оба параметра были переданы
		if group == "" || song == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required query params: group, song"})
			return
		}

		// Ищем группу в mockDB
		groupMap, ok := mockDB[group]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
			return
		}

		// Ищем песню внутри группы
		detail, ok := groupMap[song]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
			return
		}

		// Возвращаем SongDetail
		c.JSON(http.StatusOK, detail)
	})

	// Запускаем сервер на порту 8081
	log.Println("Внешнее API запущено на порту 8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
