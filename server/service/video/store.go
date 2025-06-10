package video

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/xadhithiyan/videon/config"
	awsconn "github.com/xadhithiyan/videon/service/awsConn"
	"github.com/xadhithiyan/videon/types"
)

type store struct {
	db *sql.DB
}

func CreateStore(db *sql.DB) *store {
	return &store{db: db}
}

func (s store) UploadS3(metaData types.MetaData, data []byte) error {

	var cancelFn func()
	ctx, cancelFn := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancelFn()

	key := metaData.Name + "/" + strconv.Itoa(metaData.Id)
	log.Print(key)

	_, err := awsconn.SVC.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(config.Env.AWSS3Bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
	})

	if err != nil {
		log.Print(err)
		return err
	}
	log.Print("successfully uploaded chuck: ", metaData.Id)

	return nil
}

func (s *store) AddVideoDB(userId int, metaData types.MetaData) error {

	rows, err := s.db.Query(`SELECT * FROM "video" WHERE name = $1`, metaData.Name)
	if err != nil {
		return err
	}

	u := new(types.VideoDB)
	for rows.Next() {
		err := rows.Scan(
			&u.ID,
			&u.UserId,
			&u.Name,
			&u.VideoType,
			&u.TotalChunks,
			&u.CompressVideo,
			&u.GeenrateThumbnail,
			&u.TranscodeVideo,
			&u.AddWaterMark,
			&u.VideoSummary,
		)

		if err != nil {
			return err
		}
	}
	if u.ID != 0 {
		return fmt.Errorf("video already exists on postgres")
	}

	_, err = s.db.Query(`  INSERT INTO video (userId, name, videoType, TotalChunks, 
	compressVideo, generateThumbnail, transcodeVideo, addWaterMark, videoSummary) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`, userId, metaData.Name, metaData.VideoType, metaData.TotalChunks,
		false, false, false, false, false)

	if err != nil {
		return err
	}

	log.Print("Added video to postgres")
	return nil
}
