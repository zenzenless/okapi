package api

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/influxdata/influxdb/pkg/testing/assert"
	"github.com/sdcoffey/big"
)

func TestToTimeSeries(t *testing.T) {
	data := `{
		"code": "0",
		"msg": "",
		"data": [
		  [
			"1638288000000",
			"49990",
			"60000",
			"27777.7",
			"54725.4",
			"2283737240.5569905",
			"97096270580691.14781962"
		  ],
		  [
			"1638201600000",
			"20000.1",
			"57500",
			"20000",
			"49990",
			"720276208.65373569",
			"23478595600041.51382616"
		  ]
		]
	  }`
	res := &KlineResOKX{}
	err := json.Unmarshal([]byte(data), res)
	fmt.Println(res)
	assert.NoError(t, err)
	bar := BarMap[BAR_1d]
	series, err := res.ToTimeSeries(BAR_1d)
	assert.NoError(t, err)
	fmt.Printf("%v\n", *series)
	assert.Equal(t, len(series.Candles), 2)

	assert.Equal(t, series.LastCandle().Period.Start, time.UnixMilli(1638288000000))
	assert.Equal(t, series.LastCandle().Period.End, time.UnixMilli(1638288000000).Add(bar))

	assert.Equal(t, series.LastCandle().OpenPrice, big.NewFromString("49990"))
	assert.Equal(t, series.LastCandle().ClosePrice, big.NewFromString("54725.4"))
}

func TestGeneratedOrderField(t *testing.T) {
	str := `inst_type        
	inst_id          
	ccy             
	ord_id           
	cl_ord_id         
	tag             
	px              
	sz              
	ord_type         
	side            
	pos_side         
	td_mode          
	acc_fill_sz       
	fill_px          
	trade_id         
	fill_sz          
	fill_time        
	source          
	state           
	avg_px           
	lever           
	tp_trigger_px     
	tp_trigger_px_type 
	tp_ord_px         
	sl_trigger_px     
	sl_trigger_px_type 
	sl_ord_px         
	fee_ccy          
	fee             
	rebate_ccy       
	rebate          
	tgt_ccy          
	pnl             
	category        
	u_time           
	c_time`
	fileds := strings.Split(str, "\n")
	for _, f := range fileds {
		
		fmt.Printf("field.String(\"%s\"),\n", strings.TrimSpace(f))
	}
}
