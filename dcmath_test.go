package dcmath

import (
	"testing"
)

type FundsRecord struct {
	BalanceID   int64   `json:"balance_id"`   //全局唯一ID,用于防止kafka重复消费
	UserID      int64   `json:"user_id"`      //用户ID
	Balance     float64 `json:"balance"`      //金额
	BalanceLock float64 `json:"balance_lock"` //解锁金额
	Action      int8    `json:"action"`       //1、2、3 见Action_xx 定义
	Coin        string  `json:"coin"`         //处理的币种
}

func TestMultiplyFloatPrecision(t *testing.T) {
	//tradeVolume := 1.00000000
	tradeVolume := 1.0
	makerFee := -0.00030000
	vipMakerFeeNum := MultiplyFloatPrecision(tradeVolume, makerFee, 8)
	t.Logf("vipMakerFeeNum:%f", vipMakerFeeNum)
	t.Logf("vipMakerFeeNum:%v", vipMakerFeeNum)
}

func TestStruct(t *testing.T) {
	r := FundsRecord{}
	r.Balance = 0.000300
	t.Logf("TestAbs Balance:%f", r.Balance)
	t.Logf("TestAbs Balance:%v", r.Balance)
	t.Logf("TestAbs FundsRecord:%+v", r)
}

func TestAdd(t *testing.T) {
	res := AddFloatPrecision(0.0, 0.0, 8)
	t.Log("TestAdd:", res)
}

func TestRound(t *testing.T) {
	a := 1.23456789
	b := 2.345677845234
	c := MultiplyFloatPrecision(a, b, 8)
	t.Log(c)
}

func TestAddStingPrecision(t *testing.T) {
	var num int32 = 18

	a := "11.112333333444444456768"
	b := "0.2231231231233333333338"
	c, err := AddStingPrecision(a, b, num)
	if err != nil {
		t.Log(err)
	}

	t.Log(c)
}

func TestDivideString(t *testing.T) {
	var num int32 = 18
	a := "11.1231231231"
	b := "3"
	c, err := DivideStringPrecision(a, b, num)
	if err != nil {
		t.Log(err)
	}
	t.Log(c)
}

func TestMultiplyString(t *testing.T) {
	var num int32 = 18
	a := "0.1000000000000000"
	b := "3"

	c, err := MultiplyStringPrecision(a, b, num)
	if err != nil {
		t.Log(err)
	}
	t.Log(c)
}

func TestSubtractFloatPrecision(t *testing.T) {
	var num int32 = 18
	a := "11.12"
	b := "3.2"
	c, err := SubStringPrecision(a, b, num)
	if err != nil {
		t.Log(err)
	}
	t.Log(c)
}

func TestCompareString(t *testing.T) {
	//a, err := CompareString("2222.111", "2222.111")
	a, err := CompareString("0.000000000001", "0.000000000001")
	t.Log(a)
	t.Log(err)
}

func TestCompareZeroString(t *testing.T) {
	a, err := CompareZeroString("0.000000000001")
	if err != nil {
		t.Error(err)
	}
	t.Log("0.000000000001=0?", a)

	a, err = CompareZeroString("0.00000000000")
	if err != nil {
		t.Error(err)
	}
	t.Log("0.00000000000=0? ", a)

	a, err = CompareZeroString("0.0")
	if err != nil {
		t.Error(err)
	}
	t.Log("0.0=0?", a)

	a, err = CompareZeroString("0")
	if err != nil {
		t.Error(err)
	}
	t.Log("0=0?", a)
}

func TestPrecisionProcessing(t *testing.T) {
	a, err := PrecisionProcessing("0.11", 18)
	t.Log(a)
	t.Log(err)
}
