package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type BoardRequest struct {
	Board [][]string `json:"board"`
}

type WordsResponse struct {
	Words []string `json:"words"`
}

var dictionary = map[string]bool{
	"cresset": true, "cress": true, "set": true, "chypres": true, "song": true, "go": true,
}

// Handles the input boggle board request
func HandleBoggleBoardSubmission(ctx *gin.Context) {
	var req BoardRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	words := solveBoggleBoard(req.Board)
	ctx.JSON(http.StatusOK, WordsResponse{Words: words})
}

// Solves the Boggle board in order to find all valid words
func solveBoggleBoard(board [][]string) []string {
	rows := len(board)
	cols := len(board[0])
	wordsFound := make(map[string]bool) 
	uniqueWords := make(map[string]bool)
	var validWords []string

	// Dfs algorithm for traversing through the boggle board horizontally, vertically, and diagonally.
	// This is where word paths are validated and "wordsFound" is updated
	var dfs func(row, col int, path string, visited [][]bool)
	dfs = func(row, col int, path string, visited[][]bool) {
		if row < 0 || row >= rows || col < 0 || col >= cols || visited[row][col] {
			return
		}

		path += board[row][col]
		if !dictionary[path] && !hasPrefix(path) {
			return
		}

		if len(path) >= 3 && dictionary[path] {
			wordsFound[path] = true
		}

		visited[row][col] = true
		directions := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
		for _, dir := range directions {
			dfs(row + dir[0], col + dir[1], path, visited)
		}
		visited[row][col] = false
	}
	
	// This is important so that every cell of the board is considered as a potential valid path.
	// All possible paths and word combinations are explored while keeping track of already visited cells
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			visited := make([][]bool, rows)
			for k := range visited {
				visited[k] = make([]bool, cols)
			}
			dfs(i, j, "", visited)
		}
	}

	// Making sure that duplicated words are not added to the final result due to one of the rules of boggle
	for word := range wordsFound {
		if !uniqueWords[word] {
            uniqueWords[word] = true
            validWords = append(validWords, word)
        }
	}

	return validWords
}

// Helps to optimize the checking of valid words in the dictionary when traversing paths through the board 
func hasPrefix(prefix string) bool {
	for word := range dictionary {
		if strings.HasPrefix(word, prefix) {
			return true
		}
	}

	return false
}