package rect_repo

import "context"

type repoKey string

var (
	RectRepoKey = repoKey("rect-repo")
)

func SetRectangleRepoInContext(parentCtx context.Context, rectRepo RectangleRepo) context.Context {
	return context.WithValue(parentCtx, RectRepoKey, rectRepo)
}

func GetRectangleRepoFromContex(ctx context.Context) (RectangleRepo, bool) {

	rectRepo := ctx.Value(RectRepoKey).(RectangleRepo)

	if rectRepo == nil {
		return nil, false
	}

	return rectRepo, true
}
