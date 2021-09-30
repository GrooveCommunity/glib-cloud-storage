package gcp

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"

	"google.golang.org/api/iterator"

	"log"

	"cloud.google.com/go/storage"
)

func GetObjects(bucketName string) [][]byte {
	var dataObjects [][]byte

	ctx := context.Background()

	client, err := storage.NewClient(ctx)

	defer client.Close()

	if err != nil {
		panic(err.Error())
	}

	bucket := client.Bucket(bucketName)
	it := bucket.Objects(ctx, nil)

	for {

		attr, err := it.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			panic(err)
		}

		dataObjects = append(dataObjects, getObject(attr.Name, bucket, ctx))
	}

	return dataObjects
}

func GetObject(bucketName, objectName string) []byte {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	defer client.Close()

	if err != nil {
		panic(err)
	}

	bucket := client.Bucket(bucketName)

	return getObject(objectName, bucket, ctx)
}

func getObject(objectName string, bucket *storage.BucketHandle, ctx context.Context) []byte {
	objBucket := bucket.Object(objectName)
	reader, errorReader := objBucket.NewReader(ctx)

	if errorReader != nil {
		log.Fatal("Erro na criação de reader, ", errorReader.Error())
		panic(errorReader)
	}

	defer reader.Close()

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	writer.ReadFrom(reader)

	return b.Bytes()
}

func WriteObject(i interface{}, bucketName, objectName string) {
	ctx := context.Background()

	log.Print("Criando client")

	client, err := storage.NewClient(ctx)

	defer client.Close()

	if err != nil {

		log.Printf("Erro criando client %v", err)
		panic(err.Error())
	}

	bkt := client.Bucket(bucketName)
	obj := bkt.Object(objectName)
	log.Printf("Enviando arquivo %s para bucket %s", objectName, bucketName)
	w := obj.NewWriter(ctx)

	bI, err := json.Marshal(i)

	if err != nil {
		log.Println("Erro na conversão do struct: " + err.Error())

		panic(err)
	}

	_, err = w.Write(bI)

	if err != nil {
		log.Println("Erro gravando o arquivo: " + err.Error())

		panic(err)
	}

	err = w.Close()

	if err != nil {
		log.Println("Finalizando a gravação do arquivo: " + err.Error())

		panic(err)
	}
	log.Print("Arquivo enviado")
}
