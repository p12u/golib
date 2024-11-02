package perrors

import "github.com/Southclaws/fault/ftag"

const (
	CodeInternal        ftag.Kind = ftag.Internal
	CodeNotFound        ftag.Kind = ftag.NotFound
	CodeInvalidArgument ftag.Kind = ftag.InvalidArgument
	CodeUnauthorized    ftag.Kind = "UNAUTHORIZED"
)
