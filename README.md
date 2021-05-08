# proxy-server

In this repo I have created, to use cases of Proxy server [loadbalancer, reverse]. Both are written in Native Golang.

## Understand Handle, Handler, and HandleFunc in Go
In any programming language to implement a server, we need two important things: Port and routes.
If you ever encountered the implementation of the HTTP server in golang, you must have faced the ListenAndServe() function. It accepts two parameters 1. Port with a colon like (:8023) and 2. An object of a user-defined type. Sometimes we pass nil as a second parameter and sometimes we pass some parameter why is that so, Any idea?
To know this concept we need to understand 3 terms
- Handler
- Handle
- HandleFunc

## What is a Proxy Server?
Proxy means someone has the power or authority to do something for someone else or pretends to be someone else. Let us see this with a simple example, if you are living with Bob and currently he is not at home and you are alone and suppose the phone rang and you picked up the call and caller said, “Hi, this is Martin, may I speak to Bob”. As Bob was not at home so you pretended to be Bob and had a discussion with caller on behalf of Bob. So here, you are pretending to be Bob and the caller is not aware that he is not speaking to Bob.

A Proxy server is simply a server that pretends to be an original server for clients.

Forward Proxy
A forward proxy is also called a proxy by many people. This proxy acts on the client-side. In this type of proxy, all requests from a client are sent to the proxy server first, and then that it performs some operations like logging, encryption, decryption, hide IP, etc. A firewall is a kind of forward proxy.
The use of a forward proxy is to bypass a network block i.e., run any web application that is restricted by the network.

