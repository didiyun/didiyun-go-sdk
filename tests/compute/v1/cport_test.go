package v1

import (
	"fmt"
	"github.com/didiyun/didiyun-go-sdk/base/v1"
	"github.com/didiyun/didiyun-go-sdk/compute/v1"
	"testing"
	"time"
	"context"
	c "github.com/didiyun/didiyun-go-sdk/tests/common"
)

func TestListCport(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &compute.ListCportRequest{
		Header: &base.Header{RegionId: "gz"},
		Offset: 0,
		Limit:10,
		Condition: &compute.ListCportCondition{},
	}
	out, err := iaasCli.ListCport(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestGetCportByUuid(t *testing.T){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &compute.GetCportByUuidRequest{
		Header: &base.Header{RegionId: "gz"},
		CportUuid: "2d23924e66cf43e6b8f062b7b99a87c5",
	}
	out, err := iaasCli.GetCportByUuid(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestCreateCport(t *testing.T){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &compute.CreateCportRequest{
		Header: &base.Header{RegionId: "gz"},
		Cport: &compute.CreateCportRequest_Input{
			VpcUuid:"9743948bd827420eab9ddf86998a624b",
			SubnetUuid:"4ed769cf66c547e0a9b2ec01451b1050",
			VmIp: "10.255.1.235",
			Ip: "10.255.1.3",
			PodNs: "abc",
			PodName : "abc",
			ContainerId :"abc",
		},
	}
	////////step 1 create
	CreateResp, err := iaasCli.CreateCport(ctx, in)
	panicErr(err)
	panicError(CreateResp.Error)
	fmt.Println(c.ToPrettyJsonString(CreateResp))
	var cportUuid string
	////////step 2 query job result
	success, result, err := iaasCli.WaitForJobResult(ctx, "gz", getJobUuids(CreateResp.Data)...) //轮询查询进度
	for _, job := range result {
		cportUuid = job.ResourceUuid
	}
	panicErr(err)
	fmt.Println("cportUuid: ", cportUuid)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
	////////step 3 query cport info
	query := &compute.GetCportByUuidRequest{
		Header: &base.Header{RegionId: "gz"},
		CportUuid: cportUuid,
	}
	cportInfo, err := iaasCli.GetCportByUuid(ctx, query)
	panicErr(err)
	panicError(cportInfo.Error)
	fmt.Println(c.ToPrettyJsonString(cportInfo))
}

func TestDeleteCport(t *testing.T){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &compute.DeleteCportRequest{
		Header: &base.Header{RegionId: "gz"},
		Cport: []*compute.DeleteCportRequest_Input{{
			CportUuid: "d5bf649dd04d4a079ea72f442da42f12",
		}},
	}
	out, err := iaasCli.DeleteCport(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
	success, result, err := iaasCli.WaitForJobResult(ctx, "gz", getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}