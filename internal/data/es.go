package data

import (
	"context"

	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	pelasticsearch "github.com/xiaohubai/go-grpc-layout/pkg/elasticsearch"
)

// InsertDoc: Index creates or updates a document in an index.
func (d *dataRepo) InsertDoc(ctx context.Context, indexName string, data []byte) (err error) {
	es := pelasticsearch.New(d.data.es)
	return es.InsertDoc(ctx, consts.ESIndexOperationRecord, data)
}
