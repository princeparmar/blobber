package allocation

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"0chain.net/core/common"
	// mocksAllocation "0chain.net/blobbercore/allocation/mocks"
)

func TestAllocationWantRead(t *testing.T) {
	type fields struct {
		blobberID string
		readPrice int64
	}
	type args struct {
		blobberID string
		numBlocks int64
	}
	tests := []struct {
		name string
		fields fields
		args args
		want int64
	}{
		{
			name: "WantRead / Basic",
			fields: fields{blobberID: "0", readPrice: 1},
			args: args{blobberID: "0", numBlocks: 16384},
			want: 1,
		},
		{
			name: "WantRead / Mismatched Ids",
			fields: fields{blobberID: "1", readPrice: 1},
			args: args{blobberID: "2", numBlocks: 16384},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allocation := Allocation{}
			allocation.Terms = append(allocation.Terms, &Terms{
				BlobberID: tt.fields.blobberID,
				ReadPrice: tt.fields.readPrice,
			})

			t.Log("Terms", *allocation.Terms[0], tt.args)

			got := allocation.WantRead(
				tt.args.blobberID, tt.args.numBlocks);
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAllocationWantWrite(t *testing.T) {
	type args struct {
		blobberID string
		size int64
		wmt common.Timestamp
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Table name",
			args: args{blobberID: "0",size: 0, wmt: common.Timestamp(0)}, 
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allocation := Allocation{}

			got := allocation.WantWrite(
				tt.args.blobberID, tt.args.size, tt.args.wmt);
			assert.Equal(t, tt.want, got)
		})
	}
}
