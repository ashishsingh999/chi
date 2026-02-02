package chi

import (
	"context"
	"net/http"
	"sync"
)

type GrepTestStruct struct {
	mu    sync.Mutex
	value string
	data  map[string]int
}

func NewGrepTestStruct() *GrepTestStruct {
	return &GrepTestStruct{
		value: "default",
		data:  make(map[string]int),
	}
}

func (g *GrepTestStruct) GetValue() string {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.value
}

func (g *GrepTestStruct) SetValue(v string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.value = v
}

func (g *GrepTestStruct) ProcessData(key string) int {
	g.mu.Lock()
	defer g.mu.Unlock()
	if val, ok := g.data[key]; ok {
		return val
	}
	return 0
}

type GrepTestInterface interface {
	GetValue() string
	SetValue(string)
	ProcessData(string) int
}

type GrepTestType int

const (
	GrepTestTypeA GrepTestType = iota
	GrepTestTypeB
	GrepTestTypeC
)

var GrepTestGlobalVar = "test"

func GrepTestFunction(ctx context.Context, req *http.Request) error {
	_ = ctx
	_ = req
	return nil
}

func GrepTestFunctionWithReturn() (string, error) {
	return "result", nil
}

func GrepTestFunctionWithParams(a int, b string) int {
	return len(b) + a
}

func GrepTestHelper() {
	helper := func() {
	}
	helper()
}

type GrepTestEmbedded struct {
	GrepTestStruct
	extra string
}

func GrepTestComplexFunction(data []byte) (map[string]int, error) {
	result := make(map[string]int)
	for i, b := range data {
		result[string(b)] = i
	}
	return result, nil
}
