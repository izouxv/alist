package alist

import (
	"context"

	// _ "github.com/alist-org/alist/v3/drivers"
	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/alist-org/alist/v3/internal/op"
	// "github.com/alist-org/alist/v3/internal/op"
)

type Driver interface {
	driver.Driver
}
type Obj interface {
	model.Obj
}
type Info struct {
	driver.Info
}
type UpdateProgress func(percentage float64)

func GetDriver(config map[string]string) (Driver, error) {
	name := config["name"]
	driverFunc, err := op.GetDriver(name)
	if err != nil {
		return nil, err
	}
	driv := driverFunc()
	// add := driv.GetAddition()
	// err = util.Json2Obj(util.Obj2Json(config, true), add)
	if err != nil {
		return nil, err
	}
	err = driv.Init(context.Background())
	return driv, nil
}

// func GetDriverInfoMap() (x map[string]Info) {
// 	info := op.GetDriverInfoMap()
// 	x = make(map[string]Info)
// 	for k, v := range info {
// 		x[k] = Info{v}
// 	}
// 	return
// }
