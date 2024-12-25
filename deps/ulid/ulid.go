package ulid

import (
	"fmt"
	"math/rand"

	"github.com/stuckinforloop/shit/deps/timeutils"

	"github.com/oklog/ulid/v2"
)

type Source struct {
	rnd     *rand.Rand
	nowFunc timeutils.TimeNow
}

func New(rnd *rand.Rand, nowFunc timeutils.TimeNow) *Source {
	src := &Source{
		rnd,
		nowFunc,
	}

	return src
}

func (s *Source) Generate() (string, error) {
	ms := ulid.Timestamp(s.nowFunc())
	ulid, err := ulid.New(ms, s.rnd)
	if err != nil {
		return "", fmt.Errorf("new ulid: %w", err)
	}

	return ulid.String(), nil
}
