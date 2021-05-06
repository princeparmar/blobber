package allocation

import (
	"regexp"
	"context"
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	// "0chain.net/blobbercore/datastore"
	mocksDatastore "0chain.net/blobbercore/datastore/mocks"
)

func TestNewConnectionID(t *testing.T) {
	tests := []struct {
		name string
		want *regexp.Regexp
	}{
		{
			name: "NewConnectionID / Basic",
			want: regexp.MustCompile(`^[0-9]+$`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newConnectionID();
			assert.Equal(t, tt.want.MatchString(got), true)
		})
	}
}

func TestFindAllocations(t *testing.T) {
	type args struct {
		ctx context.Context
		offset int64
	}
	type want struct {
		allocations []*Allocation
		count int64
		err error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "FindAllocations / Basic",
			args: args{
				// ctx: context.Context{},
				offset: 0,
			},
			want: want{
				allocations: []*Allocation{},
				count: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
	
			mockStoreInterface := mocksDatastore.NewMockStoreInterface(mockCtrl)
			mockStoreInterface.EXPECT().
				CreateTransaction(tt.args.ctx).
				Return(tt.args.ctx).
				Times(1)

			gotAllocations, gotCount, _ := findAllocations(
				tt.args.ctx, tt.args.offset);
			assert.Equal(t, tt.want.allocations, gotAllocations)
			assert.Equal(t, tt.want.count, gotCount)
			// reserved / assert.Equal(t, tt.want.err, gotErr)
		})
	}
}
