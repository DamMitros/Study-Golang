package data

import (
  "encoding/csv"
	"strconv"
  "strings"
  "time"
  "fmt"
  "io"
  "os"
)

func LoadCSVData(filePath string) ([]StockData, error) {
  file, err := os.Open(filePath)
  if err != nil {
    return nil, fmt.Errorf("could not open CSV file %s: %w", filePath, err)
  }
  defer file.Close()

  reader := csv.NewReader(file)
  header, err := reader.Read() 
  if err != nil {
    return nil, fmt.Errorf("could not read header row: %w", err)
  }

  expectedHeader := []string{"Date", "Close/Last", "Volume", "Open", "High", "Low"}
  if len(header) != len(expectedHeader) {
    fmt.Printf("CSV header mismatch: Expected %d columns, got %d.\n", len(expectedHeader), len(header))
  }

  var records []StockData
  lineNumber := 1 

  for {
    lineNumber++
    row, err := reader.Read()
    if err == io.EOF {break}
    if err != nil {return nil, fmt.Errorf("error reading CSV row %d: %w", lineNumber, err)}
    if len(row) < 6 {
      return nil, fmt.Errorf("error parsing CSV row %d: expected at least 6 columns, got %d", lineNumber, len(row))
    }
        
    date, err := time.Parse("01/02/2006", row[0])
    if err != nil {
      return nil, fmt.Errorf("error parsing date on row %d ('%s'): %w", lineNumber, row[0], err)
    }

    closePriceStr := strings.Replace(row[1], "$", "", -1)
    closePrice, err := strconv.ParseFloat(closePriceStr, 64)
    if err != nil {
      return nil, fmt.Errorf("error parsing close price on row %d ('%s'): %w", lineNumber, row[1], err)
    }

    volume, err := strconv.ParseInt(row[2], 10, 64)
    if err != nil {
      return nil, fmt.Errorf("error parsing volume on row %d ('%s'): %w", lineNumber, row[2], err)
    }

    openPriceStr := strings.Replace(row[3], "$", "", -1)
    openPrice, err := strconv.ParseFloat(openPriceStr, 64)
    if err != nil {
      return nil, fmt.Errorf("error parsing open price on row %d ('%s'): %w", lineNumber, row[3], err)
    }

    highPriceStr := strings.Replace(row[4], "$", "", -1)
    highPrice, err := strconv.ParseFloat(highPriceStr, 64)
    if err != nil {
      return nil, fmt.Errorf("error parsing high price on row %d ('%s'): %w", lineNumber, row[4], err)
    }

    lowPriceStr := strings.Replace(row[5], "$", "", -1)
    lowPrice, err := strconv.ParseFloat(lowPriceStr, 64)
    if err != nil {
      return nil, fmt.Errorf("error parsing low price on row %d ('%s'): %w", lineNumber, row[5], err)
    }

    records = append(records, StockData{
      Date:   date,
      Close:  closePrice,
      Volume: volume,
      Open:   openPrice,
      High:   highPrice,
      Low:    lowPrice,
    })
  }

  for i, j := 0, len(records)-1; i < j; i, j = i+1, j-1 {
    records[i], records[j] = records[j], records[i]
  }
  return records, nil
}