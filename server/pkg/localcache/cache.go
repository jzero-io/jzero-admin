package localcache

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/errorx"
	"time"
)

type Cache struct {
	Vals        map[string][]byte
	errNotFound error
}

func (mc *Cache) Del(keys ...string) error {
	return mc.DelCtx(context.Background(), keys...)
}

func (mc *Cache) DelCtx(_ context.Context, keys ...string) error {
	var be errorx.BatchError

	for _, key := range keys {
		if _, ok := mc.Vals[key]; !ok {
			be.Add(mc.errNotFound)
		} else {
			delete(mc.Vals, key)
		}
	}

	return be.Err()
}

func (mc *Cache) Get(key string, val any) error {
	return mc.GetCtx(context.Background(), key, val)
}

func (mc *Cache) GetCtx(ctx context.Context, key string, val any) error {
	bs, ok := mc.Vals[key]
	if ok {
		return json.Unmarshal(bs, val)
	}

	return mc.errNotFound
}

func (mc *Cache) IsNotFound(err error) bool {
	return errors.Is(err, mc.errNotFound)
}

func (mc *Cache) Set(key string, val any) error {
	return mc.SetCtx(context.Background(), key, val)
}

func (mc *Cache) SetCtx(ctx context.Context, key string, val any) error {
	data, err := json.Marshal(val)
	if err != nil {
		return err
	}

	mc.Vals[key] = data
	return nil
}

func (mc *Cache) SetWithExpire(key string, val any, expire time.Duration) error {
	return mc.SetWithExpireCtx(context.Background(), key, val, expire)
}

func (mc *Cache) SetWithExpireCtx(ctx context.Context, key string, val any, expire time.Duration) error {
	return mc.Set(key, val)
}

func (mc *Cache) Take(val any, key string, query func(val any) error) error {
	return mc.TakeCtx(context.Background(), val, key, query)
}

func (mc *Cache) TakeCtx(ctx context.Context, val any, key string, query func(val any) error) error {
	if _, ok := mc.Vals[key]; ok {
		return mc.GetCtx(ctx, key, val)
	}

	if err := query(val); err != nil {
		return err
	}

	return mc.SetCtx(ctx, key, val)
}

func (mc *Cache) TakeWithExpire(val any, key string, query func(val any, expire time.Duration) error) error {
	return mc.TakeWithExpireCtx(context.Background(), val, key, query)
}

func (mc *Cache) TakeWithExpireCtx(ctx context.Context, val any, key string, query func(val any, expire time.Duration) error) error {
	return mc.Take(val, key, func(val any) error {
		return query(val, 0)
	})
}
