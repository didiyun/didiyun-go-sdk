package v1

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/didiyun/didiyun-go-sdk/base/v1"

	. "github.com/didiyun/didiyun-go-sdk/compute/v1"
	c "github.com/didiyun/didiyun-go-sdk/tests/common"
)

func TestListEip(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ListEipRequest{
		Header:    &base.Header{RegionId: "gz", ZoneId: "gz01"},
		Start:     0,
		Limit:     10,
		Simplify:  false,
		Condition: &ListEipCondition{
			//Eip: "172.22.52.125",
			//EipUuids: []string{"e0dd4100a9275897ad1a56422d697956"},
		},
	}
	out, err := iaasCli.ListEip(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestGetEipByUuid(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &GetEipByUuidRequest{
		Header:  &base.Header{RegionId: "gz"},
		EipUuid: eipUuid,
	}
	out, err := iaasCli.GetEipByUuid(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestGetEipTotalCnt(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &GetEipTotalCntRequest{
		Header: &base.Header{RegionId: "gz"},
		//EipUuids: []string{},
	}
	out, err := iaasCli.GetEipTotalCnt(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestCreateEip(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &CreateEipRequest{
		Header:         &base.Header{RegionId: "gz"},
		Bandwidth:      1,
		ChargeWithFlow: true,
		BindingUuid:    dc2Uuid,
	}
	out, err := iaasCli.CreateEip(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	for _, job := range result {
		eipUuids = append(eipUuids, job.ResourceUuid)
	}
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestDetachEipFromDc2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &DetachEipFromDc2Request{
		Header: &base.Header{RegionId: "gz"},
		Eip: []*DetachEipFromDc2Request_Input{{
			EipUuid: eipUuid,
		}},
	}
	out, err := iaasCli.DetachEipFromDc2(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestAttachEipToDc2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &AttachEipToDc2Request{
		Header: &base.Header{RegionId: "gz"},
		Eip: []*AttachEipToDc2Request_Input{{
			EipUuid:     eipUuid,
			BindingUuid: dc2Uuid,
		}},
	}
	out, err := iaasCli.AttachEipToDc2(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestChangeEipBandwidth(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ChangeEipBandwidthRequest{
		Header: &base.Header{RegionId: "gz"},
		Eip: []*ChangeEipBandwidthRequest_Input{{
			EipUuid:        eipUuid,
			Bandwidth:      2,
			ChargeWithFlow: false,
		}},
	}
	out, err := iaasCli.ChangeEipBandwidth(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestDeleteEip(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &DeleteEipRequest{
		Header: &base.Header{RegionId: "gz"},
	}
	for _, eipUuid := range eipUuids {
		in.Eip = append(in.Eip, &DeleteEipRequest_Input{
			EipUuid: eipUuid,
		})
	}
	out, err := iaasCli.DeleteEip(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}
