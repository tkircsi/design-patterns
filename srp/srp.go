package srp

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

type Journal struct {
	entries []string
	Sep     string
}

var entryCount = 0

// SRP code. Managing entries is the Journal responsibility.
func (j *Journal) Add(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) Remove(index int) {
	// some removal code
}

func (j *Journal) Get(index int) string {
	// return an entry at the index position
	return ""
}

func (j *Journal) String() string {
	return strings.Join(j.entries, j.Sep)
}

func NewJournal() *Journal {
	return &Journal{
		entries: []string{},
		Sep:     "\n",
	}
}

// separation of concerns
// Break SRP. Not Journal related functions.
func (j *Journal) Save(filename string) {
	_ = os.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
	// loading from file
}

func (j *Journal) LoadFromWeb(url *url.URL) {
	// loading from web
}

func Main() {
	j := NewJournal()
	j.Add("entry 1")
	j.Add("entry 2")
	fmt.Println(j)
}
