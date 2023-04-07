package public

import (
	"bytes"
	"encoding/json"
	"sort"
	"strings"

	"github.com/james-zhang-bing/okapi"
)

type param struct {
	key   string
	value string
}

func publicArgToID(arg map[string]string) (id string) {
	var params []*param
	for i, v := range arg {
		params = append(params, &param{
			key:   i,
			value: v,
		})
	}
	sort.Slice(params, func(i, j int) bool {
		return params[i].key < params[j].key
	})
	var str []string
	for _, v := range params {
		str = append(str, v.key, v.value)
	}
	id = strings.Join(str, "_")
	return id
}

type PublicChannel[ReqType any, DataType any] struct {
	Req        ReqType
	ReceivedCh chan *DataType
	Channel    string
}

func (p *PublicChannel[T, Ch]) GetArg() map[string]string {
	m := okapi.MI2MS(p.Req)
	m["channel"] = p.Channel
	return m
}

func (p *PublicChannel[T, Ch]) ReceiveData(b []byte) {
	d := json.NewDecoder(bytes.NewReader(b))
	var rec []Ch
	err := d.Decode(&rec)
	if err != nil {
		log.Errorf("decode %s failed %s", p.Channel, err)
		return
	}
	for _, v := range rec {
		select {
		case p.ReceivedCh <- &v:
		default:
			log.Warnf("the chan for receiving channel: %s  req:%s was full", p.Channel, p.Req)
		}
	}
}

func (p *PublicChannel[T, Ch]) Init(name string, req T) {
	p.Channel = name
	p.ReceivedCh = make(chan *Ch, 1024)
	p.Req = req

}
