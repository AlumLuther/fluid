package dataload

// DataLoadValue defines all configurations used in DataLoad prefetch batch job
type DataLoadValue struct {
	DataLoadInfo DataLoadInfo `yaml:"dataloader,omitempty"`
}

type DataLoadInfo struct {
	BackoffLimit int32  `yaml:"backoffLimit,omitempty"`
	Threads      int32  `yaml:"threads,omitempty"`
	MountPath    string `yaml:"mountPath,omitempty"`
	Image        string `yaml:"image"`
	NumWorker    int32  `yaml:"numWorker,omitempty"`
}
