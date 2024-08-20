package alist

import (
	"context"

	// _ "github.com/alist-org/alist/v3/drivers"
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/alist-org/alist/v3/internal/op"
)

func ClearCache(storage Driver, path string) {
	op.ClearCache(storage, path)
}
func Key(storage Driver, path string) string {
	return op.Key(storage, path)
}
func List(ctx context.Context, storage Driver, path string, args model.ListArgs) (o []Obj, e error) {
	obs, err := op.List(ctx, storage, path, args)
	if err != nil {
		return nil, err
	}
	for _, it := range obs {
		o = append(o, it)
	}
	return
}
func Get(ctx context.Context, storage Driver, path string) (Obj, error) {
	return op.Get(ctx, storage, path)
}
func GetUnwrap(ctx context.Context, storage Driver, path string) (Obj, error) {
	return op.GetUnwrap(ctx, storage, path)
}
func Link(ctx context.Context, storage Driver, path string, args model.LinkArgs) (*model.Link, Obj, error) {
	return op.Link(ctx, storage, path, args)
}
func Other(ctx context.Context, storage Driver, args model.FsOtherArgs) (interface{}, error) {
	return op.Other(ctx, storage, args)
}
func MakeDir(ctx context.Context, storage Driver, path string, lazyCache ...bool) error {
	return op.MakeDir(ctx, storage, path, lazyCache...)
}
func Move(ctx context.Context, storage Driver, srcPath, dstDirPath string, lazyCache ...bool) error {
	return op.Move(ctx, storage, srcPath, dstDirPath, lazyCache...)
}
func Rename(ctx context.Context, storage Driver, srcPath, dstName string, lazyCache ...bool) error {
	return op.Rename(ctx, storage, srcPath, dstName, lazyCache...)
}
func Copy(ctx context.Context, storage Driver, srcPath, dstDirPath string, lazyCache ...bool) error {
	return op.Copy(ctx, storage, srcPath, dstDirPath, lazyCache...)
}
func Remove(ctx context.Context, storage Driver, path string) error {
	return op.Remove(ctx, storage, path)
}
func Put(ctx context.Context, storage Driver, dstDirPath string, file model.FileStreamer, up UpdateProgress, lazyCache ...bool) error {
	return op.Put(ctx, storage, dstDirPath, file, func(percentage float64) {
		up(percentage)
	}, lazyCache...)
}
