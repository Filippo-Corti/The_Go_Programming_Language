package main

//Given the pointer to the first element node of a list, returns the number of elements
func Length(first *Node) (c int) {
	var curs *Node
	for curs = first; curs != nil; curs = curs.Next {
		c++
	}
	return
}