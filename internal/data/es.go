package data

import (
	"bytes"
	"context"
	"encoding/json"
)

// ESInsertDoc: Index creates or updates a document in an index.
func (d *dataRepo) ESInsertDoc(ctx context.Context, indexName string, data interface{}) (err error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		return err
	}
	res, err := d.data.es.Index(indexName, &buf)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return err
}
