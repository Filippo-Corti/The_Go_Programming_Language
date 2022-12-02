package main

//Given the pointers to the first element of two linked lists,
//appends all the elements of the second list to the tail of the first list.
//Returns the pointer to the new first element of the first list 
func Concat(first1, first2 *Node) (newFirst *Node) {
	var curs1, curs2 *Node
	//If the first list is nil, the new list is the second list
	if first1 == nil {
		newFirst = first2
		return
	}
	for curs1 = first1; curs1.Next != nil; curs1 = curs1.Next {} 
	//curs1 is now pointing to the last element of the first list
	for curs2 = first2; curs2 != nil; curs2 = curs2.Next {
		curs1.Next = curs2
		curs1 = curs2
	}
	newFirst = first1
	return
}