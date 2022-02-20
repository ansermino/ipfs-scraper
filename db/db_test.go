package db

import (
	"context"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/assert"
	"ipfs-scraper/config"
	"ipfs-scraper/models"
	"testing"
	"time"
)

var TestPage = &models.PageInfo{
	Url:   "http://website.com",
	Title: "A page",
}

var UpdatedTestPage = &models.PageInfo{
	Url:   "http://website.com",
	Title: "NewPage",
}

var testTime = time.Now()
var testCid, _ = cid.Decode("QmbWqxBEKC3P8tqsKc98xmWNzrzDtRLMiMPL8wBuTGsMnR")

var TestPageVersionA = &models.PageVersion{
	Url:       "1",
	Timestamp: testTime,
	Cid:       testCid,
}

var TestPageVersionB = &models.PageVersion{
	Url:       "1",
	Timestamp: testTime.Add(time.Minute),
	Cid:       testCid,
}

func TestMongoStore_CreateViewUpsertPageInfo(t *testing.T) {
	cfg := &config.Database{
		URI:      "mongodb://root:rootpassword@127.0.0.1:27017",
		Database: "testdb",
	}

	db, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}

	err = db.CreatePageInfo(context.TODO(), TestPage)
	if err != nil {
		t.Fatal(err)
	}

	res, err := db.ViewPageInfo(context.TODO(), TestPage)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, TestPage, res)

	err = db.UpsertPageInfo(context.TODO(), UpdatedTestPage)
	assert.NoError(t, err)

	res, err = db.ViewPageInfo(context.TODO(), TestPage)
	assert.NoError(t, err)
	assert.Equal(t, UpdatedTestPage, res)
}

func TestMongoStore_CreateViewViewAllPageVersion(t *testing.T) {
	cfg := &config.Database{
		URI:      "mongodb://root:rootpassword@127.0.0.1:27017",
		Database: "testdb",
	}

	db, err := New(cfg)
	assert.NoError(t, err)

	err = db.CreatePageVersion(context.TODO(), TestPageVersionA)
	assert.NoError(t, err)

	err = db.CreatePageVersion(context.TODO(), TestPageVersionB)
	assert.NoError(t, err)

	versions, err := db.ViewPageVersions(context.TODO(), "1")
	assert.NoError(t, err)

	assert.Equal(t, []*models.PageVersion{TestPageVersionA, TestPageVersionB}, versions)

	version, err := db.ViewPageVersion(context.TODO(), TestPageVersionA)
	assert.NoError(t, err)
	assert.EqualValues(t, TestPageVersionA, version)
}
