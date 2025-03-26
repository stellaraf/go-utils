package ttime_test

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.stellar.af/go-utils/ttime"
)

func Test_PossibleTime(t *testing.T) {
	type T struct {
		Time ttime.PossibleTime `json:"time"`
	}
	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		var tt *T
		data := []byte(`{"time":"2025-03-26T08:04:19Z"}`)
		err := json.Unmarshal(data, &tt)
		require.NoError(t, err)
	})
	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		var tt *T
		data := []byte(`{"time":""}`)
		err := json.Unmarshal(data, &tt)
		require.NoError(t, err)
		assert.True(t, tt.Time.IsZero())
	})
	t.Run("null", func(t *testing.T) {
		t.Parallel()
		var tt *T
		data := []byte(`{"time":null}`)
		err := json.Unmarshal(data, &tt)
		require.NoError(t, err)
		assert.True(t, tt.Time.IsZero())
	})
}

func Test_Timestamp(t *testing.T) {
	t.Run("base", func(t *testing.T) {
		t.Parallel()
		now := time.Now()
		tsb := []byte(strconv.FormatInt(now.Unix(), 10))
		var val *ttime.Timestamp
		err := json.Unmarshal(tsb, &val)
		require.NoError(t, err)
		assert.Equal(t, now.Unix(), val.Unix())
	})
	t.Run("raw", func(t *testing.T) {
		tsb := []byte(`1722444364652000000`)
		var val *ttime.Timestamp
		err := json.Unmarshal(tsb, &val)
		require.NoError(t, err)
		assert.Equal(t, 2024, val.Year(), "year")
		assert.Equal(t, 7, int(val.Month()), "month")
		assert.Equal(t, 31, val.Day(), "day")
		assert.Equal(t, 16, val.Hour(), "hour")
		assert.Equal(t, 46, val.Minute(), "minute")
		assert.Equal(t, 4, val.Second(), "second")
	})
}
