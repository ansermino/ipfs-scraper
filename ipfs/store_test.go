package ipfs

import "testing"

func TestStoreDir(t *testing.T) {
	_, err := StoreDir("/ip4/127.0.0.1/tcp/5001", "../scraper/en.wikipedia.org")
	if err != nil {
		t.Fatal(err)
	}
}
