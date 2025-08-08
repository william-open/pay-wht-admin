package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
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
