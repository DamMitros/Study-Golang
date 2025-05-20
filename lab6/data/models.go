package data

import "time"

type StockData struct {
  Date   time.Time
  Close  float64
  Volume int64
  Open   float64
  High   float64
  Low    float64
}

type FileLoader func(string) ([]StockData, error)

type Indicator struct {
  Name        string
  Description string
  Execute     func([]StockData) (any, error)
}