package perrors

import "github.com/Southclaws/fault/ftag"

const (
	Internal        ftag.Kind = ftag.Internal
	NotFound        ftag.Kind = ftag.NotFound
	InvalidArgument ftag.Kind = ftag.InvalidArgument
	Unauthorized    ftag.Kind = "UNAUTHORIZED"
)
