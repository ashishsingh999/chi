package chi

import (
	"context"
	"net/http"
	"sync"
)

// TestChunkTypes is a comprehensive test file to verify all chunk types
// are extracted correctly by the Go parser

// Package-level constants (should be extracted as 'module' chunk type)
const (
	TestConst1 = "value1"
	TestConst2 = 42
	TestConst3 = true
)

// Package-level variables (should be extracted as 'module' chunk type)
var (
	TestVar1 string
	TestVar2 int
	TestVar3 bool
)

// SimpleStruct is a basic struct type (should be extracted as 'class' chunk type)
type SimpleStruct struct {
	Field1 string
	Field2 int
	Field3 bool
}

// EmbeddedStruct demonstrates struct embedding (should extract embedded types)
type EmbeddedStruct struct {
	SimpleStruct // Embedded type
	Field4      string
}

// InterfaceType is a basic interface (should be extracted as 'interface' chunk type)
type InterfaceType interface {
	Method1() string
	Method2(newValue string) error
}

// EmbeddedInterface demonstrates interface embedding
type EmbeddedInterface interface {
	InterfaceType // Embedded interface
	Method3() bool
}

// TypeAlias is a type alias (should be extracted as 'type' chunk type)
type TypeAlias = string

// TypeDefinition is a type definition (should be extracted as 'type' chunk type)
type TypeDefinition int

// StandaloneFunction is a package-level function (should be extracted as 'function' chunk type)
func StandaloneFunction(param1 string, param2 int) (string, error) {
	result := param1 + string(rune(param2))
	return result, nil
}

// StandaloneFunctionWithContext demonstrates function with context
func StandaloneFunctionWithContext(ctx context.Context, req *http.Request) error {
	_ = ctx
	_ = req
	return nil
}

// TestStruct is a struct with methods (struct extracted as 'class', methods as 'method')
type TestStruct struct {
	mu    sync.Mutex
	value string
}

// Method1 is a method on TestStruct (should be extracted as 'method' chunk type)
func (t *TestStruct) Method1() string {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.value
}

// Method2 is another method with parameters
func (t *TestStruct) Method2(newValue string) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.value = newValue
	return nil
}

// Method3 demonstrates method with context
func (t *TestStruct) Method3(ctx context.Context) (string, error) {
	_ = ctx
	return t.value, nil
}

// InitFunction is a special init function (should be detected and marked)
func init() {
	TestVar1 = "initialized"
}

// ComplexFunction demonstrates a more complex function with multiple statements
func ComplexFunction(input []string) map[string]int {
	result := make(map[string]int)
	for i, item := range input {
		if item != "" {
			result[item] = i
		}
	}
	return result
}

// FunctionWithErrorHandling demonstrates error handling patterns
func FunctionWithErrorHandling(data []byte) (string, error) {
	if len(data) == 0 {
		return "", nil
	}
	
	result := string(data)
	if len(result) > 100 {
		return result[:100], nil
	}
	
	return result, nil
}

// FunctionWithGoroutine demonstrates goroutine usage
func FunctionWithGoroutine(ch chan string) {
	go func() {
		ch <- "done"
	}()
}

// FunctionWithDefer demonstrates defer usage
func FunctionWithDefer() {
	defer func() {
		// cleanup
	}()
}

// GenericStruct demonstrates generic types (Go 1.18+)
type GenericStruct[T any] struct {
	Value T
}

// GenericMethod demonstrates generic methods
func (g *GenericStruct[T]) GetValue() T {
	return g.Value
}

// FunctionWithVariadic demonstrates variadic functions
func FunctionWithVariadic(args ...string) string {
	result := ""
	for _, arg := range args {
		result += arg
	}
	return result
}

// FunctionWithNamedReturns demonstrates named return values
func FunctionWithNamedReturns(x, y int) (sum int, product int) {
	sum = x + y
	product = x * y
	return
}

// FunctionWithInterfaceParam demonstrates interface parameters
func FunctionWithInterfaceParam(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = r
}

// FunctionWithChannel demonstrates channel usage
func FunctionWithChannel(ch chan int) {
	ch <- 42
	close(ch)
}

// FunctionWithSelect demonstrates select statement
func FunctionWithSelect(ch1, ch2 chan string) string {
	select {
	case msg := <-ch1:
		return msg
	case msg := <-ch2:
		return msg
	default:
		return "default"
	}
}

// FunctionWithSwitch demonstrates switch statement
func FunctionWithSwitch(value int) string {
	switch value {
	case 1:
		return "one"
	case 2:
		return "two"
	default:
		return "other"
	}
}

// FunctionWithTypeAssertion demonstrates type assertions
func FunctionWithTypeAssertion(val interface{}) string {
	if str, ok := val.(string); ok {
		return str
	}
	return "not a string"
}

// FunctionWithTypeSwitch demonstrates type switches
func FunctionWithTypeSwitch(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v
	case int:
		return "int"
	default:
		return "unknown"
	}
}

// FunctionWithMap demonstrates map operations
func FunctionWithMap(data map[string]int) int {
	sum := 0
	for key, value := range data {
		_ = key
		sum += value
	}
	return sum
}

// FunctionWithSlice demonstrates slice operations
func FunctionWithSlice(data []int) []int {
	result := make([]int, 0, len(data))
	for _, v := range data {
		if v > 0 {
			result = append(result, v)
		}
	}
	return result
}

// FunctionWithPointer demonstrates pointer operations
func FunctionWithPointer(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

// FunctionWithRecursion demonstrates recursive functions
func FunctionWithRecursion(n int) int {
	if n <= 1 {
		return 1
	}
	return n * FunctionWithRecursion(n-1)
}

// FunctionWithClosure demonstrates closures
func FunctionWithClosure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

// FunctionWithMultipleReturns demonstrates multiple return values
func FunctionWithMultipleReturns(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, nil
	}
	return a + b, a * b, nil
}

// FunctionWithStructLiteral demonstrates struct literals
func FunctionWithStructLiteral() SimpleStruct {
	return SimpleStruct{
		Field1: "test",
		Field2: 42,
		Field3: true,
	}
}

// FunctionWithInterfaceLiteral demonstrates interface satisfaction
func FunctionWithInterfaceLiteral() InterfaceType {
	ts := &TestStruct{value: "test"}
	_ = ts.Method2("test") // Ensure method exists
	return ts
}

// FunctionWithEmbeddedFieldAccess demonstrates embedded field access
func FunctionWithEmbeddedFieldAccess() EmbeddedStruct {
	return EmbeddedStruct{
		SimpleStruct: SimpleStruct{
			Field1: "embedded",
		},
		Field4: "field4",
	}
}
