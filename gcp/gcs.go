package gcp

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"

	"log"

	"cloud.google.com/go/storage"
	//	"google.golang.org/api/iterator"
)

/*func GetObject(id, bucketName string) string {
	object := ""

	ctx := context.Background()

	client := getConnection(ctx)

	bucket := client.Bucket(bucketName)

	storageObjHdl := bucket.Object(id)

	if storageObjHdl != nil {
		return object
	}

	log.Println(storageObjHdl)
	/*
		//query := &storage.Query{Prefix: id}

		rdr, err := bucket.Object(id).NewReader(ctx)

		defer rdr.Close()

		if err != nil {
			panic(err)
		}

		b, err := io.ReadAll(rdr)

		if err != nil {
			panic(err)
		}

		log.Println(string(b))

		//var bucket

		//it := bkt.Objects(ctx, query)


	return object

}*/

func getConnection(ctx context.Context) *storage.Client {

	client, err := storage.NewClient(ctx)

	defer client.Close()

	if err != nil {
		panic(err.Error())
	}

	return client
}

func WriteObject(i interface{}, bucketName, objectName string) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)

	defer client.Close()

	if err != nil {
		panic(err.Error())
	}

	bkt := client.Bucket(bucketName)
	obj := bkt.Object(objectName)

	w := obj.NewWriter(ctx)

	bI, err := json.Marshal(i)

	if err != nil {
		log.Println("Erro na conversão do struct: " + err.Error())

		panic(err)
	}

	var b bytes.Buffer
	writ := gzip.NewWriter(&b)
	writ.Write(bI)
	writ.Close()

	w.Write(b.Bytes())
	w.Close()

}
