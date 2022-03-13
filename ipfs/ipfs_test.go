package ipfs

import "testing"

func TestPing(t *testing.T) {
	err := PingNode("/ip4/127.0.0.1/tcp/5001")
	if err != nil {
		t.Fatal(err)
	}

}
