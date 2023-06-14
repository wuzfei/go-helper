package compress

import (
	"os"
	"path/filepath"
	"strings"
)

const pSep = string(os.PathSeparator)

type Match func(filePath string) bool

var DefaultMatch Match = func(filePath string) bool { return true }

func FileMatch(regs ...string) Match {
	return fileMatch(regs, true)
}

func ReFileMatch(regs ...string) Match {
	return fileMatch(regs, false)
}

func fileMatch(regs []string, re bool) Match {
	if len(regs) == 0 {
		return DefaultMatch
	}
	for i := range regs {
		regs[i] = strings.TrimPrefix(regs[i], pSep)
	}
	return func(path string) bool {
		ok := false
		for i := range regs {
			ok, _ = filepath.Match(regs[i], path)
			if ok {
				break
			}
		}
		if re {
			return ok
		}
		return !ok
	}
}
