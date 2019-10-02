package helpers

import "github.com/SpectraLogic/ds3_go_sdk/ds3/models"

type ReadTransferStrategy struct {
    BlobStrategy ReadBlobStrategy
    Options      ReadBulkJobOptions
    Listeners    ListenerStrategy
}

// Defines the options to use on the get bulk job
type ReadBulkJobOptions struct {
    Aggregating *bool
    ChunkClientProcessingOrderGuarantee models.JobChunkClientProcessingOrderGuarantee
    ImplicitJobIdResolution *bool
    name *string
    priority models.Priority
}

// Strategy for how to blob objects for writing
type ReadBlobStrategy interface {
    BlobStrategy
}