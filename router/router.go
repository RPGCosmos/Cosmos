package router

type Input struct {
	args []string
}

type Action func(inp Input) error

type Router struct {
	routes map[string]Action
}

func NewRouter() Router {
	return Router{
		routes: make(map[string]Action),
	}
}

func (r *Router) Register(route string, caller Action) error {
	r.routes[route] = caller

	return nil
}

func (r *Router) Unregister(route string) error {
	delete(r.routes, route)

	return nil
}

func (r Router) Handle(route string, inp Input) error {
	return r.routes[route](inp)
}