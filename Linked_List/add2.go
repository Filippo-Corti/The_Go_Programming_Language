package main

//Given the pointer to the first element node of a list and a new string
//Adds the latter in the correct position to maintain the list in (lexicographic) order.
//Returns the pointer to the new First element
func AddInOrder(first *Node, x string) (newFirst *Node){
	var curs, prev *Node
	newNode := new(Node)
	newNode.Value = x

	if first == nil {
		newFirst = newNode
		return
	}
	//List is not empty
	prev = nil
	for curs = first; curs != nil && curs.Value < newNode.Value; curs = curs.Next {
		prev = curs
	}
	//newNode comes before curs
	newNode.Next = curs
	if prev == nil { //Inserting as first element
		newFirst = newNode //newNode comes after nothing 
		return
	}
	//Inserting the newNode between prev and curs
	prev.Next = newNode //newNode comes after prev
	newFirst = first
	return
}