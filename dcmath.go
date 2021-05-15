package dcmath

import (
	"math"
	"reflect"
	"strconv"

	"github.com/shopspring/decimal"
)

// AddFloat decimal类型加法
// return d1 + d2
func AddFloat(d1, d2 float64) float64 {
	decimalD1 := decimal.NewFromFloat(d1)
	decimalD2 := decimal.NewFromFloat(d2)
	decimalResult := decimalD1.Add(decimalD2)
	float64Result, _ := decimalResult.Float64()
	return float64Result
}

// NewStringDecimal 生成StringsDecimal
// return d1 + d2
func NewStringDecimal(d string) (decimal.Decimal, error) {
	return decimal.NewFromString(d)
}

// newStringsDecimal 生成两个SringsDecimal
// return d1 + d2
func newStringsDecimal(d1, d2 string) (decimal.Decimal, decimal.Decimal, error) {
	var decimalD1, decimalD2 decimal.Decimal

	decimalD1, err := decimal.NewFromString(d1)
	if err != nil {
		return decimalD1, decimalD1, err
	}

	decimalD2, err = decimal.NewFromString(d2)
	if err != err {
		return decimalD1, decimalD2, err
	}

	return decimalD1, decimalD2, nil
}

// AddStingPrecision decimal类型加法,带精度
// return d1 + d2
func AddStingPrecision(d1, d2 string, precision int32) (string, error) {
	decimalD1, decimalD2, err := newStringsDecimal(d1, d2)
	if err != nil {
		return "", err
	}

	return decimalD1.Add(decimalD2).Truncate(precision).String(), nil
}

// AddFloatPrecision decimal类型加法,带精度
// return d1 + d2
func AddFloatPrecision(d1, d2 float64, precision int) float64 {
	decimalD1 := decimal.NewFromFloat(d1)
	decimalD2 := decimal.NewFromFloat(d2)
	decimalResult := decimalD1.Add(decimalD2)
	float64Result, _ := decimalResult.Float64()
	return Round64(float64Result, precision)
}

// SubtractFloat decimal类型减法
// return d1 - d2
func SubtractFloat(d1, d2 float64) float64 {
	decimalD1 := decimal.NewFromFloat(d1)
	decimalD2 := decimal.NewFromFloat(d2)
	decimalResult := decimalD1.Sub(decimalD2)
	float64Result, _ := decimalResult.Float64()
	return float64Result
}

// SubStringPrecision decimal类型减法,带精度
// return d1 - d2
func SubStringPrecision(d1, d2 string, precision int32) (string, error) {
	decimalD1, decimalD2, err := newStringsDecimal(d1, d2)
	if err != nil {
		return "", err
	}

	return decimalD1.Sub(decimalD2).Truncate(precision).String(), nil
}

// SubtractFloatPrecision decimal类型减法,带精度
// return d1 - d2
func SubtractFloatPrecision(d1, d2 float64, precision int) float64 {
	decimalD1 := decimal.NewFromFloat(d1)
	decimalD2 := decimal.NewFromFloat(d2)
	decimalResult := decimalD1.Sub(decimalD2)
	float64Result, _ := decimalResult.Float64()
	return Round64(float64Result, precision)
}

// MultiplyStringPrecision decimal类型乘法
// return d1 * d2
func MultiplyStringPrecision(d1, d2 string, precision int32) (string, error) {
	decimalD1, decimalD2, err := newStringsDecimal(d1, d2)
	if err != nil {
		return "", err
	}

	return decimalD1.Mul(decimalD2).Truncate(precision).String(), nil
}

// MultiplyFloat decimal类型乘法
// return d1 * d2
func MultiplyFloat(d1, d2 float64) float64 {
	decimalD1 := decimal.NewFromFloat(d1)
	decimalD2 := decimal.NewFromFloat(d2)
	decimalResult := decimalD1.Mul(decimalD2)
	float64Result, _ := decimalResult.Float64()
	return float64Result
}

// MultiplyFloatPrecision decimal类型乘法,带精度
// return d1 * d2
func MultiplyFloatPrecision(d1, d2 float64, precision int) float64 {
	decimalD1 := decimal.NewFromFloat(d1)
	decimalD2 := decimal.NewFromFloat(d2)
	decimalResult := decimalD1.Mul(decimalD2)
	float64Result, _ := decimalResult.Float64()
	return Round64(float64Result, precision)
}

// DivideStringPrecision decimal类型除法
// return d1 / d2
func DivideStringPrecision(d1, d2 string, precision int32) (string, error) {
	decimalD1, decimalD2, err := newStringsDecimal(d1, d2)
	if err != nil {
		return "", err
	}
	decimal.DivisionPrecision = int(precision)
	return decimalD1.Div(decimalD2).Truncate(precision).String(), nil
}

// DivideFloat decimal类型除法
// return d1 / d2
func DivideFloat(d1, d2 float64) float64 {
	decimalD1 := decimal.NewFromFloat(d1)
	decimalD2 := decimal.NewFromFloat(d2)
	decimalResult := decimalD1.Div(decimalD2)
	float64Result, _ := decimalResult.Float64()
	return float64Result
}

// DivideFloatPrecision decimal类型除法,带精度
// return d1 / d2
func DivideFloatPrecision(d1, d2 float64, precision int) float64 {
	decimalD1 := decimal.NewFromFloat(d1)
	decimalD2 := decimal.NewFromFloat(d2)
	decimalResult := decimalD1.Div(decimalD2)
	float64Result, _ := decimalResult.Float64()
	return Round64(float64Result, precision)
}

// Round 由于直接取整不会四舍五入，故将该值加上 0.5 / 10的n次方（n为精确到小数点后几位）
// 再将结果乘10的n次方取整后除于10的n次方
// 浮点类型保留小数点后n位精度
func Round(f interface{}, n int) (r float64, err error) {
	pow10N := math.Pow10(n)
	switch f.(type) {
	case float32:
		v := reflect.ValueOf(f).Interface().(float32)
		r = math.Trunc((float64(v)+0.5/pow10N)*pow10N) / pow10N
	case float64:
		v := reflect.ValueOf(f).Interface().(float64)
		r = math.Trunc((v+0.5/pow10N)*pow10N) / pow10N
	}
	return r, err
}

// Round32 浮点类型保留小数点后n位精度
func Round32(f float32, n int) (r float64) {
	pow10N := math.Pow10(n)
	r = math.Trunc((float64(f)+0.5/pow10N)*pow10N) / pow10N
	return r
}

// Round64Carry 浮点类型保留小数点后n位精度
func Round64Carry(f float64, n int) (r float64) {
	pow10N := math.Pow10(n)
	r = math.Trunc((f+0.5/pow10N)*pow10N) / pow10N
	return r
}

// Round64 直接舍去n位后的小数,不进位
func Round64(f float64, n int) (r float64) {
	pow10N := math.Pow10(n)
	decimalD1 := decimal.NewFromFloat(f)
	decimalD2 := decimal.NewFromFloat(pow10N)
	decimalResult := decimalD1.Mul(decimalD2)
	res1, _ := decimalResult.Float64()
	res := math.Trunc(res1)
	r = res / pow10N
	return r
}

// f*pow10N会造成浮点数不精确
//func Round64(f float64, n int) (r float64) {
//	pow10N := math.Pow10(n)
//	r = math.Trunc(f*pow10N) / pow10N
//	return r
//}

const min = 0.00000001

func Compare(f1, f2 float64) int8 {
	if math.Dim(math.Abs(f1-f2), min) == 0 {
		return 0
	}

	num1 := decimal.NewFromFloat(f1)
	num2 := decimal.NewFromFloat(f2)
	max := decimal.Max(num1, num2)
	if max == num1 { //num1 > num2
		return 1
	}

	//num 1 < num2
	return -1
}

//CompareString
//     -1 if d1 <  d2
//      0 if d1 == d2
//     +1 if d1 >  d2
func CompareString(d1, d2 string) (int, error) {

	decimalD1, decimalD2, err := newStringsDecimal(d1, d2)
	if err != nil {
		return 0, err
	}

	return decimalD1.Cmp(decimalD2), nil
}

// CompareZeroString 和0比较
//     -1 if d1 <  0
//      0 if d1 == 0
//     +1 if d1 >  0
func CompareZeroString(d1 string) (int, error) {

	decimalD1, decimalD2, err := newStringsDecimal(d1, "0")
	if err != nil {
		return 0, err
	}

	return decimalD1.Cmp(decimalD2), nil
}

func AbsString(s string) (string, error) {
	d1, err := NewStringDecimal(s)
	if err != nil {
		return "", err
	}
	return d1.Abs().String(), nil
}

func Float64ToString(value float64, precision int) string {
	return strconv.FormatFloat(value, 'f', precision, 64)
}

func DoubleStringsDecimal(d1, d2 string) (decimal.Decimal, decimal.Decimal, error) {
	var decimalD1, decimalD2 decimal.Decimal

	decimalD1, err := decimal.NewFromString(d1)
	if err != nil {
		return decimalD1, decimalD1, err
	}

	decimalD2, err = decimal.NewFromString(d2)
	if err != err {
		return decimalD1, decimalD2, err
	}

	return decimalD1, decimalD2, nil
}

func PrecisionProcessing(d1 string, precision int32) (string, error) {
	decimalD1, err := NewStringDecimal(d1)
	if err != nil {
		return "", err
	}
	return decimalD1.Truncate(precision).String(), nil
}
