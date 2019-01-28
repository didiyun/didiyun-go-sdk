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

func TestListSg(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ListSgRequest{
		Header: &base.Header{RegionId: "gz"},
		Start:  0,
		Limit:  10,
		Condition: &ListSgCondition{
			//SgUuids:    []string{},
			VpcUuid: vpcUuid,
			//Dc2Uuid:    "953777262c9e5bd48d1a5379ca220811",
			//Dc2Exclude: true,
		},
	}
	out, err := iaasCli.ListSg(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestGetSgTotalCnt(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &GetSgTotalCntRequest{
		Header: &base.Header{RegionId: "gz"},
		//SgUuids:    []string{},
		VpcUuid: vpcUuid,
		//Dc2Uuid:    "953777262c9e5bd48d1a5379ca220811",
		//Dc2Exclude: false,
	}
	out, err := iaasCli.GetSgTotalCnt(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestCreateSg(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &CreateSgRequest{
		Header:  &base.Header{RegionId: "gz"},
		Name:    "test-sdk-create-sg-2",
		VpcUuid: vpcUuid,
		SgRule: []*CreateSgRuleInput{{
			Type:        "Ingress",
			Protocol:    "TCP",
			StartPort:   23,
			EndPort:     45,
			AllowedCidr: "10.0.0.0/16",
		}, {
			Type:        "Egress",
			Protocol:    "UDP",
			StartPort:   67,
			EndPort:     89,
			AllowedCidr: "10.0.0.0/16",
		}},
	}
	out, err := iaasCli.CreateSg(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, "gz", getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	for _, job := range result {
		sgUuid = job.ResourceUuid
	}
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestDeleteSg(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &DeleteSgRequest{
		Header: &base.Header{RegionId: "gz"},
		Sg: []*DeleteSgRequest_Input{{
			SgUuid: sgUuid,
		}},
	}
	out, err := iaasCli.DeleteSg(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, "gz", getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
}

func TestChangeSgName(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ChangeSgNameRequest{
		Header: &base.Header{RegionId: "gz"},
		Sg: []*ChangeSgNameRequest_Input{{
			SgUuid: sgUuid,
			Name:   "test-change-sg-name",
		}},
	}
	out, err := iaasCli.ChangeSgName(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, "gz", getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestAttachDc2ToSg(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &AttachDc2ToSgRequest{
		Header: &base.Header{RegionId: "gz"},
		Dc2: []*AttachDc2ToSgRequest_Dc2{{
			Dc2Uuid: dc2Uuid,
		}},
		Sg: []*AttachDc2ToSgRequest_Sg{{
			SgUuid: sgUuid,
		}},
	}
	out, err := iaasCli.AttachDc2ToSg(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, "gz", getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
}

func TestDetachDc2FromSg(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &DetachDc2FromSgRequest{
		Header: &base.Header{RegionId: "gz"},
		Dc2: []*DetachDc2FromSgRequest_Dc2{{
			Dc2Uuid: dc2Uuid,
		}},
		Sg: []*DetachDc2FromSgRequest_Sg{{
			SgUuid: sgUuid,
		}},
	}
	out, err := iaasCli.DetachDc2FromSg(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, "gz", getJobUuids(out.Data)...) //轮询查询进度
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
}

func TestListSgRule(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ListSgRuleRequest{
		Header: &base.Header{RegionId: "gz"},
		Start:  0,
		Limit:  10,
		Condition: &ListSgRuleCondition{
			SgUuid:  sgUuid,
			Dc2Uuid: "",
			//Type:    "Ingress",
		},
	}
	out, err := iaasCli.ListSgRule(ctx, in)
	panicErr(err)
	panicError(out.Error)
	for _, sgRule := range out.Data {
		sgRuleUuids = append(sgRuleUuids, sgRule.SgRuleUuid)
	}
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestGetSgRuleTotalCnt(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &GetSgRuleTotalCntRequest{
		Header:  &base.Header{RegionId: "gz", ZoneId: "gz01"},
		SgUuid:  sgUuid,
		Dc2Uuid: "",
		Type:    "Egress",
	}
	out, err := iaasCli.GetSgRuleTotalCnt(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestCreateSgRule(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &CreateSgRuleRequest{
		Header: &base.Header{RegionId: "gz"},
		SgUuid: sgUuid,
		SgRule: []*CreateSgRuleInput{{
			Type:        "Ingress",
			Protocol:    "TCP",
			StartPort:   145,
			EndPort:     267,
			AllowedCidr: "10.12.13.0/24",
		}},
	}
	out, err := iaasCli.CreateSgRule(ctx, in)
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

func TestDeleteSgRule(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &DeleteSgRuleRequest{
		Header: &base.Header{RegionId: "gz"},
	}
	for _, sgRuleUuid := range sgRuleUuids {
		in.SgRule = append(in.SgRule, &DeleteSgRuleRequest_Input{
			SgRuleUuid: sgRuleUuid,
		})
	}
	out, err := iaasCli.DeleteSgRule(ctx, in)
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
