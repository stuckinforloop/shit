package ulid

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stuckinforloop/shit/deps/timeutils"

	"github.com/stretchr/testify/require"
)

var (
	testULIDOne = "01JFYY7M4G06AFVGQT5ZYC0GEK"
	testULIDTwo = "01JFYY7M4GZW908PVKS1Q4ZYAZ"
)

func TestULID(t *testing.T) {
	rnd := rand.New(rand.NewSource(0))
	var timeNow timeutils.TimeNow = func() time.Time {
		return timeutils.FoundingTimeUTC
	}
	source := New(rnd, timeNow)

	idOne, err := source.Generate()
	require.NoError(t, err)
	require.Equal(t, testULIDOne, idOne)

	idTwo, err := source.Generate()
	require.NoError(t, err)
	require.Equal(t, testULIDTwo, idTwo)
}
