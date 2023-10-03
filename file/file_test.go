package file

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	// list, err := GetHistory()
	// log.Print(list, err)
	err := WhiteHistory(HistoryTiem{
		Type:    "text",
		Content: "test",
		Time:    1,
		IP:      "127.0.0.1",
	})
	log.Println(err)
	assert.Equal(t, 1, 0)
}

func TestHistoryList(t *testing.T) {
	list, err := GetHistory()
	log.Print(list, err)

	assert.Equal(t, 1, 0)
}
