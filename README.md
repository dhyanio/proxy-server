# proxy-server

In this repo I have created, to use cases of Proxy server [loadbalancer, reverse...]. Both are written in Native Golang.
- Compression server
- File server
- Loadbalancer
- Revese proxy
- Cache service

## Understand Handle, Handler, and HandleFunc in Go
In any programming language to implement a server, we need two important things: <b>Port and routes</b>.
If you ever encountered the implementation of the HTTP server in Golang, you must have faced the ListenAndServe() function. It accepts two parameters 1. Port with a colon like (:8080) and 2. An object of a user-defined type. Sometimes we pass nil as a second parameter and sometimes we pass some parameter why is that so, Any idea?
To know this concept we need to understand 3 terms
- Interface
- Handler
- Handle
- HandleFunc

### Interface
Interface is just to define the behavior of anything like I want to create a very simple mobile game where on-screen there are many items like bird, car, human, etc and on click of any item, you will get the sound of that item like bird mumbles, car horns, and human speaks so all item have common behavior that is everyone makes some sound so I create interface with one method.
```go
    interface voice {
        sound();
    }
```
Now we need to define all struct with the implementation of this interface.
```go
    // Bird Item
    type bird struct {}
    func (b *bird) sound() {
    // Add media of bird voice
    }
    // Car Item
    type car struct {}
    func (c *car) sound() {
    // Add media of car voice
    }
    // human Item
    type human struct {}
    func (h *human) sound() {
    // Add media of human voice
    }
```
Now it’s a piece of cake, we need to create an instance of the interface and assign an object of all struct which has implemented that interface.
```go
    // Create variable of interface
    var v voice
    v := bird object
    v.sound() // Assigning bird object in interface variable
    v := car object
    v.sound() // Assigning car object in interface variable
    v := human object
    v.sound() // Assigning human object in interface variable
```
If you see here we have a variable of the interface and we are assigning objects of all struct that implement that interface. Now a question arises what is the use of this?

Suppose in-game if anyone touches any item (like bird, car, human, etc) and if we create a method that gives an instance of a touched item.
```go
    func getTouchedItemObject() {
    // Return instance of any touched object
    }
    v := getTouchedItemObject();
    v.sound() // it will say sound of any touched thing in the game
```
Now after this code it’s a cup of tea to maintain the code if you want to add a new object in the game go for it, just create a struct of that object and define the method and boom getTouchedItemObject() will give the instance of the object and v.sound() will speak the language of the thing, no need to touch any existing code.

### Handler
The handler is nothing, just an interface with only one method and that method has two parameters: response writer and a pointer of request.
```go
    type Handler interface {
    ServeHTTP(ResponseWriter, *Request) 
    }
```
Now as I mentioned earlier the second parameter of ListenAndServe() is Object of a user-defined type its correct definition is an object of a user-defined type that implements Handler interface.

### Handle
The handle is a function with two parameters: 1. A pattern of the route and 2. An object of a user-defined type that implements the handler interface similar to ServeAndListen().

Now let’s see the code and understand the flow of server
```go
    type person struct { 
    name string
    age int
    }
    // Implement handle interface
    func (p *person) serveHTTP(w http.writeResponse, r *http.Request) {
    // Callback of `/blogs` path
    }
    func main() {
    var p person
    mux := http.NewServeMux()
    mux.handle(“/blogs”, p)
    http.ListenAndServe(“:80”, mux)
    }
```
The above code register `/blogs` path with handler object so whenever the `blogs` path hit the server, handle function search for the path and takes the second parameter value (handler object, p) and calls serveHTTP function, if p doesn’t implement serveHTTP function then code will give an error. Don’t worry about http.NewServeMux() it just a function that returns an object of a struct which implemented handle interface.

Till now we understand the use of handler and handle and with the help of both, we can implement a route system then what is the use of HandleFunc? To understand this we need to understand a situation, let assume we need to build a web application and add 1000 route in our application now for that we need to call 1000 times to handle function to register the path and for each path needs to write serveHTTP function it means for each serveHTTP we need to define 1000 struct type.
```go
    type blog struct {}
    type comment struct {}
    type tag struct {}
    . . .
    . . .
    . . .
    . . .
    func (b *blog) serveHTTP(w http.writeResponse, r *http.Request) { 
    // Callback of /blogs path 
    }
    func (c *comment) serveHTTP(w http.writeResponse, r *http.Request) {
    // Callback of / comments 
    }
    func (t *tag) serveHTTP(w http.writeResponse, r *http.Request) { 
    // Callback of /tag path 
    }
    . . .
    . . .
    . . .
    . . .
    func main() {
    var t tag
    var b blog
    var c comment
    . . .
    . . .
    . . .
    mux := http.NewServeMux()
    mux.handle(“/blogs”, b)
    mux.handle(“/comments”, c)
    mux.handle(“/tags”, t)
    . . .
    . . .
    . . .
    . . .
    http.ListenAndServe(“:80”, mux)
    }
```
There is a lot of chaos in code like first it’s hard to maintain this code and if we need to change anything in callback it’s hard to find which callback belongs to which route to solve this problem here handlerFunc comes in picture.
### HandlerFunc
```go
    type HandlerFunc func(ResponseWriter, *Request) // ServeHTTP calls f(w, r).
    func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) { 
    f(w, r) 
    }
```
HandlerFunc is a user-defined type that implements a handler interface and it makes any normal function as an HTTP handler. It means if I have a normal function with two parameters ResponseWriter, *Request then no need to create a struct
```go
    func getBlogs(w ResponseWriter, r *Request) {
    // Callback of `/blogs` route
    }
```
Now we will use this above function as a callback of /blogs route.
```go
    func main() {
    var b blog
    mux := http.NewServeMux()
    mux.handleFunc(“/blogs”, “getBlogs”)
    http.ListenAndServe(“:80”, mux)
    }
```
Now one question left with the above discusses that sometimes ListenAndServe() has nil as a second parameter. Why is that so? The reason for that is if we pass nil then golang automatically creates an object of http.NewServeMux() which is a default one already defined in http package and passes that object as an argument.



## What is a Proxy Server?
Proxy means someone has the power or authority to do something for someone else or pretends to be someone else. Let us see this with a simple example, if you are living with Bob and currently he is not at home and you are alone and suppose the phone rang and you picked up the call and caller said, “Hi, this is Martin, may I speak to Bob”. As Bob was not at home so you pretended to be Bob and had a discussion with caller on behalf of Bob. So here, you are pretending to be Bob and the caller is not aware that he is not speaking to Bob.

A Proxy server is simply a server that pretends to be an original server for clients.

Forward Proxy
A forward proxy is also called a proxy by many people. This proxy acts on the client-side. In this type of proxy, all requests from a client are sent to the proxy server first, and then that it performs some operations like logging, encryption, decryption, hide IP, etc. A firewall is a kind of forward proxy.
The use of a forward proxy is to bypass a network block i.e., run any web application that is restricted by the network.

