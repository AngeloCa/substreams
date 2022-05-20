package substreams

import (
	"context"

	pbsubstreams "github.com/streamingfast/substreams/pb/sf/substreams/v1"
)

type ResponseFunc func(resp *pbsubstreams.Response) error

func NewBlockScopedDataResponse(in *pbsubstreams.BlockScopedData) *pbsubstreams.Response {
	return &pbsubstreams.Response{
		Message: &pbsubstreams.Response_Data{Data: in},
	}
}

func NewModulesProgressResponse(in []*pbsubstreams.ModuleProgress) *pbsubstreams.Response {
	return &pbsubstreams.Response{
		Message: &pbsubstreams.Response_Progress{Progress: &pbsubstreams.ModulesProgress{Modules: in}},
	}
}

type BlockHook func(ctx context.Context, clock *pbsubstreams.Clock) error
type PostJobHook func(ctx context.Context, clock *pbsubstreams.Clock) error
