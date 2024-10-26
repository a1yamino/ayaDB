package file

import "ayaDB/pkg/codec"

type WalFile struct {
	file *LogFile
}

func (wf *WalFile) Close() error { return wf.file.Close() }

func (wf *WalFile) Write(entry *codec.Entry) error {
	walData := codec.WalCodec(entry)
	return wf.file.Write(walData)
}

func OpenWalFile(opt *Options) *WalFile { return &WalFile{file: OpenLogFile(opt)} }
