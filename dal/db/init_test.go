package db

import (
	"context"
	"testing"
)

func Test_initCache(t *testing.T) {
	initCache()
	if err := Redis_client.Ping(context.Background()).Err(); err != nil {
		t.Error(err)
	}
}
