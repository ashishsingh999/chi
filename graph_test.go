package chi

import (
	"context"
	"net/http"
	"sync"
)

// GraphTestStruct is a struct that should be extracted as 'class' type
type GraphTestStruct struct {
	mu    sync.Mutex
	value string
	data  map[string]int
}

// NewGraphTestStruct creates a new GraphTestStruct instance
// This function should create CALLS relationship to GraphTestStruct
func NewGraphTestStruct() *GraphTestStruct {
	return &GraphTestStruct{
		value: "default",
		data:  make(map[string]int),
	}
}

// GetValue is a method on GraphTestStruct
// Should create CONTAINS relationship from GraphTestStruct to this method
func (g *GraphTestStruct) GetValue() string {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.value
}

// SetValue is another method that should be extracted
func (g *GraphTestStruct) SetValue(v string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.value = v
}

// ProcessData processes data and calls other methods
// Should create CALLS relationships
func (g *GraphTestStruct) ProcessData(key string) int {
	g.mu.Lock()
	defer g.mu.Unlock()
	if val, ok := g.data[key]; ok {
		return val
	}
	return 0
}

// GraphTestInterface defines an interface
// Should be extracted as 'interface' type
type GraphTestInterface interface {
	GetValue() string
	SetValue(string)
	ProcessData(string) int
}

// GraphTestType is a type alias
// Should be extracted as 'type' type
type GraphTestType int

const (
	// GraphTestTypeA is a constant
	GraphTestTypeA GraphTestType = iota
	// GraphTestTypeB is another constant
	GraphTestTypeB
	// GraphTestTypeC is another constant
	GraphTestTypeC
)

// GraphTestGlobalVar is a global variable
// Should be extracted as 'variable' type
var GraphTestGlobalVar = "test"

// GraphTestFunction is a standalone function
// Should be extracted as 'function' type
func GraphTestFunction(ctx context.Context, req *http.Request) error {
	_ = ctx
	_ = req
	return nil
}

// GraphTestFunctionWithReturn returns values
func GraphTestFunctionWithReturn() (string, error) {
	return "result", nil
}

// GraphTestFunctionWithParams has parameters
func GraphTestFunctionWithParams(a int, b string) int {
	return len(b) + a
}

// GraphTestEmbedded embeds GraphTestStruct
// Should create EXTENDS relationship from GraphTestEmbedded to GraphTestStruct
type GraphTestEmbedded struct {
	GraphTestStruct
	extra string
}

// GraphTestComplexFunction demonstrates complex usage
// Should create CALLS relationships to make, len, etc.
func GraphTestComplexFunction(data []byte) (map[string]int, error) {
	result := make(map[string]int)
	for i, b := range data {
		result[string(b)] = i
	}
	return result, nil
}

// GraphTestImplementsStruct implements GraphTestInterface
// Should create IMPLEMENTS relationship (implicit in Go)
type GraphTestImplementsStruct struct {
	value string
}

// GetValue implements GraphTestInterface
func (g *GraphTestImplementsStruct) GetValue() string {
	return g.value
}

// SetValue implements GraphTestInterface
func (g *GraphTestImplementsStruct) SetValue(v string) {
	g.value = v
}

// ProcessData implements GraphTestInterface
func (g *GraphTestImplementsStruct) ProcessData(key string) int {
	return 0
}

// GraphTestUsesType demonstrates USES_TYPE relationships
type GraphTestUsesType struct {
	user *GraphTestStruct
	data map[string]*GraphTestStruct
}

// GraphTestCalls demonstrates CALLS relationships
func GraphTestCalls() {
	instance := NewGraphTestStruct()
	instance.SetValue("test")
	value := instance.GetValue()
	_ = value
}

// GraphTestInstantiation demonstrates INSTANTIATES relationships
func GraphTestInstantiation() {
	// Direct instantiation
	_ = GraphTestStruct{value: "test"}
	
	// Pointer instantiation
	_ = &GraphTestStruct{value: "test"}
	
	// Using new()
	_ = new(GraphTestStruct)
	
	// Embedded struct instantiation
	_ = GraphTestEmbedded{
		GraphTestStruct: GraphTestStruct{value: "embedded"},
		extra:           "extra",
	}
}
