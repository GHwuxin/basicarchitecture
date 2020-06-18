package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

// GetIntStr is from string get int str
func GetIntStr(str string) string {
	reg := regexp.MustCompile(`[1-9]\d*`)
	matchStr := reg.FindString(str)
	if matchStr == "" {
		return "0"
	}
	return matchStr
}

// GetInt is from string get int
func GetInt(str string) int {
	reg := regexp.MustCompile(`[1-9]\d*`)
	matchStr := reg.FindString(str)
	if matchStr == "" {
		return 0
	}
	num, err := strconv.Atoi(matchStr)
	if err != nil {
		return 0
	}
	return num
}

// Float32ToByte is trans float32 to []byte
func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

// ByteToFloat32 is trans []byte to float32
func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

// Float64ToByte is trans float64 to []byte
func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

// ByteToFloat64 is trans []byte to float64
func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

// ByteToInt32 byte to int32
func ByteToInt32(by []byte) int32 {
	var num int32
	bBUF := bytes.NewBuffer(by)
	binary.Read(bBUF, binary.LittleEndian, num)
	return num
}

// Int32ToByte int32 to byte
func Int32ToByte(num int32) []byte {
	bBUF := new(bytes.Buffer)
	binary.Write(bBUF, binary.LittleEndian, num)
	return bBUF.Bytes()
}

// ByteToInt64 byte to int64
func ByteToInt64(by []byte) int64 {
	var num int64
	bBUF := bytes.NewBuffer(by)
	binary.Read(bBUF, binary.LittleEndian, num)
	return num
}

// Int64ToByte int64 to byte
func Int64ToByte(num int64) []byte {
	bBUF := new(bytes.Buffer)
	binary.Write(bBUF, binary.LittleEndian, num)
	return bBUF.Bytes()
}

// Round is ceil round float
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

// Float64Round is keep n decimal
func Float64Round(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}
