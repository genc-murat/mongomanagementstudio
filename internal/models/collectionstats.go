package models

type CollectionStats struct {
	Ns              string `json:"ns,omitempty"`
	Size            int    `json:"size,omitempty"`
	Count           int    `json:"count,omitempty"`
	AvgObjSize      int    `json:"avgObjSize,omitempty"`
	StorageSize     int    `json:"storageSize,omitempty"`
	FreeStorageSize int    `json:"freeStorageSize,omitempty"`
	Capped          bool   `json:"capped,omitempty"`
	WiredTiger      struct {
		Metadata struct {
			FormatVersion int `json:"formatVersion,omitempty"`
		} `json:"metadata,omitempty"`
		CreationString string `json:"creationString,omitempty"`
		Type           string `json:"type,omitempty"`
		URI            string `json:"uri,omitempty"`

		BlockManager struct {
			AllocationsRequiringFileExtension int `json:"allocations requiring file extension,omitempty" bson:"allocations requiring file extension,omitempty"`
			BlocksAllocated                   int `json:"blocks allocated,omitempty" bson:"blocks allocated,omitempty"`
			BlocksFreed                       int `json:"blocks freed,omitempty" bson:"blocks freed,omitempty"`
			CheckpointSize                    int `json:"checkpoint size,omitempty" bson:"checkpoint size,omitempty"`
			FileAllocationUnitSize            int `json:"file allocation unit size,omitempty" bson:"file allocation unit size,omitempty"`
			FileBytesAvailableForReuse        int `json:"file bytes available for reuse,omitempty" bson:"file bytes available for reuse,omitempty"`
			FileMagicNumber                   int `json:"file magic number,omitempty" bson:"file magic number,omitempty"`
			FileMajorVersionNumber            int `json:"file major version number,omitempty" bson:"file major version number,omitempty"`
			FileSizeInBytes                   int `json:"file size in bytes,omitempty" bson:"file size in bytes,omitempty"`
			MinorVersionNumber                int `json:"minor version number,omitempty" bson:"minor version number,omitempty"`
		} `json:"block-manager,omitempty" bson:"block-manager,omitempty"`
	} `json:"wiredTiger,omitempty" bson:"wiredTiger,omitempty"`
	Nindexes       int `json:"nindexes,omitempty"`
	TotalIndexSize int `json:"totalIndexSize,omitempty"`
	TotalSize      int `json:"totalSize,omitempty"`

	ScaleFactor int `json:"scaleFactor,omitempty"`
	Ok          int `json:"ok,omitempty"`
}
