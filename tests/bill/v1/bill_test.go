package v1

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/didiyun/didiyun-go-sdk/base/v1"

	. "github.com/didiyun/didiyun-go-sdk/bill/v1"
	c "github.com/didiyun/didiyun-go-sdk/tests/common"
)

func panicError(err2 *base.Error, err error) {
	if err != nil {
		panic("Query Err: " + err.Error())
	}
	if err2 != nil && err2.Errno != 0 {
		panic("Errno != 0: " + strconv.Itoa(int(err2.Errno)) + "ErrMsg: " + err2.Errmsg + " requestId: " + err2.RequestId)
	}
}

func GetBillClient() BillClient {
	conn, err := c.DialTCP(c.ServerAddr)
	if err != nil {
		panic(err)
	}
	return NewBillClient(conn)
}

func TestCheckDc2Price(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &CheckDc2PriceRequest{
		Header: &base.Header{RegionId: "gz"},
		IsChange: true,
		Dc2Goods: &CheckDc2PriceInput{
			PayPeriod: 1,
			Dc2Model:  "dc2.e1.large4",
			Dc2Uuid:"92cd0ec57d3455ecb457b16096e9024c",
		},
	}
	out, err := GetBillClient().CheckDc2Price(ctx, in)
	panicError(out.Error, err)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestCheckEipPrice(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &CheckEipPriceRequest{
		Header: &base.Header{RegionId: "gz"},
		EipGoods: &CheckEipPriceInput{
			PayPeriod: 2,
			Bandwidth: 1,
		},
	}
	out, err := GetBillClient().CheckEipPrice(ctx, in)
	panicError(out.Error, err)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestCheckEbsPrice(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &CheckEbsPriceRequest{
		Header: &base.Header{RegionId: "gz"},
		EbsGoods: &CheckEbsPriceInput{
			PayPeriod: 2,
			Size:      30,
			DiskType:  "HDD",
		},
	}
	out, err := GetBillClient().CheckEbsPrice(ctx, in)
	panicError(out.Error, err)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestContinueList(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ContinueListRequest{
		Header:    &base.Header{RegionId: "gz", ZoneId: "gz01"},
		Start:     0,
		Limit:     10,
		Condition: &ContinueListCondition{
			//ResourceType:"dc2",
		},
	}
	out, err := GetBillClient().ContinueList(ctx, in)
	panicError(out.Error, err)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestContinueMonthly(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ContinueMonthlyRequest{
		Header:    &base.Header{RegionId: "gz"},
		PayPeriod: 2,
		Resource: []*ResourceItemInput{
			{
				ResourceUuid: "e0dd4100a9275897ad1a56422d697956",
				ResourceType: "eip",
			},
		},
	}
	out, err := GetBillClient().ContinueMonthly(ctx, in)
	panicError(out.Error, err)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestChangeToMonthly(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ChangeToMonthlyRequest{
		Header:    &base.Header{RegionId: "gz"},
		PayPeriod: 1,
		Resource: []*ResourceItemInput{
			{
				ResourceUuid: "86e85c57c6e85bbebf8d76f55166d31e",
				ResourceType: "dc2",
			},
		},
	}
	out, err := GetBillClient().ChangeToMonthly(ctx, in)
	panicError(out.Error, err)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestChangeExpireStrategy(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &ChangeExpireStrategyRequest{
		Header:       &base.Header{RegionId: "gz"},
		AutoRenewCnt: 5,
		AutoSwitch:   false,
		Resource: []*ResourceItemInput{{
			ResourceUuid: "e0dd4100a9275897ad1a56422d697956",
			ResourceType: "eip",
		}},
	}
	out, err := GetBillClient().ChangeExpireStrategy(ctx, in)
	panicError(out.Error, err)
	fmt.Println(c.ToPrettyJsonString(out))
}

func TestGetChargeInfoAndSpec(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1200)
	defer cancel()
	in := &GetChargeInfoAndSpecRequest{
		Header:       &base.Header{RegionId: "gz"},

		Resource: []*GetChargeInfoInput{{
			ResourceUuid: "c97ec813180656bbbc6e4474c76a8f2b",
		}},
	}
	out, err := GetBillClient().GetDc2ChargeInfoAndSpec(ctx, in)
	panicError(out.Error, err)
	fmt.Println(c.ToPrettyJsonString(out))
}