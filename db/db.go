package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"ipfs-scraper/config"
	"ipfs-scraper/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ Database = &mongoStore{}

var ErrNotFound = errors.New("document not found")

type mongoStore struct {
	client          *mongo.Client
	pageInfoColl    *mongo.Collection
	pageVersionColl *mongo.Collection
	timeout         time.Duration
}

type Database interface {
	CreatePageInfo(ctx context.Context, pg *models.PageInfo) error
	ViewPageInfo(ctx context.Context, pg *models.PageInfo) (*models.PageInfo, error)
	UpsertPageInfo(ctx context.Context, pg *models.PageInfo) error
	ViewAllPages(ctx context.Context) ([]*models.PageInfo, error)
	CreatePageVersion(ctx context.Context, pv *models.PageVersion) error
	ViewPageVersion(ctx context.Context, pv *models.PageVersion) (*models.PageVersion, error)
	ViewPageVersions(ctx context.Context, pageId string) ([]*models.PageVersion, error)
}

func New(cfg *config.Database) (Database, error) {
	timeout := time.Duration(5) * time.Second
	opts := options.Client()

	opts.ApplyURI(cfg.URI)
	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, fmt.Errorf("connecton error [%s]: %w", opts.GetURI(), err)
	}

	// connect to the mongoDB cluster
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// test the connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	} else {
		log.Info().Str("addr", cfg.URI).Msg("Successfully connected to mongodb")
	}

	return &mongoStore{
		client:          client,
		pageInfoColl:    client.Database(cfg.Database).Collection("pageInfo"),
		pageVersionColl: client.Database(cfg.Database).Collection("pageVersions"),
		timeout:         timeout,
	}, nil
}

func (s *mongoStore) UpsertPageInfo(ctx context.Context, pg *models.PageInfo) error {
	_, err := s.ViewPageInfo(ctx, pg)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return s.CreatePageInfo(ctx, pg)
		}
		return err
	}

	return s.UpdatePageInfo(ctx, pg)
}

func (s *mongoStore) CreatePageInfo(ctx context.Context, pg *models.PageInfo) error {
	_, err := s.ViewPageInfo(ctx, pg)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			_, err = s.pageInfoColl.InsertOne(ctx, pg)
			return err
		}
		return err
	}
	return nil
}

func (s *mongoStore) UpdatePageInfo(ctx context.Context, pg *models.PageInfo) error {
	filter := bson.D{
		{Key: "url", Value: pg.Url},
	}
	_, err := s.pageInfoColl.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: pg}})
	if err != nil {
		return err
	}
	return nil
}

func (s *mongoStore) ViewPageInfo(ctx context.Context, pg *models.PageInfo) (*models.PageInfo, error) {
	filter := bson.D{
		{Key: "url", Value: pg.Url},
	}
	res := new(models.PageInfo)
	err := s.pageInfoColl.FindOne(ctx, filter).Decode(res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

func (s *mongoStore) ViewAllPages(ctx context.Context) ([]*models.PageInfo, error) {
	var versions []*models.PageInfo
	cursor, err := s.pageInfoColl.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		// create a value into which the single document can be decoded
		page := new(models.PageInfo)
		err := cursor.Decode(page)
		if err != nil {
			return nil, err
		}

		versions = append(versions, page)
	}
	return versions, nil
}

func (s *mongoStore) CreatePageVersion(ctx context.Context, pv *models.PageVersion) error {
	//_, err := s.ViewPageVersion(ctx, pv)
	//if err != nil {
	//	if errors.Is(err, ErrNotFound) {
	_, err := s.pageVersionColl.InsertOne(ctx, pv)
	return err
	//	}
	//	return err
	//}
	//return nil
}

func (s *mongoStore) ViewPageVersion(ctx context.Context, pv *models.PageVersion) (*models.PageVersion, error) {
	filter := bson.D{
		{Key: "url", Value: pv.Url},
	}
	res := new(models.PageVersion)
	err := s.pageVersionColl.FindOne(ctx, filter).Decode(res)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

func (s *mongoStore) ViewPageVersions(ctx context.Context, title string) ([]*models.PageVersion, error) {
	var versions []*models.PageVersion
	cursor, err := s.pageVersionColl.Find(ctx, bson.D{{Key: "title", Value: title}})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		// create a value into which the single document can be decoded
		peer := new(models.PageVersion)
		err := cursor.Decode(peer)
		if err != nil {
			return nil, err
		}

		versions = append(versions, peer)
	}
	return versions, nil
}

//func (s *mongoStore) ViewAll(ctx context.Context) ([]*models.Peer, error) {
//	var peers []*models.Peer
//	cursor, err := s.pageInfoColl.Find(ctx, bson.D{{Key: "is_connectable", Value: true}})
//	if err != nil {
//		return nil, err
//	}
//
//	for cursor.Next(ctx) {
//		// create a value into which the single document can be decoded
//		peer := new(models.Peer)
//		err := cursor.Decode(peer)
//		if err != nil {
//			return nil, err
//		}
//
//		peers = append(peers, peer)
//	}
//	return peers, nil
//}
