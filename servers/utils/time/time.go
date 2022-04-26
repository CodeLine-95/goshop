package time

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"
const timezone = "Asia/Shanghai"

type LocalTime struct {
	time.Time
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	second := t.Unix()
	return []byte(strconv.FormatInt(second, 10)), nil
}

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = LocalTime{Time: now}
	return
}
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t LocalTime) String() string {
	return time.Now().Format(timeFormat)
}

func (t LocalTime) local() time.Time {
	loc, _ := time.LoadLocation(timezone)
	return t.In(loc)
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// 将指定格式的日期，转成成 YYYY-MM-DD HH:ii:ss
//func (t LocalTime) FormatDateString(MyDateString any) string {
//	// 将日期字符串解析为Go的time对象第一个参数指定格式，第二个是日期字符串
//	myDate, err := time.Parse(timeFormat, MyDateString.(string))
//	if err != nil {
//		return ""
//	}
//	// Format使用与parse相同的格式样式，或者我们可以使用预先生成的常量
//	return myDate.Format(timeFormat)
//}
