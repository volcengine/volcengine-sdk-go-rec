volcengine data/predict api sdk, go version
<br>
```go
package main

import (
	"github.com/google/uuid"
	"github.com/volcengine/volcengine-sdk-go-rec/byteair"
	"github.com/volcengine/volcengine-sdk-go-rec/byteair/protocol"
	"github.com/volcengine/volcengine-sdk-go-rec/core"
	"github.com/volcengine/volcengine-sdk-go-rec/core/logs"
	"github.com/volcengine/volcengine-sdk-go-rec/core/metrics"
	"github.com/volcengine/volcengine-sdk-go-rec/core/option"
	"time"
)

var client byteair.Client

func init() {
	client, _ = (&byteair.ClientBuilder{}).
		// 必传,租户id.
		TenantId("xxx").
		// 必传,项目id.
		ApplicationId("xxx").
		// 必传,密钥AK,获取方式:【火山引擎控制台】->【个人信息】->【密钥管理】中获取.
		AK("xxx").
		// 必传,密钥SK,获取方式:【火山引擎控制台】->【个人信息】->【密钥管理】中获取.
		SK("xxx").
		// 必传,国内使用RegionAirCn.
		Region(core.RegionAirCn).
		Build()
	metrics.Init() // metrics上报init.建议开启,方便字节侧排查问题.
	logs.Level = logs.LevelInfo
}

func Write() {
	// 此处为示例数据,实际调用时需注意字段类型和格式
	dataList := []map[string]interface{}{
		{ // 第一条数据
			"id":            "1",
			"title":         "test_title1",
			"status":        0,
			"brand":         "volcengine",
			"pub_time":      1583641807,
			"current_price": 1.1,
		},
		{ // 第二条数据
			"id":            "2",
			"title":         "test_title2",
			"status":        1,
			"brand":         "volcengine",
			"pub_time":      1583641503,
			"current_price": 2.2,
		},
	}
	// topic为枚举值，请参考API文档
	topic := "item"
	// 同步离线天级数据，需要指定日期
	date, _ := time.Parse("2006-01-02", "2022-01-01")
	opts := []option.Option{
		// 测试数据/预同步阶段("pre_sync"),历史数据同步（"history_sync"）和增量天级数据上传（"incremental_sync_daily"）
		option.WithStage("pre_sync"),
		// 必传，要求每次请求的Request-Id不重复，若未传，sdk会默认为每个请求添加
		option.WithRequestId(uuid.NewString()),
		// 必传，数据产生日期，实际传输时需修改为实际日期
		option.WithDataDate(date),
	}
	rsp, err := client.WriteData(dataList, topic, opts...)
	if err != nil {
		logs.Error("[WriteData] occur error, msg:%s", err.Error())
		return
	}
	if !rsp.GetStatus().GetSuccess() {
		logs.Error("[WriteData] failure")
		return
	}
	logs.Info("[WriteData] success")
	return
}

func Done() {
	date, _ := time.Parse("2006-01-02", "2022-01-01")
	// 已经上传完成的数据日期，可在一次请求中传多个
	dateList := []time.Time{date}
	// 与离线天级数据传输的topic保持一致
	topic := "item"
	opts := []option.Option{
		// 测试数据/预同步阶段("pre_sync"),历史数据同步（"history_sync"）和增量天级数据上传（"incremental_sync_daily"）
		option.WithStage("pre_sync"),
		// 必传，要求每次请求的Request-Id不重复，若未传，sdk会默认为每个请求添加
		option.WithRequestId(uuid.NewString()),
	}
	rsp, err := client.Done(dateList, topic, opts...)
	if err != nil {
		logs.Error("[Done] occur error, msg: %s", err.Error())
		return
	}
	if !rsp.GetStatus().GetSuccess() {
		logs.Error("[Done] failure")
		return
	}
	logs.Info("[Done] success")
	return
}

func Predict() {
	req := &protocol.PredictRequest{
		User: &protocol.PredictUser{
			Uid: "uid1",
		},
		Context: &protocol.PredictContext{
			Spm:   "1$##$2$##$3$##$4",
			Extra: map[string]string{"extra_key": "extra_value"},
		},
		CandidateItems: []*protocol.PredictCandidateItem{&protocol.PredictCandidateItem{
			Id: "item_id1",
		}},
		ParentItem: &protocol.PredictParentItem{
			Id: "item_id2",
		},
	}
	opts := []option.Option{
		option.WithRequestId(uuid.NewString()),
		option.WithScene("default"),
		// 是否开启SPM路由.开启的话需要保证请求体里的SPM存在且绑定了栏位.
		// server会根据body里的SPM路由到选择的栏位.
		option.WithHeaders(map[string]string{
			"Enable-spm-route": "true",
		}),
	}
	rsp, err := client.Predict(req, opts...)
	if err != nil {
		logs.Error("[predict] occur error, msg: %s", err.Error())
		return
	}
	if !rsp.GetSuccess() {
		logs.Error("[predict] failure")
		return
	}
	rsp.GetRequestId()
	logs.Info("[predict] success")
}

func Callback() {
	req := &protocol.CallbackRequest{
		Uid:   "uid1",
		Scene: "default",
		Items: []*protocol.CallbackItem{
			{
				Id:    "item_id1",
				Pos:   "pos1",
				Extra: "{\"reason\":\"exposure\"}",
			},
			{
				Id:    "item_id1",
				Pos:   "pos1",
				Extra: "{\"reason\":\"filter\"}",
			},
		},
		PredictRequestId: "ds12ad61hnwo",
		Context: &protocol.CallbackContext{
			Spm:   "1$##$2$##$3$##$4",
			Extra: map[string]string{"extra_key": "extra_value"},
		},
		Extra: nil,
	}
	opts := []option.Option{
		option.WithRequestId(uuid.NewString()),
	}
	rsp, err := client.Callback(req, opts...)
	if err != nil {
		logs.Error("[callback] occur error, msg: %s", err.Error())
		return
	}
	if !rsp.GetSuccess() {
		logs.Error("[callback] failure")
		return
	}
	logs.Info("[callback] success")
}
```
