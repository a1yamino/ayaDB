package file

import "os"

type LogFile struct {
	f *os.File
}

func (lf *LogFile) Close() error {
	return lf.f.Close()
}

func (lf *LogFile) Write(data []byte) error {
	return nil
}

type Options struct {
	name string
}

func OpenLogFile(opt *Options) *LogFile {
	lf := &LogFile{}
	lf.f, _ = os.Create(opt.name)
	return lf
}
