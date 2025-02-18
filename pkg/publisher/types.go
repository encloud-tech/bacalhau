package publisher

import (
	"context"

	"github.com/bacalhau-project/bacalhau/pkg/model"
)

// PublisherProvider returns a publisher for the given publisher type
type PublisherProvider = model.Provider[model.Publisher, Publisher]

// Publisher is the interface for publishing results of a job
// The job spec will choose which publisher(s) it wants to use
// (there can be multiple publishers configured)
type Publisher interface {
	model.Providable

	// Validate the job's publisher configuration
	ValidateJob(ctx context.Context, j model.Job) error

	// compute node
	//
	// once the results have been verified we publish them
	// this will result in a "publish" event that will keep track
	// of the details of the storage spec where the results live
	// the returned storage spec might be nill as jobs
	// can have multiple publishers and some publisher
	// implementations don't concern themselves with storage
	// (e.g. notify slack)
	PublishResult(
		ctx context.Context,
		executionID string,
		job model.Job,
		resultPath string,
	) (model.StorageSpec, error)
}
