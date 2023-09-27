package main

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
	_ "github.com/vektra/mockery/v2/cmd"
)

//go:generate go run github.com/vektra/mockery/v2 --name blockBlobClient --filename block_blob_client_mock.go --replace-type github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/generated.BlockBlobClientPutBlobFromURLResponse=github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob.UploadBlobFromURLResponse --replace-type github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/generated.BlockBlobClientCommitBlockListResponse=github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob.CommitBlockListResponse --replace-type github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/generated.BlockBlobClientStageBlockFromURLResponse=github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob.StageBlockFromURLResponse --replace-type github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/generated.BlockBlobClientGetBlockListResponse=github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob.GetBlockListResponse --replace-type github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/generated.BlobClientDeleteResponse=github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob.DeleteResponse --replace-type github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/generated.BlockListType=github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob.BlockListType
type blockBlobClient interface {
	CommitBlockList(ctx context.Context, base64BlockIDs []string, options *blockblob.CommitBlockListOptions) (blockblob.CommitBlockListResponse, error)
	Delete(ctx context.Context, o *blob.DeleteOptions) (blob.DeleteResponse, error)
	GetBlockList(ctx context.Context, listType blockblob.BlockListType, options *blockblob.GetBlockListOptions) (blockblob.GetBlockListResponse, error)
	StageBlockFromURL(ctx context.Context, base64BlockID string, sourceURL string, options *blockblob.StageBlockFromURLOptions) (blockblob.StageBlockFromURLResponse, error)
	UploadBlobFromURL(ctx context.Context, copySource string, options *blockblob.UploadBlobFromURLOptions) (blockblob.UploadBlobFromURLResponse, error)
}

func main() {

}
