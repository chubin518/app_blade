package database

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

var (
	_ json.Marshaler   = (LocalTime)(time.Now())
	_ json.Unmarshaler = (*LocalTime)(nil)
	_ sql.Scanner      = (*LocalTime)(nil)
	_ driver.Valuer    = (LocalTime)(time.Now())
)

const (
	timeFormat = "2006-01-02 15:04:05"
	timeZone   = "Asia/Shanghai"
)

type LocalTime time.Time

// Value implements driver.Valuer.
func (lt LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(lt)
	// Check whether the given time is equal to the default zero time
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

// Scan implements sql.Scanner.
func (lt *LocalTime) Scan(value any) (err error) {
	if value, ok := value.(time.Time); ok {
		*lt = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to LocalTime", value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (lt *LocalTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*lt = LocalTime(now)
	return
}

// MarshalJSON implements json.Marshaler.
func (lt LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(lt).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (lt LocalTime) IsZero() bool {
	return time.Time(lt).IsZero()
}

func NowTime() LocalTime {
	return LocalTime(time.Now())
}

func ZeroTime() LocalTime {
	return LocalTime{}
}
