package uniswap

// Option is a function that is used to configure a RequestHandler.
// todo get rid of this if we end up having no options
type Option func(impl *UniswapRequestHandlerImpl)
