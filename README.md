# godatacollections
Data Collections in Go that use generics

## Usage

The interfaces for different structures lives in the root of this project
For flexability interfaces allow for swapping out different implementations without modifying your code

For the concrete implementations they live in the packages named with their type
Utilizing the concrete types allow for a slight performance boost as you don't need the
additional pointer jump that interfaces in go require
However, I would recommend using interfaces unless you have identified that they are your bottleneck with most reward of eliminating

## Why use this library instead of others

Started off with me wanting to be able to have a generic tree that allowed for the type of the value that is
utilized for figuring out where in the tree the data lives to be different than the value being stored. 

The libraries that I came accross that retricted both concepts to be the same type. This works for primatives but what if
I want to have something that is a large struct but has a small id. Sharing the id around in other areas is cheap but the struct is not.

Another aspect of existing libraries is that they require for the type that is being stored to be compareable which again seemed too restrictive.

This library attempts to solve both issues by having where it makes sense a seperate type K that is the small Key for storing a large value T
