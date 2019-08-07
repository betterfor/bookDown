// Desc: 使用map的时候加锁
package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/betterfor/NovelDown/models"
	"strconv"
	"strings"
	"sync"
	"time"
)

type MapLock struct {
	Lock *sync.RWMutex
	Data map[string]interface{}
}

// 新建map
func NewMapWithData(data map[string]interface{}) *MapLock {
	return &MapLock{
		Lock: new(sync.RWMutex),
		Data: data,
	}
}

func NewMapLock() *MapLock {
	return &MapLock{
		Lock: new(sync.RWMutex),
		Data: make(map[string]interface{}),
	}
}

// 获取原始map数据
func (m *MapLock) GetData() map[string]interface{} {
	m.Lock.RLock()
	//logs.Debug("[GetData] Rlock.")
	defer func() {
		m.Lock.RUnlock()
		//logs.Debug("[GetData] RUnlock.")
	}()
	return m.Data
}

// 覆盖数据
func (m *MapLock) PutData(data map[string]interface{}) {
	m.Lock.Lock()
	//logs.Debug("[PutData] lock.")
	defer func() {
		m.Lock.Unlock()
		//logs.Debug("[PutData] Unlock.")
	}()
	m.Data = data
}

// 清空数据
func (m *MapLock) ClearData() {
	m.Lock.Lock()
	//logs.Debug("[ClearData] lock.")
	defer func() {
		m.Lock.Unlock()
		//logs.Debug("[ClearData] Unlock.")
	}()
	m.Data = make(map[string]interface{})
}

// 查询是否存在某个key
func (m *MapLock) IsExist(k string) bool {
	m.Lock.RLock()
	//logs.Debug("[IsExist] Rlock.")
	defer func() {
		m.Lock.RUnlock()
		//logs.Debug("[IsExist] RUnlock.")
	}()
	_, ok := m.Data[k]
	return ok
}

// 获取数据,只返回数据忽略不存在的情况
func (m *MapLock) GetV(key string) interface{} {
	v, _ := m.Get(key)
	return v
}

// 返回数据字符串形式
func (m *MapLock) GetVString(key string) string {
	v, ok := m.Get(key)
	if ok {
		switch v.(type) {
		case string:
			return v.(string)
			break
		case int:
			return strconv.Itoa(int(v.(int)))
		case int32:
			return strconv.Itoa(int(v.(int32)))
		case int64:
			return strconv.Itoa(int(v.(int64)))
		case []int:
			return IntsToStr(v.([]int))
		case []int32:
			return Int32sToStr(v.([]int32))
		case []int64:
			return Int64sToStr(v.([]int64))
		case []string:
			return StrsToStr(v.([]string))
		}
	}
	return ""
}

// 获取数据
func (m *MapLock) Get(key string) (interface{}, bool) {
	m.Lock.RLock()
	//logs.Debug("[Get] Rlock.")
	defer func() {
		m.Lock.RUnlock()
		//logs.Debug("[Get] RUnlock.")
	}()
	data := m.Data
	if _, ok := data[key]; ok {
		return data[key], true
	}
	return nil, false
}

// 删除某个key\value
func (m *MapLock) Delete(k string) {
	m.Lock.Lock()
	//logs.Debug("[Delete] lock.")
	defer func() {
		m.Lock.Unlock()
		//logs.Debug("[Delete] Unlock.")
	}()
	delete(m.Data, k)
}

// 添加数据
func (m *MapLock) Put(k string, v interface{}) {
	m.Lock.Lock()
	//logs.Debug("[Put] lock.")
	defer func() {
		m.Lock.Unlock()
		//logs.Debug("[Put] Unlock.")
	}()
	m.Data[k] = v
}

// 添加元素到 数组类型 的指定索引位置
// 只有在 eType == models.ArrayElemTypeCommon 时，e参数才有意义
func (m *MapLock) ArrayPut(k string, i int, eType models.MapArrayElemType, e interface{}) error {
	var data interface{}
	var size int

	m.Lock.Lock()
	//logs.Debug("[ArrayPut] Lock.")
	defer func() {
		m.Lock.Unlock()
		//logs.Debug("[ArrayPut] Unlock.")
	}()

	_, ok := m.Data[k]
	if !ok {
		switch e.(type) {
		case int:
			tmp := make([]int, i)
			if eType == models.ArrayElemTypeCommon {
				tmp[i] = e.(int)
			} else if eType == models.ArrayElemTypeInc {
				tmp[i]++
			} else if eType == models.ArrayElemTypeDec {
				tmp[i]--
			} else if eType == models.ArrayElemTypeAcc {
				tmp[i] += e.(int)
			} else if eType == models.ArrayElemTypeSubAcc {
				tmp[i] -= e.(int)
			} else if eType == models.ArrayElemTypeMax {
				if tmp[i] < e.(int) {
					tmp[i] = e.(int)
				}
			} else if eType == models.ArrayElemTypeMin {
				if tmp[i] > e.(int) {
					tmp[i] = e.(int)
				}
			}
			data = tmp
		case int64:
			tmp := make([]int64, i)
			if eType == models.ArrayElemTypeCommon {
				tmp[i] = e.(int64)
			} else if eType == models.ArrayElemTypeInc {
				tmp[i]++
			} else if eType == models.ArrayElemTypeDec {
				tmp[i]--
			} else if eType == models.ArrayElemTypeAcc {
				tmp[i] += e.(int64)
			} else if eType == models.ArrayElemTypeSubAcc {
				tmp[i] -= e.(int64)
			} else if eType == models.ArrayElemTypeMax {
				if tmp[i] < e.(int64) {
					tmp[i] = e.(int64)
				}
			} else if eType == models.ArrayElemTypeMin {
				if tmp[i] > e.(int64) {
					tmp[i] = e.(int64)
				}
			}
			data = tmp
		case string:
			tmp := make([]string, i)
			tmp[i] = e.(string)
			data = tmp
		default:
			return errors.New(fmt.Sprint("[ArrayPut] Value of ", e, "'s type not defined."))
		}
	} else {
		var values interface{}
		if _, ok := m.Data[k]; ok {
			values = m.Data[k]
		}

		switch values.(type) {
		case []int:
			size = len(values.([]int))
			if i >= size {
				goto OutOfIndex
			}
			if eType == models.ArrayElemTypeCommon {
				values.([]int)[i] = e.(int)
			} else if eType == models.ArrayElemTypeInc {
				values.([]int)[i]++
			} else if eType == models.ArrayElemTypeDec {
				values.([]int)[i]--
			} else if eType == models.ArrayElemTypeAcc {
				values.([]int)[i] += e.(int)
			} else if eType == models.ArrayElemTypeSubAcc {
				values.([]int)[i] -= e.(int)
			} else if eType == models.ArrayElemTypeMax {
				if values.([]int)[i] < e.(int) {
					values.([]int)[i] = e.(int)
				}
			} else if eType == models.ArrayElemTypeMin {
				if values.([]int)[i] > e.(int) {
					values.([]int)[i] = e.(int)
				}
			}
			data = values
		case []int64:
			size = len(values.([]int64))
			if i >= size {
				goto OutOfIndex
			}
			if eType == models.ArrayElemTypeCommon {
				values.([]int64)[i] = e.(int64)
			} else if eType == models.ArrayElemTypeInc {
				values.([]int64)[i]++
			} else if eType == models.ArrayElemTypeDec {
				values.([]int64)[i]--
			} else if eType == models.ArrayElemTypeAcc {
				values.([]int64)[i] += e.(int64)
			} else if eType == models.ArrayElemTypeSubAcc {
				values.([]int64)[i] -= e.(int64)
			} else if eType == models.ArrayElemTypeMax {
				if values.([]int64)[i] < e.(int64) {
					values.([]int64)[i] = e.(int64)
				}
			} else if eType == models.ArrayElemTypeMin {
				if values.([]int64)[i] > e.(int64) {
					values.([]int64)[i] = e.(int64)
				}
			}
			data = values
		case []string:
			size = len(values.([]string))
			if i >= size {
				goto OutOfIndex
			}
			values.([]string)[i] = e.(string)
			data = values
		default:
			return errors.New("[ArrayPut] value's type in key of '" + k + "' not defined.")
		}
	}

	m.Data[k] = data
	return nil
OutOfIndex:
	return errors.New(fmt.Sprint("[ArrayPut] Array's length is ", size, ", while put the value on index of", i))
}

// 添加元素到 数组类型 的value中
func (m *MapLock) ArrayAppend(k string, e interface{}) error {
	var data interface{}
	m.Lock.Lock()
	//logs.Debug("[ArrayAppend] Lock.")
	defer func() {
		m.Lock.Unlock()
		//logs.Debug("[ArrayAppend] Unlock.")
	}()

	_, ok := m.Data[k]
	if !ok {
		switch e.(type) {
		case int:
			tmp := make([]int, 0, 3)
			data = append(tmp, e.(int))
		case int64:
			tmp := make([]int64, 0, 3)
			data = append(tmp, e.(int64))
		case string:
			tmp := make([]string, 0, 3)
			data = append(tmp, e.(string))
		default:
			return errors.New(fmt.Sprint("[ArrayAppend] Value of ", e, "'s type not defined."))
		}
	} else {
		var values interface{}
		if _, ok := m.Data[k]; ok {
			values = m.Data[k]
		}

		switch values.(type) {
		case []int:
			data = append(values.([]int), e.(int))
		case []int64:
			data = append(values.([]int64), e.(int64))
		case []string:
			data = append(values.([]string), e.(string))
		default:
			return errors.New("[ArrayAppend] value's type in key of '" + k + "' not defined.")
		}
	}

	m.Data[k] = data
	return nil
}

// 返回map的key\value形式的字符串
func (m *MapLock) String() string {
	v, err := json.Marshal(m.GetData())
	if err == nil {
		return string(v)
	}
	return err.Error()
}

// 返回元素个数
func (m *MapLock) Count() int {
	m.Lock.RLock()
	//logs.Debug("[Count] RLock.")
	defer func() {
		m.Lock.RUnlock()
		//logs.Debug("[Count] RUnlock.")
	}()
	return len(m.Data)
}

// 避免频繁更新,加锁n秒后可操作
func WriteLock(key string, lock *MapLock, timeout int64) bool {
	if len(lock.GetData()) > 0 {
		v, err := lock.Get(key)
		if err {
			last := v.(int64)
			if time.Now().Unix()-last < timeout {
				return false
			}
		}
	}
	lock.Put(key, time.Now().Unix())
	return true
}

func StrToInt64s(str string) (int, []int64, error) {
	strs := strings.Split(str, " ")
	size := len(strs)
	ret := make([]int64, size)
	for i, v := range strs {
		value, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			logs.Error("[StrToInt64s] str of", v, "transfer to int64 err:", err)
			return 0, nil, err
		}
		ret[i] = value
	}
	return size, ret, nil
}

func StrToInts(str string) (int, []int, error) {
	strs := strings.Split(str, " ")
	size := len(strs)
	ret := make([]int, size)
	for i, v := range strs {
		value, err := strconv.Atoi(v)
		if err != nil {
			logs.Error("[StrToInts] str of", v, "transfer to int err:", err)
			return 0, nil, err
		}
		ret[i] = value
	}
	return size, ret, nil
}

// 返回数组内容，不包含首位的'['\']'符号
func IntsToStr(values []int) string {
	if len(values) == 0 {
		return ""
	}
	str := fmt.Sprint(values)
	str = str[1 : len(str)-1] // 删除首位的'['\']'
	return str
}

// 返回数组内容，不包含首位的'['\']'符号
func Int32sToStr(values []int32) string {
	if len(values) == 0 {
		return ""
	}
	str := fmt.Sprint(values)
	str = str[1 : len(str)-1] // 删除首位的'['\']'
	return str
}

// 返回数组内容，不包含首位的'['\']'符号
func Int64sToStr(values []int64) string {
	if len(values) == 0 {
		return ""
	}
	str := fmt.Sprint(values)
	str = str[1 : len(str)-1] // 删除首位的'['\']'
	return str
}

func StrsToStr(values []string) string {
	if len(values) == 0 {
		return ""
	}
	str := fmt.Sprint(values)
	str = str[1 : len(str)-1] // 删除首位的'['\']'
	return str
}

func StrToStrs(str string) []string {
	return strings.Split(str, " ")
}

func Bytes2Obj(input []byte, output interface{}) error {
	var err error
	if len(input) != 0 {
		return json.Unmarshal(input, &output)
	} else {
		err = errors.New("input data is empty.")
	}
	return err
}

// output: 必须传入地址，否则无法改变内容
func String2Obj(input string, output interface{}) error {
	return Bytes2Obj([]byte(input), output)
}
