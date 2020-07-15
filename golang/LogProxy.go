package utils

import "fmt"

type LogProxy struct{}

func (l *LogProxy) Error(any ...interface{}) {
	if any != nil {
		for _, item := range any {
			fmt.Print(item)
		}
		fmt.Println()
	}
}

func (l *LogProxy) Fatal(any ...interface{}) {
	if any != nil {
		for _, item := range any {
			fmt.Print(item)
		}
		fmt.Println()
	}
}

func (l *LogProxy) Info(any ...interface{}) {
	if any != nil {
		for _, item := range any {
			fmt.Print(item)
		}
		fmt.Println()
	}
}

func (l *LogProxy) Debug(any ...interface{}) {
	if any != nil {
		for _, item := range any {
			fmt.Print(item)
		}
		fmt.Println()
	}
}

func (l *LogProxy) Warn(any ...interface{}) {
	if any != nil {
		for _, item := range any {
			fmt.Print(item)
		}
		fmt.Println()
	}
}

func (l *LogProxy) Trace(any ...interface{}) {
	if any != nil {
		for _, item := range any {
			fmt.Print(item)
		}
		fmt.Println()
	}
}
