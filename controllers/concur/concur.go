/*
Create an endpoint that receives as input a list of words and returns the number of occurrences
(TF: text frequency & DF: document frequency) of these words in a collection of files already stored
on the server.

The program should store, for each word, :
the number of times the users searched for it, the last TF and the last DF.
the TF and the DF of each search made on it
 TF: the number of occurrences in all files
 DF: the number of documents in which the word was found.
*/

package concur

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/Man4ct/belajar-golang-gorm/initializers"
	"github.com/Man4ct/belajar-golang-gorm/models"
	"github.com/gin-gonic/gin"
)

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
	wordInfoMap = make(map[string]*WordInfo)
}

func SearchHandler(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing query parameter"})
		return
	}

	words := strings.Fields(query)

	var wg sync.WaitGroup
	wg.Add(len(words))
	for _, word := range words {
		go func(word string) {
			defer wg.Done()
			processWord(word)
		}(word)
	}
	wg.Wait()

	c.JSON(http.StatusOK, wordInfoMap)
}

func processWord(word string) {
	mutex.Lock()
	defer mutex.Unlock()

	var books []models.Book
	err := initializers.DB.Model(&models.Book{}).Where("description ILIKE ?", "%"+word+"%").Find(&books).Error
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
