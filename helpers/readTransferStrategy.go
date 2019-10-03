package helpers

import "github.com/SpectraLogic/ds3_go_sdk/ds3/models"

type ReadTransferStrategy struct {
    BlobStrategy ReadBlobStrategy
    Options      ReadBulkJobOptions
}

// Defines the options to use on the get bulk job
type ReadBulkJobOptions struct {
    Aggregating *bool
    ChunkClientProcessingOrderGuarantee models.JobChunkClientProcessingOrderGuarantee
    ImplicitJobIdResolution *bool
    Name *string
    Priority models.Priority
}

// Strategy for how to blob objects for writing
type ReadBlobStrategy interface {
    BlobStrategy
}