package main

type dbState struct {
	StateID    int64 `gorm:"PRIMARY_KEY"`
	Open       bool
	LastChange int64
}

type dbTemperature struct {
	Value    float32
	Unit     string
	Location string `gorm:"PRIMARY_KEY"`
	Changed  int64
}

type dbHumidity struct {
	Value    float32
	Unit     string
	Location string `gorm:"PRIMARY_KEY"`
	Changed  int64
}
