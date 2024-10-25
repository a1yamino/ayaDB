package file

import "os"

type LogFile struct {
	f *os.File
}

type Options struct {
	name string
}

func OpenLogFile(opt *Options) *LogFile {
	lf := &LogFile{}
	lf.f, _ = os.Create(opt.name)
	return lf
}
