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

func (s *server) Start(ctx context.Context, in *LoadConfigReq) (out *ErrorResp, _ error) {
	var err error

	// only error use this
	defer func() {
		out = &ErrorResp{}
		if err != nil {
			out.Error = err.Error()
			instance = nil
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

	// only error use this
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

func (s *server) Exit(ctx context.Context, in *EmptyReq) (out *EmptyResp, _ error) {
	out = &EmptyResp{}

	// Connection closed
	os.Exit(0)
	return
}

func (s *server) Test(ctx context.Context, in *TestReq) (out *TestResp, _ error) {
	var err error
	out = &TestResp{Ms: 0}

	defer func() {
		if err != nil {
			out.Error = err.Error()
		}
	}()

	if in.Mode == TestMode_UrlTest {
		i := libcore.NewV2rayInstance()
		err = i.LoadConfig(in.Config.CoreConfig)
		if err != nil {
			return
		}
		err = i.Start()
		if err != nil {
			return
		}
		var t int32
		t, err = libcore.UrlTestV2ray(i, in.Inbound, in.Url, in.Timeout)
		out.Ms = t // sn: ms==0 是错误
		i.Close()
	} else {
		// TODO  impl
	}

	return
}
