package file

type Manifest struct {
	f *LogFile
}

func OpenManifestFile(opt *Options) *Manifest {
	return &Manifest{}
}
