package timeutils

import "time"

var FoundingTime int64 = 1735131386
var FoundingTimeUTC time.Time = time.Unix(FoundingTime, 0).UTC()

type TimeNow func() time.Time
