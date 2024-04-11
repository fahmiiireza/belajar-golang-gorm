package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	controllers "github.com/Man4ct/belajar-golang-gorm/controllers/users"
	"github.com/Man4ct/belajar-golang-gorm/initializers"
	"github.com/Man4ct/belajar-golang-gorm/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository struct {
	*gorm.DB
}

// Book represents the Book model in the database.
type Book struct {
	ID          uint             `gorm:"primaryKey"`
	ISBN        string           `gorm:"unique;not null"`
	Title       string           `gorm:"not null"`
	Language    string           `gorm:"not null"`
	TotalCopy   int              `gorm:"not null"`
	Description string           // Add Description field
	CreatedBy   uint             `gorm:"not null"`
	Librarian   models.Librarian `gorm:"foreignKey:CreatedBy"`
}
type WordInfo struct {
	TF          int // Text Frequency
	DF          int // Document Frequency
	LastTF      int // Last Text Frequency
	LastDF      int // Last Document Frequency
	SearchCount int // Number of times the word was searched
}

var (
	wordInfoMap map[string]*WordInfo
	mutex       sync.Mutex
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
	wordInfoMap = make(map[string]*WordInfo)
}

func main() {

	r := gin.Default()
	r.POST("/user", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/user/:id", controllers.GetUser)
	r.PATCH("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	r.POST("/librarian", controllers.CreateLibrarian)
	r.GET("/librarian/:id", controllers.GetLibrarian)
	r.PATCH("/librarian/:id", controllers.UpdateLibrarian)

	r.GET("/search", searchHandler)

	// Start the server
	if err := r.Run(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func searchHandler(c *gin.Context) {
	// Parse request parameters
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing query parameter"})
		return
	}

	// Split query string into words
	words := strings.Fields(query)

	// Process each word concurrently
	var wg sync.WaitGroup
	wg.Add(len(words))
	for _, word := range words {
		go func(word string) {
			defer wg.Done()
			processWord(word)
		}(word)
	}
	wg.Wait()

	// Respond with word information
	c.JSON(http.StatusOK, wordInfoMap)
}

func processWord(word string) {
	mutex.Lock()
	defer mutex.Unlock()

	var books []Book
	err := initializers.DB.Model(&Book{}).Where("description ILIKE ?", "%"+word+"%").Find(&books).Error
	if err != nil {
		fmt.Printf("Error querying books: %v\n", err)
		return
	}

	tf := 0
	for _, book := range books {
		tf += strings.Count(strings.ToLower(book.Description), strings.ToLower(word))
	}

	info, exists := wordInfoMap[word]
	if !exists {
		info = &WordInfo{}
		wordInfoMap[word] = info
	}
	info.LastTF = info.TF
	info.LastDF = info.DF
	info.DF = len(books)
	info.TF = tf
	info.SearchCount++

}
