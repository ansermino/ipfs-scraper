package scraper

import (
	"testing"
)

const TEST_URL = "https://en.wikipedia.org/wiki/InterPlanetary_File_System"

func TestFetchPage(t *testing.T) {
	_, err := FetchPage(TEST_URL)
	if err != nil {
		t.Fatal(err)
	}
}
