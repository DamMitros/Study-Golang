package trend

import (
  "fmt"
  "zad6/data" 
)

func CalculateSMA(stockData []data.StockData, period int) ([]float64, error) {
  if period <= 0 {
    return nil, fmt.Errorf("period must be positive")
  }
  if len(stockData) < period {
    return nil, fmt.Errorf("not enough data points (%d) for period %d", len(stockData), period)
  }

  smaValues := make([]float64, len(stockData)-period+1)
  sum := 0.0

  for i := 0; i < period; i++ {
    sum += stockData[i].Close
  }
  smaValues[0] = sum / float64(period)

  for i := period; i < len(stockData); i++ {
    sum = sum - stockData[i-period].Close + stockData[i].Close
    smaValues[i-period+1] = sum / float64(period)
  }

  return smaValues, nil
}