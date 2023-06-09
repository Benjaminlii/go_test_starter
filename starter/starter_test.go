package starter_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/Benjaminlii/go_test_starter/starter"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	ans := starter.Hello("Benjamin")
	assert.Equal(t, "Hello Benjamin", ans)

	another_ans := starter.Hello("Tao")
	assert.Equal(t, "Hello Tao", another_ans)

}

func TestOddOrEven(t *testing.T) {
	t.Run("Test Positive Number", func(t *testing.T) {
		assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
		assert.Equal(t, "43 is an odd number", starter.OddOrEven(43))
		assert.Equal(t, "42 is an even number", starter.OddOrEven(42))
	})
	t.Run("Test Negative Number", func(t *testing.T) {
		assert.Equal(t, "-45 is an odd number", starter.OddOrEven(-45))
	})
	t.Run("Test Edge Case", func(t *testing.T) {
		assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
	})
}

func TestCheckhealth(t *testing.T) {
  t.Run("Check health status", func(t *testing.T) {
    req := httptest.NewRequest("GET", "http://github.com/", nil)
    writer := httptest.NewRecorder()
    starter.Checkhealth(writer, req)
    response := writer.Result()
    body, err := io.ReadAll(response.Body)

    assert.Equal(t, "health check passed", string(body))
    assert.Equal(t, 200, response.StatusCode)
    assert.Equal(t,
      "text/plain; charset=utf-8",
      response.Header.Get("Content-Type"))

    assert.Equal(t, nil, err)
  })
}

