package gcp

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"

	"gopkg.in/yaml.v2"

	"log"

	"cloud.google.com/go/storage"
	"github.com/GrooveCommunity/glib-cloud-storage/entity"
	//	"google.golang.org/api/iterator"
)

func GetObject(bucketName, objectName string, dataObject *entity.DataObject) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)

	defer client.Close()

	if err != nil {
		panic(err.Error())
	}

	bucket := client.Bucket(bucketName)
	objBucket := bucket.Object(objectName)

	reader, errorReader := objBucket.NewReader(ctx)
	if errorReader != nil {
		log.Fatal("Erro na criação de reader, ", errorReader.Error())
		panic(err)
	}

	defer reader.Close()

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	writer.ReadFrom(reader)

	errUnmarsh := yaml.Unmarshal(b.Bytes(), &dataObject)

	if errUnmarsh != nil {
		log.Fatal("Erro no unmarshal\n", errUnmarsh.Error())
	}

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


	return object*/

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
