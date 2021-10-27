package handler

import (
	"context"
	"encoding/hex"
	blobbergrpc "github.com/0chain/blobber/code/go/0chain.net/blobbercore/blobbergrpc/proto"
	"github.com/0chain/blobber/code/go/0chain.net/blobbercore/reference"
	"github.com/0chain/blobber/code/go/0chain.net/core/common"
	"github.com/0chain/blobber/code/go/0chain.net/core/encryption"
	"google.golang.org/grpc/metadata"
	"net/http"
	"testing"
)

func TestBlobberGRPCService_MarketplaceShareInfo(t *testing.T) {
	bClient, tdController := setupHandlerTests(t)
	allocationTx := randString(32)

	pubKey, _, signScheme := GeneratePubPrivateKey(t)
	clientSignature, _ := signScheme.Sign(encryption.Hash(allocationTx))
	pubKeyBytes, _ := hex.DecodeString(pubKey)
	clientId := encryption.Hash(pubKeyBytes)

	err := tdController.ClearDatabase()
	if err != nil {
		t.Fatal(err)
	}

	path := "/"
	allocationId := "exampleId"
	pathHash := reference.GetReferenceLookup(allocationId, path)
	err = tdController.AddMarketplaceShareInfoTestData(allocationTx, pubKey, clientId, allocationId, path, pathHash)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name           string
		context        metadata.MD
		input          *blobbergrpc.MarketplaceShareInfoRequest
		expectingError bool
	}{
		{
			name: "Success Insert share ",
			context: metadata.New(map[string]string{
				common.ClientHeader:          clientId,
				common.ClientSignatureHeader: clientSignature,
				common.ClientKeyHeader:       pubKey,
			}),
			input: &blobbergrpc.MarketplaceShareInfoRequest{
				Allocation:          allocationTx,
				EncryptionPublicKey: pubKey,
				AuthTicket:          "",
				HttpMethod:          http.MethodPost,
				Path:                path,
				RefereeClientId:     "",
			},
			expectingError: false,
		},
		{
			name: "Success Revoke share ",
			context: metadata.New(map[string]string{
				common.ClientHeader:          clientId,
				common.ClientSignatureHeader: clientSignature,
				common.ClientKeyHeader:       pubKey,
			}),
			input: &blobbergrpc.MarketplaceShareInfoRequest{
				Allocation:          allocationTx,
				EncryptionPublicKey: pubKey,
				AuthTicket:          "",
				HttpMethod:          http.MethodDelete,
				Path:                path,
				RefereeClientId:     "abcdefgh",
			},
			expectingError: false,
		},
		{
			name: "Invalid Marketplace share method",
			context: metadata.New(map[string]string{
				common.ClientHeader:          clientId,
				common.ClientSignatureHeader: clientSignature,
				common.ClientKeyHeader:       pubKey,
			}),
			input: &blobbergrpc.MarketplaceShareInfoRequest{
				Allocation:          allocationTx,
				EncryptionPublicKey: pubKey,
				AuthTicket:          "",
				HttpMethod:          http.MethodGet,
				Path:                "/file.txt",
				RefereeClientId:     "abcdefgh",
			},
			expectingError: true,
		},
	}

	for _, tc := range testCases {
		ctx := context.Background()
		ctx = metadata.NewOutgoingContext(ctx, tc.context)
		_, err := bClient.MarketplaceShareInfo(ctx, tc.input)
		if err != nil {
			if !tc.expectingError {
				t.Fatal(err)
			}
			continue
		}

		if tc.expectingError {
			t.Fatal("expected error")
		}
	}
}