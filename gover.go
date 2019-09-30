package gover

import (
	"strconv"
	"sync"
	"time"
)

var (
	GitCommit string
	BuildTime string
	Version   string
)

var (
	gitCommitShort string
	buildTime      time.Time
	m              sync.Mutex
)

func GetGitCommitShort() string {
	m.Lock()
	defer m.Unlock()

	if gitCommitShort == "" {
		if len(GitCommit) > 8 {
			gitCommitShort = GitCommit[:8]
		} else {
			gitCommitShort = GitCommit
		}
	}

	return gitCommitShort
}

var timeFormats = []string{
	time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822, time.RFC822Z,
	time.RFC850, time.RFC1123, time.RFC1123Z, time.RFC3339, time.RFC3339Nano,
	time.Stamp, time.StampMilli, time.StampMicro, time.StampNano,
}

func GetBuildTime() time.Time {
	m.Lock()
	defer m.Unlock()

	if buildTime.IsZero() {
		timeInt, err := strconv.ParseInt(BuildTime, 10, 64)
		if err == nil {
			buildTime = time.Unix(timeInt, 0)
		} else {
			for _, layout := range timeFormats {
				buildTime, err = time.Parse(layout, BuildTime)
				if err != nil {
					break
				}
			}
		}
	}
	return buildTime
}
