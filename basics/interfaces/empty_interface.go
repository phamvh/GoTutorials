package main

/*
EmptyInterface
Empty interface has no method.
An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type.
For example, fmt.Print takes any number of arguments of type interface{}.

Note that Go has an interface, named any, which is an empty interface. And Println() actually takes args of type "any".
*/
type EmptyInterface interface{}
