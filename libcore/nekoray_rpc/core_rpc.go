package nekoray_rpc

import (
	"context"
	"errors"
	"libcore"
	"libcore/device"
	"os"

	"github.com/sirupsen/logrus"
)

var instance *libcore.V2RayInstance

func setupCore() {
	device.IsNekoray = true
	libcore.InitCore("", "", "", nil, ".", "moe.nekoray.pc:bg", true, 50*1024)
}

func (s *server) Start(ctx context.Context, in *LoadConfigReq) (resp *ErrorResp, _ error) {
	var err error
	defer func() {
		resp = &ErrorResp{}
		if err != nil {
			resp.Error = err.Error()
		}
	}()

	logrus.Println("Start:", in.CoreConfig)

	if instance != nil {
		err = errors.New("Already started...")
		return
	}

	instance = libcore.NewV2rayInstance()

	err = instance.LoadConfig(in.CoreConfig)
	if err != nil {
		return
	}

	err = instance.Start()
	if err != nil {
		return
	}

	return
}

func (s *server) Stop(ctx context.Context, in *EmptyReq) (resp *ErrorResp, _ error) {
	var err error
	defer func() {
		resp = &ErrorResp{}
		if err != nil {
			resp.Error = err.Error()
		}
	}()

	if instance != nil {
		err = instance.Close()
		instance = nil
	}

	return
}

func (s *server) Exit(ctx context.Context, in *EmptyReq) (resp *ErrorResp, _ error) {
	var err error
	defer func() {
		resp = &ErrorResp{}
		if err != nil {
			resp.Error = err.Error()
		}
	}()

	// Connection closed
	os.Exit(0)

	return
}
