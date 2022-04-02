package sqlstorage

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

const configFile = "config.yaml"

func TestStorage(t *testing.T) {
	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		t.Skip(configFile + " file does not exists")
	}

	configContent, _ := os.ReadFile(configFile)

	var config struct {
		Storage struct {
			URL string
		}
	}

	err := yaml.Unmarshal(configContent, &config)
	if err != nil {
		t.Fatal("Failed to unmarshal config", err)
	}

	ctx := context.Background()
	storage := New(ctx, config.Storage.URL)
	if err := storage.Connect(ctx); err != nil {
		t.Fatal("Failed to connect to DB server", err)
	}

	t.Run("test SQL", func(t *testing.T) {
		tx, err := storage.conn.BeginTx(ctx, pgx.TxOptions{
			IsoLevel:       pgx.Serializable,
			AccessMode:     pgx.ReadWrite,
			DeferrableMode: pgx.NotDeferrable,
		})
		if err != nil {
			t.Fatal("Failed to connect to DB server", err)
		}

		findAccout, err := storage.FindAccount("test1")
		if err != nil {
			t.FailNow()
			return
		}
		fmt.Println(findAccout)

		//userID := uuid.New()
		//postID, _ := uuid.Parse("fcfa069c-28a9-48d6-b48d-befc8133f2b4")
		//
		//find, err := storage.Find(postID)
		//if err != nil {
		//	t.FailNow()
		//	return
		//}

		//created, err := time.Parse("2006-01-02 15:04:05", "2022-03-13 12:00:00")
		//if err != nil {
		//	t.FailNow()
		//	return
		//}
		//
		//post := sqlstorage.NewPost(
		//	"Title",
		//	created,
		//	"Description",
		//	userID,
		//)
		//
		//err = storage.Create(*post)
		//if err != nil {
		//	t.FailNow()
		//	return
		//}
		//
		//saved, err := storage.FindAll()
		//if err != nil {
		//	t.FailNow()
		//	return
		//}
		//require.Len(t, saved, 1)
		//require.Equal(t, *post, saved[0])
		//
		//post.Title = "Test title"
		//
		//err = storage.Update(*post)
		//if err != nil {
		//	t.FailNow()
		//	return
		//}
		//
		//saved, err = storage.FindAll()
		//if err != nil {
		//	t.FailNow()
		//	return
		//}
		//require.Len(t, saved, 1)
		//require.Equal(t, *post, saved[0])
		//
		//err = storage.Delete(post.ID)
		//if err != nil {
		//	t.FailNow()
		//	return
		//}
		//
		//saved, err = storage.FindAll()
		//if err != nil {
		//	t.FailNow()
		//	return
		//}
		//require.Len(t, saved, 0)

		//pagination, err := storage.Pagination(10, 1)
		//if err != nil {
		//	t.FailNow()
		//	return
		//}
		//fmt.Println(pagination)
		//require.Len(t, pagination, 10)

		//search, err := storage.Search("Post", "")
		//if err != nil {
		//	t.FailNow()
		//	return
		//}
		//
		//fmt.Println(search)

		//comment := sqlstorage.NewComment(
		//	"test4",
		//	"Content4",
		//	userID,
		//	postID,
		//)
		//
		//err = storage.CreateComment(*comment)
		//if err != nil {
		//	t.FailNow()
		//	return
		//}
		//require.Nil(t, nil, err)
		//
		//comment.Content = "Test change comment"
		//
		//err = storage.UpdateComment(*comment)
		//if err != nil {
		//	t.FailNow()
		//	return
		//}
		//require.Nil(t, nil, err)
		//
		//err = storage.DeleteComment(comment.ID)
		//if err != nil {
		//	t.FailNow()
		//	return
		//}
		//require.Nil(t, nil, err)

		err = tx.Rollback(ctx)
		if err != nil {
			t.Fatal("Failed to rollback tx", err)
		}
	})
}
