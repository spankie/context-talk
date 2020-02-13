# Context package

- What is a context value.
- How to create context values.
- How to define functions that takes contexts.
- How context can make your http clients and servers more efficient and less wasteful


Context was added to the go standard library with go1.7. Although it existed as a package in golang.org/x/net/context

The context package is a very small package with just two `types` and three functions.

The `Context` type has three exported methods as well, while the CancelFunc has none.

## Use of context

The main use of the context package is for cancelation and cancelation propagation.

Context is also used to send values. especially when they need to be scoped.

The background context is a bare bone context without any values or timers.

When you create more than one context, you are creating a tree of context. and they all start with the background context.

You create a new context using context.Background();


there are two errors from ctx.Err() from canceled or from deadline exceeded.



## Using context

You can use context to determine when to stop and particular task.

In a web server handler, context can be used to determine when the client
has disconnected, so you can stop what ever is being done and just return an error.

In a client request, it can also be used to stop a request if the request is no
needed.

## Using context values

Like was mentioned earlier, context values should be used for request scoped values
that needs to be passed around. This is because they do not make your code very readable
as they are not seen explicitly, and the data must be safe for access by multiple go routines.

it is a nice practice to use keys that are of an unexported type when being used in a package.
that way, no one can change the values of these keys in the context


> Side Note: You can have two variables with the same address when they are created as zero size in memory in succession.
