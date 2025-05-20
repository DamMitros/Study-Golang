package volatility

import (
	"fmt"
	"math"
  "zad6/data"
)

func CalculateATR(stockData []data.StockData, period int) ([]float64, error) {
  if period <= 0 {
    return nil, fmt.Errorf("period must be positive")
  }
  if len(stockData) < period {
    return nil, fmt.Errorf("not enough data points (%d) for period %d", len(stockData), period)
  }

  trueRanges := make([]float64, len(stockData))
  atrValues := make([]float64, len(stockData)-period+1)
  trueRanges[0] = stockData[0].High - stockData[0].Low 

  for i := 1; i < len(stockData); i++ {
    highLow := stockData[i].High - stockData[i].Low
    highPrevClose := math.Abs(stockData[i].High - stockData[i-1].Close)
    lowPrevClose := math.Abs(stockData[i].Low - stockData[i-1].Close)
    trueRanges[i] = math.Max(highLow, math.Max(highPrevClose, lowPrevClose))
  }

  sumTR := 0.0
  for i := range period {
    sumTR += trueRanges[i]
  }
  atrValues[0] = sumTR / float64(period)

  for i := period; i < len(stockData); i++ {
    atrValues[i-period+1] = (atrValues[i-period]*(float64(period-1)) + trueRanges[i]) / float64(period)
  }

  return atrValues, nil
}