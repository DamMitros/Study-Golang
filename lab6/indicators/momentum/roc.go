package momentum

import (
	"fmt"
	"zad6/data"
)

func CalculateROC(stockData []data.StockData, period int) ([]float64, error) {
  if period <= 0 {
    return nil, fmt.Errorf("period must be positive")
  }
  if len(stockData) < period+1 {
    return nil, fmt.Errorf("not enough data points (%d) for period %d", len(stockData), period)
  }

  rocValues := make([]float64, len(stockData)-period)

  for i := period; i < len(stockData); i++ {
  	closeNPeriodsAgo := stockData[i-period].Close
    if closeNPeriodsAgo == 0 {
      rocValues[i-period] = 0
    } else {
      rocValues[i-period] = ((stockData[i].Close - closeNPeriodsAgo) / closeNPeriodsAgo) * 100
    }
  }

  return rocValues, nil
}