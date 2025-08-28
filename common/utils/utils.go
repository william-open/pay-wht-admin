package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 正则验证
// expr 正则表达式
// content 要验证的内容
func CheckRegex(expr, content string) bool {

	r, err := regexp.Compile(expr)
	if err != nil {
		return false
	}

	return r.MatchString(content)
}

// 比较工具
// 检查元素item是否存在于切片slice中
// 如果存在，返回true；如果不存在，返回false
func Contains[T comparable](slice []T, item T) bool {

	for _, value := range slice {
		if value == item {
			return true
		}
	}

	return false
}

// 过滤器
// 条件函数返回true，元素会被包含在结果中
func Filter[T interface{}](slice []T, condition func(T) bool) (result []T) {

	for _, value := range slice {
		if condition(value) {
			result = append(result, value)
		}
	}

	return result
}

// 脱敏工具
func Desensitize(content string, start, end int) string {

	if start < 0 || end < 0 || start > end {
		return content
	}

	var contentRune []rune

	for key, value := range content {
		if key >= start && key <= end {
			contentRune = append(contentRune, '*')
		} else {
			contentRune = append(contentRune, value)
		}
	}

	return string(contentRune)
}

// 字符串转为int数组
func StringToIntSlice(param, char string) ([]int, error) {

	intSlice := make([]int, 0)

	if param == "" {
		return intSlice, nil
	}

	stringSlice := strings.Split(param, char)

	for _, str := range stringSlice {

		num, err := strconv.Atoi(str)
		if err != nil {
			intSlice = append(intSlice, num)
			return nil, errors.New(str + "转换失败：" + err.Error())
		}

		intSlice = append(intSlice, num)
	}

	return intSlice, nil
}

// RandomString 返回随机字符串
func RandomString(length int) string {
	var allRandomStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	byteList := make([]byte, length)
	for i := 0; i < length; i++ {
		byteList[i] = allRandomStr[rand.Intn(62)]
	}
	return string(byteList)
}

// MakeMd5 制作MD5
func MakeMd5(data string) string {
	sum := md5.Sum([]byte(data))
	return hex.EncodeToString(sum[:])
}

// 通过MD5生成APIKEY值
func GenerateApiKey() string {
	b := make([]byte, 16) // 128-bit
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	hash := md5.Sum(b)
	return hex.EncodeToString(hash[:]) // 返回32位16进制字符串
}

// ✅ 生成类似 320266 的商户号
// 生成商户编号（随机8位数字）
func GenerateMerchantCode() string {
	seed := time.Now().UnixNano()
	return fmt.Sprintf("%08d", seed%1000000)
}

// GenerateMerchantOrderNo 生成商户订单号
func GenerateMerchantOrderNo(merchantID int64) string {
	timestamp := time.Now().Format("20060102150405") // 精确到秒
	randPart := rand.Intn(10000)                     // 4位随机数
	return fmt.Sprintf("M%d%s%04d", merchantID, timestamp, randPart)
}

// GetTimestampMs 生成 13 位时间戳（毫秒）
func GetTimestampMs() int64 {
	return time.Now().UnixNano() / 1e6
}

// GenerateProductName 通用英文名称生成器（适合国际化产品）
func GenerateProductName() string {
	adjectives := []string{"Smart", "Ultra", "Secure", "Fast", "Global", "Prime", "Easy", "Next"}
	nouns := []string{"Pay", "Link", "Flow", "Channel", "Bridge", "Token", "Gate", "Stream"}

	rand.Seed(time.Now().UnixNano())
	adj := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	suffix := rand.Intn(1000)

	return fmt.Sprintf("%s%s-%03d", adj, noun, suffix)
}

// GenerateSign 生成签名（用于请求或验证）
func GenerateSign(params map[string]string, secretKey string) string {
	keys := make([]string, 0, len(params))
	for k, v := range params {
		if k == "sign" || strings.TrimSpace(v) == "" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	for i, k := range keys {
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(params[k])
		if i < len(keys)-1 {
			sb.WriteString("&")
		}
	}
	sb.WriteString("&key=")
	sb.WriteString(secretKey)

	//log.Printf("签名query字符串:%v", sb.String())
	hash := md5.Sum([]byte(sb.String()))
	signStr := strings.ToUpper(hex.EncodeToString(hash[:]))
	//log.Printf("签名值: %v", signStr)
	return signStr
}

func StructToMapString(req interface{}) map[string]string {
	result := make(map[string]string)
	val := reflect.ValueOf(req)
	typ := reflect.TypeOf(req)

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}
		// 去除 json 标签中的 ",omitempty" 等
		jsonKey := jsonTag
		if commaIdx := len(jsonKey); commaIdx > 0 {
			if idx := indexOf(jsonKey, ','); idx != -1 {
				jsonKey = jsonKey[:idx]
			}
		}
		// 获取字段值
		value := val.Field(i).Interface()
		result[jsonKey] = toString(value)
	}
	return result
}

func StructToMapInterface(req interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(req)
	typ := reflect.TypeOf(req)

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}
		// 去除 json 标签中的 ",omitempty" 等
		jsonKey := jsonTag
		if commaIdx := len(jsonKey); commaIdx > 0 {
			if idx := indexOf(jsonKey, ','); idx != -1 {
				jsonKey = jsonKey[:idx]
			}
		}
		// 获取字段值
		value := val.Field(i).Interface()
		result[jsonKey] = toString(value)
	}
	return result
}

func toString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case fmt.Stringer:
		return val.String()
	default:
		return fmt.Sprintf("%v", val)
	}
}

func indexOf(s string, sep byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == sep {
			return i
		}
	}
	return -1
}

func HandleIpList(ipStr string) []string {
	var validIPs []string

	if ipStr == "" {
		return validIPs
	}

	// 拆分（无论是否包含逗号都统一处理）
	ips := strings.Split(ipStr, ",")
	for _, ip := range ips {
		ip = strings.TrimSpace(ip)
		if ip == "" {
			continue
		}
		if net.ParseIP(ip) != nil {
			validIPs = append(validIPs, ip)
		}
	}

	return validIPs
}

func DeduplicateIPs(ipList []string) []string {
	seen := make(map[string]struct{})
	var result []string

	for _, ip := range ipList {
		ip = strings.TrimSpace(ip)
		if ip == "" {
			continue
		}
		if _, exists := seen[ip]; !exists {
			seen[ip] = struct{}{}
			result = append(result, ip)
		}
	}

	return result
}

func SecurePaymentPassword(length int) string {
	var digitRunes = []rune("0123456789")

	// 用 sync.Pool 复用 rand 实例，避免频繁创建
	var randPool = sync.Pool{
		New: func() any {
			src := rand.NewSource(time.Now().UnixNano())
			return rand.New(src)
		},
	}
	r := randPool.Get().(*rand.Rand)
	defer randPool.Put(r)

	password := make([]rune, length)
	for i := range password {
		password[i] = digitRunes[r.Intn(len(digitRunes))]
	}
	return string(password)
}
