package utils

import (
	"hash/crc32"
	"strconv"
	"time"
)

// ToString is any type trans to string
func ToString(v interface{}) string {
	switch f := v.(type) {
	case bool:
		if f {
			return "true"
		}
		return "false"
	case float32:
		return strconv.FormatFloat(float64(f), 'E', -1, 32)
	case float64:
		return strconv.FormatFloat(f, 'E', -1, 64)
	case int:
		return strconv.Itoa(f)
	case int8:
		return strconv.FormatInt(int64(f), 10)
	case int16:
		return strconv.FormatInt(int64(f), 10)
	case int32:
		return strconv.FormatInt(int64(f), 10)
	case int64:
		return strconv.FormatInt(f, 10)
	case uint:
		return strconv.FormatUint(uint64(f), 10)
	case uint8:
		return strconv.FormatUint(uint64(f), 10)
	case uint16:
		return strconv.FormatUint(uint64(f), 10)
	case uint32:
		return strconv.FormatUint(uint64(f), 10)
	case uint64:
		return strconv.FormatUint(f, 10)
	case time.Time:
		return f.Format("2006-01-02 15:04:05")
	case string:
		return f
	default:
		return ""
	}
}

// GetHashcode get string hash code
func GetHashcode(str string) int {
	v := int(int32(crc32.ChecksumIEEE([]byte(str))))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	return 0
}
