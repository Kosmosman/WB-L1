### 1. Какой самый эффективный способ конкатенации строк?
Лучше всего, согласно бенчмаркам, показал себя метод strings.Join

Тесты:
```go
package tests

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

var (
	str = []string{
		GenerateLongString(),
		GenerateLongString(),
		GenerateLongString(),
	}
)

func GenerateLongString() string {
	b := make([]byte, 1000000)
	for i := 0; i < 1000000; i++ {
		b[i] = byte(rand.Intn(80) + '!')
	}
	return string(b)
}

func BenchmarkConcatWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Join(str, "")
	}
}

func BenchmarkConcatWithSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s", str[0], str[1], str[2])
	}
}

func BenchmarkConcatWithPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = str[0] + str[1] + str[2]
	}
}

func BenchmarkConcatWithBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := strings.Builder{}
		builder.WriteString(str[0])
		builder.WriteString(str[1])
		builder.WriteString(str[2])
		_ = builder.String()
	}
}

func BenchmarkConcatInByteView(b *testing.B) {
	byteString := make([]byte, 3000001)
	for i := 0; i < b.N; i++ {
		pos := 0
		for _, s := range str {
			pos += copy(byteString[pos:], s)
		}
		_ = string(byteString[:])
	}
}
```
Результаты:
```
/tmp/GoLand/___gobench_WB_L1_Questions_bech.test -test.v -test.paniconexit0 -test.bench . -test.run ^$
goos: linux
goarch: amd64
pkg: WB-L1/Questions/bech
cpu: 11th Gen Intel(R) Core(TM) i7-11800H @ 2.30GHz
BenchmarkConcatWithJoin
BenchmarkConcatWithJoin-16       	    7737	    151372 ns/op
BenchmarkConcatWithSprintf
BenchmarkConcatWithSprintf-16    	    1104	   1375307 ns/op
BenchmarkConcatWithPlus
BenchmarkConcatWithPlus-16       	    6594	    187880 ns/op
BenchmarkConcatWithBuilder
BenchmarkConcatWithBuilder-16    	    2636	    430146 ns/op
BenchmarkConcatInByteView
BenchmarkConcatInByteView-16     	    3681	    329147 ns/op
PASS
```

### 2. Что такое интерфейсы, как они применяются в Go?