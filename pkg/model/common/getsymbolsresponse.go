package common

import "github.com/shopspring/decimal"

type GetSymbolsResponse struct {
	Status string   `json:"status"`
	Data   []Symbol `json:"data"`
	Ts     string   `json:"ts"`
	Full   int32    `json:"full"`
}

type SymbolPartition struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Weight int    `json:"weight"`
}

type Symbol struct {
	Tags  string            `json:"tags"`
	State string            `json:"state"`
	Wr    string            `json:"wr"`
	Sc    string            `json:"sc"`
	P     []SymbolPartition `json:"p"`
	Bcdn  string            `json:"bcdn"`
	Qcdn  string            `json:"qcdn"`
	Elr   interface{}       `json:"elr"`
	Tpp   int               `json:"tpp"`
	Tap   int               `json:"tap"`
	Fp    int               `json:"fp"`
	Smlr  interface{}       `json:"smlr"`
	Flr   interface{}       `json:"flr"`
	Whe   bool              `json:"whe"`
	Cd    bool              `json:"cd"`
	Te    bool              `json:"te"`
	Sp    string            `json:"sp"`
	D     interface{}       `json:"d"`
	Bc    string            `json:"bc"`
	Qc    string            `json:"qc"`
	Toa   int64             `json:"toa"`
	Ttp   int               `json:"ttp"`
	W     int               `json:"w"`
	Lr    decimal.Decimal   `json:"lr"`
	Dn    string            `json:"dn"`
}

type GetMarketSymbolsResponse struct {
	Status string         `json:"status"`
	Data   []MarketSymbol `json:"data"`
	Ts     string         `json:"ts"`
	Full   int32          `json:"full"`
}

type MarketSymbol struct {
	Symbol  string          `json:"symbol"`
	State   string          `json:"state"`
	Bc      string          `json:"bc"`
	Qc      string          `json:"qc"`
	Pp      int             `json:"pp"` //交易对价格精度
	Ap      int             `json:"ap"`
	Sp      string          `json:"sp"`
	Vp      int             `json:"vp"`
	Minoa   decimal.Decimal `json:"minoa"`
	Maxoa   decimal.Decimal `json:"maxoa"`
	Minov   decimal.Decimal `json:"minov"`
	Lominoa decimal.Decimal `json:"lominoa"`
	Lomaxoa decimal.Decimal `json:"lomaxoa"`
	Lomaxba decimal.Decimal `json:"lomaxba"`
	Lomaxsa decimal.Decimal `json:"lomaxsa"`
	Smminoa decimal.Decimal `json:"smminoa"`
	Blmlt   decimal.Decimal `json:"blmlt"`
	Slmgt   decimal.Decimal `json:"slmgt"`
	Smmaxoa decimal.Decimal `json:"smmaxoa"`
	Bmmaxov decimal.Decimal `json:"bmmaxov"`
	Msormlt decimal.Decimal `json:"msormlt"`
	Mbormlt decimal.Decimal `json:"mbormlt"`
	Maxov   decimal.Decimal `json:"maxov"`
	U       string          `json:"u"`
	Mfr     decimal.Decimal `json:"mfr"`
	Ct      string          `json:"ct"`
	Rt      string          `json:"rt"`
	Rthr    decimal.Decimal `json:"rthr"`
	In      decimal.Decimal `json:"in"`
	At      string          `json:"at"`
	Tags    string          `json:"tags"`
}
