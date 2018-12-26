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

func TestListEbs(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ListEbsRequest{
		Header:    &base.Header{RegionId: "gz", ZoneId: "gz01"},
		Start:     0,
		Limit:     10,
		Simplify:  false,
		Condition: &ListEbsCondition{
			//Dc2Uuids: []string{"b8bf159d2ce258f98b27aedafcaadc91"},
		},
	}
	out, err := iaasCli.ListEbs(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestGetEbsByUuid(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &GetEbsByUuidRequest{
		Header:  &base.Header{RegionId: "gz"},
		EbsUuid: ebsUuid,
	}
	out, err := iaasCli.GetEbsByUuid(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestGetEbsTotalCnt(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &GetEbsTotalCntRequest{
		Header: &base.Header{RegionId: "gz"},
		//Dc2Uuids: []string{"b8bf159d2ce258f98b27aedafcaadc91"},
	}
	out, err := iaasCli.GetEbsTotalCnt(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestCreateEbs(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &CreateEbsRequest{
		Header:   &base.Header{RegionId: "gz", ZoneId: "gz01"},
		Count:    1,
		Name:     "sdk-test-ebs",
		Size:     23,
		DiskType: "SSD",
		//Dc2Uuid:  dc2Uuid,
	}
	out, err := iaasCli.CreateEbs(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, getJobUuids(out.Data)...) //轮询查询进度
	for _, job := range result {
		ebsUuids = append(ebsUuids, job.ResourceUuid)
	}
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestDetachEbsFromDc2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &DetachEbsRequest{
		Header: &base.Header{RegionId: "gz"},
		Ebs: []*DetachEbsRequest_Input{{
			EbsUuid: ebsUuid,
		}},
	}
	out, err := iaasCli.DetachEbs(ctx, in)
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

func TestAttachEbsToDc2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &AttachEbsRequest{
		Header: &base.Header{RegionId: "gz"},
		Ebs: []*AttachEbsRequest_Input{{
			EbsUuid: ebsUuid,
			Dc2Uuid: dc2Uuid,
		}},
	}
	out, err := iaasCli.AttachEbs(ctx, in)
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

func TestChangeEbsSize(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ChangeEbsSizeRequest{
		Header: &base.Header{RegionId: "gz"},
		Ebs: []*ChangeEbsSizeRequest_Input{{
			EbsUuid: ebsUuid,
			Size:    33,
		}},
	}
	out, err := iaasCli.ChangeEbsSize(ctx, in)
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

func TestChangeEbsName(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ChangeEbsNameRequest{
		Header: &base.Header{RegionId: "gz"},
		Ebs: []*ChangeEbsNameRequest_Input{{
			EbsUuid: ebsUuid,
			Name:    "TestChangeEbsName",
		}},
	}
	out, err := iaasCli.ChangeEbsName(ctx, in)
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

func TestDeleteEbs(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &DeleteEbsRequest{
		Header: &base.Header{RegionId: "gz"},
	}
	for _, ebsUuid := range ebsUuids {
		in.Ebs = append(in.Ebs, &DeleteEbsRequest_Input{EbsUuid: ebsUuid})
	}
	out, err := iaasCli.DeleteEbs(ctx, in)
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
