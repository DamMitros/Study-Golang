package ui

import (
  "bufio"
  "fmt"
  "os"
  "path/filepath"
  "strconv"
  "strings"
  "zad6/data"
  "zad6/indicators/momentum"
  "zad6/indicators/trend"
  "zad6/indicators/volatility"
)

type Console struct {
  reader     *bufio.Reader
  stockData  []data.StockData
  loaders    map[string]data.FileLoader
  indicators map[string]data.Indicator
}

func NewConsole() *Console {
  c := &Console{
    reader:     bufio.NewReader(os.Stdin),
    loaders:    make(map[string]data.FileLoader),
    indicators: make(map[string]data.Indicator),
  }

  c.loaders[".csv"] = data.LoadCSVData
  c.registerIndicators()
  return c
}

func (c *Console) registerIndicators() {
  c.indicators["sma"] = data.Indicator{
    Name:        "Simple Moving Average (SMA)",
    Description: "Calculates average price over specified period",
    Execute: func(data []data.StockData) (any, error) {
      fmt.Print("Enter period for SMA: ")
      periodStr, _ := c.reader.ReadString('\n')
      periodStr = strings.TrimSpace(periodStr)
      period, err := strconv.Atoi(periodStr)
      if err != nil {
        return nil, fmt.Errorf("invalid period: %w", err)
      }
      return trend.CalculateSMA(data, period)
    },
  }

  c.indicators["roc"] = data.Indicator{
  	Name:        "Rate of Change (ROC)",
    Description: "Measures price change as percentage over period",
    Execute: func(data []data.StockData) (any, error) {
      fmt.Print("Enter period for ROC: ")
      periodStr, _ := c.reader.ReadString('\n')
      periodStr = strings.TrimSpace(periodStr)
      period, err := strconv.Atoi(periodStr)
      if err != nil {
        return nil, fmt.Errorf("invalid period: %w", err)
      }
    	return momentum.CalculateROC(data, period)
    },
  }

  c.indicators["atr"] = data.Indicator{
    Name:        "Average True Range (ATR)",
    Description: "Measures market volatility",
    Execute: func(data []data.StockData) (any, error) {
      fmt.Print("Enter period for ATR: ")
      periodStr, _ := c.reader.ReadString('\n')
      periodStr = strings.TrimSpace(periodStr)
      period, err := strconv.Atoi(periodStr)
      if err != nil {
        return nil, fmt.Errorf("invalid period: %w", err)
			}
      return volatility.CalculateATR(data, period)
		},
  }
}

func (c *Console) readInput() string {
  input, _ := c.reader.ReadString('\n')
  return strings.TrimSpace(input)
}

func (c *Console) Start() {
  fmt.Println("=== Stock Market Analysis Tool ===")
  for {
  	fmt.Println("\nMain Menu:")
  	fmt.Println("1. Load data file")
  	fmt.Println("2. Calculate indicators")
  	fmt.Println("3. Display loaded data")
  	fmt.Println("4. Exit")
  	fmt.Print("\nChoose an option (1-4): ")

  	choice := c.readInput()

  	switch choice {
			case "1":
				c.loadDataMenu()
			case "2":
				if len(c.stockData) == 0 {
					fmt.Println("No data loaded. Please load data first.")
					continue
				}
				c.indicatorMenu()
			case "3":
				c.displayData()
			case "4":
				fmt.Println("Exiting...")
				return
			default:
				fmt.Println("Invalid option. Please try again.")
		}
  }
}

func (c *Console) loadDataMenu() {
	fmt.Println("\n=== Load Data ===")
	fmt.Println("1. Use default file (HistoricalData_1747599915878.csv)")
	fmt.Println("2. Select custom file")
	fmt.Print("\nChoose an option (1-2): ")

	choice := c.readInput()

	var filePath string
	switch choice {
    case "1":
			filePath = "HistoricalData_1747599915878.csv"
    case "2":
			fmt.Print("Enter file path: ")
			filePath = c.readInput()
    default:
			fmt.Println("Invalid option. Using default file.")
			filePath = "HistoricalData_1747599915878.csv"
  }

  ext := strings.ToLower(filepath.Ext(filePath))
  loader, ok := c.loaders[ext]
  if !ok {
    fmt.Printf("Unsupported file format: %s\n", ext)
    return
  }

  data, err := loader(filePath)
  if err != nil {
    fmt.Printf("Error loading data: %s\n", err)
    return
  }

  c.stockData = data
  fmt.Printf("Successfully loaded %d records from %s\n", len(data), filePath)
}

func (c *Console) indicatorMenu() {
  fmt.Println("\n=== Calculate Indicators ===")
  fmt.Println("Available indicators:")

	options := make([]string, 0, len(c.indicators))
	i := 1
	for key, ind := range c.indicators {
		fmt.Printf("%d. %s - %s\n", i, ind.Name, ind.Description)
		options = append(options, key)
		i++
	}

	fmt.Print("\nSelect indicator (1-", len(options), "): ")
	choiceStr := c.readInput()
	choice, err := strconv.Atoi(choiceStr)
	if err != nil || choice < 1 || choice > len(options) {
		fmt.Println("Invalid selection.")
		return
	}

	selectedKey := options[choice-1]
	indicator := c.indicators[selectedKey]
	result, err := indicator.Execute(c.stockData)
	if err != nil {
		fmt.Printf("Error calculating %s: %s\n", indicator.Name, err)
		return
	}

	fmt.Printf("\n=== %s Results ===\n", indicator.Name)
	fmt.Println(formatIndicatorResult(result))
}

func (c *Console) displayData() {
	if len(c.stockData) == 0 {
		fmt.Println("No data loaded.")
		return
	}

	fmt.Println("\n=== Stock Data ===")
	fmt.Println("Date\t\tOpen\tHigh\tLow\tClose\tVolume")
	fmt.Println("-----------------------------------------------------")

	limit := min(len(c.stockData), 10)
	for i := range limit {
		record := c.stockData[i]
		fmt.Printf("%s\t%.2f\t%.2f\t%.2f\t%.2f\t%d\n",
			record.Date.Format("2006-01-02"),
			record.Open,
			record.High,
			record.Low,
			record.Close,
			record.Volume)
	}

	if len(c.stockData) > limit {
		fmt.Printf("\n... and %d more records (showing first %d)\n", len(c.stockData)-limit, limit)
	}
}

func formatIndicatorResult(result any) string {
	switch v := result.(type) {
		case []float64:
			if len(v) > 10 {
				truncated := v[:10]
				return fmt.Sprintf("%v\n... (showing first 10 of %d values)", truncated, len(v))
			}
			return fmt.Sprintf("%v", v)
		case map[string]float64:
			var sb strings.Builder
			for date, value := range v {
				sb.WriteString(fmt.Sprintf("%s: %.4f\n", date, value))
			}
			return sb.String()
		default:
			return fmt.Sprintf("%v", v)
  }
}
