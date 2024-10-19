package middleware

type OapiStrictMiddlewareFunc func(f StrictHandlerFunc, operationId string) StrictHandlerFunc

var DefaultOapiStrictMiddlewares = []OapiStrictMiddlewareFunc{
	OapiLoggerMiddleware,
}
