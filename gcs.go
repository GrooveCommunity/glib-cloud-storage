package glibcloudstorage

import (
	"context"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func GetObject(name string) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)

	defer client.Close()

	if err != nil {
		panic(err.Error())
	}
}

/*func getConnection() {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)

	defer client.Close()

}

/*func writeObject(ctx Context, i interface{}) {
	ctx := context.Background()


}*/
