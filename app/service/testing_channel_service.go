package service

import (
	"fmt"
	"log"
	"strconv"
	"wht-admin/app/dto"
	"wht-admin/common/curl"
	"wht-admin/common/utils"
	"wht-admin/config"
)

type TestingChannelService struct{}

// CreateCollectOrder 代收订单
func (s *TestingChannelService) CreateCollectOrder(param dto.TestCreatePayUpstreamProductRequest) (string, error) {

	// 测试下游商户信息
	var merchant dto.MerchantDetailResponse
	var testMerchantId = config.Data.TestDownMerchant.MerchantID
	mId, err := strconv.ParseInt(testMerchantId, 10, 64) // base 10, 返回 int64
	if err != nil {
		fmt.Println("转换失败:", err)
	} else {
		fmt.Println("转换成功:", mId)
	}
	merchant = (&MerchantService{}).GetMerchantByMerchantId(int(mId))

	var requestParam dto.TestCreateOrderReq

	requestParam.Amount = param.Amount
	requestParam.MerchantNo = merchant.AppId
	requestParam.Version = "V1"
	requestParam.TranFlow = utils.GenerateMerchantOrderNo(int64(merchant.MId))
	requestParam.TranDatetime = strconv.FormatInt(utils.GetTimestampMs(), 10)
	requestParam.PayType = param.PayType
	requestParam.NotifyUrl = config.Data.TestDownMerchant.NotifyUrl
	requestParam.ProductInfo = utils.GenerateProductName()
	pendingSignMap := utils.StructToMapString(requestParam)
	// 提取参数做签名（排除 Sign 字段）
	requestParam.Sign = utils.GenerateSign(pendingSignMap, merchant.ApiKey)
	// 请求下单服务
	var apiUrl = config.Data.TestDownMerchant.ApiUrl
	log.Printf("请求GO API服务，下单地址: %v,请求参数: %+v", apiUrl, utils.StructToMapInterface(requestParam))
	body, err := curl.DefaultClient().Send(&curl.RequestParam{
		Header: map[string]interface{}{
			"Content-Type": "application/json",
		},
		Url:    apiUrl,
		Method: "post",
		Json:   utils.StructToMapInterface(requestParam),
	})
	if err != nil {
		return body, err
	}
	log.Printf("请求GO API服务，返回数据: %+v", body)
	return body, nil

}
