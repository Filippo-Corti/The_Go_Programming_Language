//Provides a definition of the basic datatype required for lists
package main

type Node struct {
	Value string 
	Next *Node //This is a recursive datatype
}