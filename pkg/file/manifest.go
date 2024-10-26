package file

type Manifest struct {
	f *LogFile
}

func (m *Manifest) Close() error {
	return m.f.Close()
}

func OpenManifestFile(opt *Options) *Manifest {
	return &Manifest{}
}
