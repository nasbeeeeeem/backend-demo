package main

import (
	"backend-demo/pkg/interface/api/server"
	"context"
	"log"
	"unsafe"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

func init() {

}

func main() {
	// クライアントの作成
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// シークレットneon_dsnへのアクセス
	resourceName := "projects/" + "811974289373" + "/secrets/" + "neon_dsn" + "/versions/latest"
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: resourceName,
	}

	// シークレット上にアクセスする
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Fatalf("failed to access secret version: %v", err)
	}
	dsn := *(*string)(unsafe.Pointer(&result.Payload.Data))
	// dsn := ""
	server.Server(dsn)
}
