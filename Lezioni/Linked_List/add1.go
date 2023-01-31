package main

//Given the pointer to the first element node of a list and a new string
//Adds the latter as the FIRST element of the list.
//Returns the pointer to the new First element
func AddFront(first *Node, x string) (newFirst *Node){
	newFirst = new(Node)
	newFirst.Value = x
	newFirst.Next = first
	return
}

//Given the pointer to the first element node of a list and a new string
//Adds the latter as the LAST element of the list.
//Returns the pointer to the new First element (It can be useful is first is nil)
func AddLast(first *Node, x string) (newFirst *Node) {
	var curs *Node
	newNode := new(Node)
	newNode.Value = x

	if first == nil {
		newFirst = newNode
		return
	}
	//List is non-empty
	newFirst = first
	for curs = first; curs.Next != nil; curs = curs.Next {}
	//Curs is now pointing to the last node of the list
	curs.Next = newNode
	return
}