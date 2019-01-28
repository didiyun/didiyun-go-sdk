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

func TestListSnap(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ListSnapshotRequest{
		Header:    &base.Header{RegionId: "gz", ZoneId: "gz01"},
		Start:     0,
		Limit:     10,
		Simplify:  false,
		Condition: &ListSnapshotCondition{
			//EbsUuid: "c1dbc495456956ca8c8fe7a270f94865",
			//SnapName: "test-sdk2",
		},
	}
	out, err := iaasCli.ListSnapshot(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestGetSnapTotalCnt(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &GetSnapshotTotalCntRequest{
		Header: &base.Header{RegionId: "gz"},
		//EbsUuid: "c1dbc495456956ca8c8fe7a270f94865",
		SnapName: "test",
	}
	out, err := iaasCli.GetSnapshotTotalCnt(ctx, in)
	panicErr(err)
	panicError(out.Error)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestCreateSnap(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &CreateSnapshotRequest{
		Header:   &base.Header{RegionId: "pre"},
		Dc2Uuid:  "cf0fe02457a455108e18c655060ed25f",
		SnapName: "test-create-snap",
	}
	out, err := iaasCli.CreateSnapshot(ctx, in)
	fmt.Println(c.ToPrettyJsonString(out))
	panicErr(err)
	panicError(out.Error)
	success, result, err := iaasCli.WaitForJobResult(ctx, "gz", getJobUuids(out.Data)...) //轮询查询进度
	for _, job := range result {
		snapUuid = job.ResourceUuid
	}
	panicErr(err)
	fmt.Println("Success: ", success, "result: ", result)
	if !success {
		panic("not success")
	}
}

func TestDeleteSnap(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &DeleteSnapshotRequest{
		Header: &base.Header{RegionId: "gz"},
		Snap: []*DeleteSnapshotRequest_Input{{
			SnapUuid: snapUuid,
		}},
	}
	out, err := iaasCli.DeleteSnapshot(ctx, in)
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

func TestRevertSnap(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &RevertSnapshotRequest{
		Header: &base.Header{RegionId: "gz"},
		Snap: []*RevertSnapshotRequest_Input{{
			SnapUuid: snapUuid,
		}},
	}
	out, err := iaasCli.RevertSnapshot(ctx, in)
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

func TestChangeSnapshotName(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ChangeSnapshotNameRequest{
		Header: &base.Header{RegionId: "gz"},
		Snap: []*ChangeSnapshotNameRequest_Input{{
			SnapUuid: snapUuid,
			Name:     "test-change-snap-name",
		}},
	}
	out, err := iaasCli.ChangeSnapshotName(ctx, in)
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
