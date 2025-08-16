package service

import (
	"fmt"
	"gorm.io/gorm"
	"ruoyi-go/app/dto"
	"ruoyi-go/framework/dal"
	"sort"
	"time"
)

type OrderReceiveService struct{}

// ---------------------- 分表规则 ----------------------
func getOrderTable(id uint, yearMonth string) string {
	tableIndex := id % 4 // 每月分4张表
	return fmt.Sprintf("p_order_%s_p%d", yearMonth, tableIndex)
}

func generateOrderTableNames(yearMonth string) []string {
	return []string{
		fmt.Sprintf("p_order_%s_p0", yearMonth),
		fmt.Sprintf("p_order_%s_p1", yearMonth),
		fmt.Sprintf("p_order_%s_p2", yearMonth),
		fmt.Sprintf("p_order_%s_p3", yearMonth),
	}
}

// ---------------------- 查询单个订单 ----------------------
func (s *OrderReceiveService) GetOrderReceiveById(id uint, yearMonth string) (*dto.OrderReceiveDetailResponse, error) {
	if yearMonth == "" {
		yearMonth = time.Now().Format("200601")
	}

	tableName := getOrderTable(id, yearMonth)
	var order dto.OrderReceiveDetailResponse

	err := dal.GormOrder.Table(tableName).Where("id = ?", id).First(&order).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("order not found")
		}
		return nil, err
	}
	return &order, nil
}

// ---------------------- 查询订单列表（聚合分页） ----------------------
func (s *OrderReceiveService) GetOrderReceiveList(param dto.OrderReceiveListRequest, isPaging bool) ([]dto.OrderReceiveListResponse, int) {
	var allOrders []dto.OrderReceiveListResponse
	var count int64

	yearMonth := param.YearMonth
	if yearMonth == "" {
		yearMonth = time.Now().Format("200601")
	}

	// 遍历所有分片表
	for _, table := range generateOrderTableNames(yearMonth) {
		var tableOrders []dto.OrderReceiveListResponse
		query := dal.GormOrder.Table(table)

		if param.Keyword != "" {
			query = query.Where("order_id LIKE ?", "%"+param.Keyword+"%")
		}
		if param.Status != "" {
			query = query.Where("status = ?", param.Status)
		}

		var tableCount int64
		query.Count(&tableCount)
		count += tableCount

		query.Find(&tableOrders)
		allOrders = append(allOrders, tableOrders...)
	}

	// 内存排序（按创建时间倒序）
	sort.Slice(allOrders, func(i, j int) bool {
		return allOrders[i].CreateTime.After(allOrders[j].CreateTime)
	})

	// 内存分页
	if isPaging {
		start := (param.PageNum - 1) * param.PageSize
		end := start + param.PageSize
		if start >= len(allOrders) {
			return []dto.OrderReceiveListResponse{}, int(count)
		}
		if end > len(allOrders) {
			end = len(allOrders)
		}
		return allOrders[start:end], int(count)
	}

	return allOrders, int(count)
}
