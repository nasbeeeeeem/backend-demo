package main

import (
	"backend-demo/pkg/interface/api/server"
)

func init() {

}

func main() {
	// クライアントの作成
	// ctx := context.Background()
	// client, err := secretmanager.NewClient(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // シークレットneon_dsnへのアクセス
	// resourceName := "projects/466355925872/secrets/neon_dsn/versions/1"
	// req := &secretmanagerpb.AccessSecretVersionRequest{
	// 	Name: resourceName,
	// }

	// // シークレット上にアクセスする
	// result, err := client.AccessSecretVersion(ctx, req)
	// if err != nil {
	// 	log.Fatalf("failed to access secret version: %v", err)
	// }
	// dsn := *(*string)(unsafe.Pointer(&result.Payload.Data))
	dsn := ""
	server.Server(dsn)
}
