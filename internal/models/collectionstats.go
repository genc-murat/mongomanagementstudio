package models

type CollectionStats struct {
	NS              string     `json:"ns,omitempty"`
	Size            int64      `json:"size,omitempty"`
	Count           int64      `json:"count,omitempty"`
	AvgObjSize      int64      `json:"avgObjSize,omitempty"`
	StorageSize     int64      `json:"storageSize,omitempty"`
	FreeStorageSize int64      `json:"freeStorageSize,omitempty"`
	Capped          bool       `json:"capped,omitempty"`
	Max             int64      `json:"max,omitempty"`
	MaxSize         int64      `json:"maxSize,omitempty"`
	WiredTiger      WiredTiger `json:"wiredTiger,omitempty"`
}

type WiredTiger struct {
	BlockManager BlockManager `json:"block-manager,omitempty"`
}

type BlockManager struct {
	AllocationsRequiringFileExtension int64 `json:"allocations requiring file extension,omitempty"`
	BlocksAllocated                   int64 `json:"blocks allocated,omitempty"`
	BlocksFreed                       int64 `json:"blocks freed,omitempty"`
	CheckpointSize                    int64 `json:"checkpoint size,omitempty"`
	FileAllocationUnitSize            int64 `json:"file allocation unit size,omitempty"`
	FileBytesAvailableForReUse        int64 `json:"files bytes available for reuse,omitempty"`
	FileMagicNumber                   int64 `json:"file magic number,omitempty"`
	FileMajorVersionNumber            int64 `json:"file major version number,omitempty"`
	FileSizeInBytes                   int64 `json:"file size in bytes,omitempty"`
	MinorVersionNumber                int64 `json:"minor version number,omitempty"`
}
