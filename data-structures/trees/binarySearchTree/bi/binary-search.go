package binarySearchTree

//Node represents a node in the binarySearchTree
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

//BinarySearchTree represents a binary search tree
type BinarySearchTree struct {
	Root *Node
}

//Insert, inserts a new value into BST
func (bst *BinarySearchTree) Insert(data int) {
	bst.Root = bst.insertRecursive(bst.Root, data)
}
func (bst *BinarySearchTree) insertRecursive(root *Node, data int) *Node {
	if root == nil {
		return &Node{Data: data}
	}
	if data < root.Data {
		root.Left = bst.insertRecursive(root.Left, data)
	} else if data > root.Data {
		root.Right = bst.insertRecursive(root.Right, data)
	}
	return root
}

//Search, searches for a value in the BST and returns true if found and
//false, otherwise
func (bst *BinarySearchTree) Search(data int) bool {
	return bst.searchRecursive(bst.Root, data)
}
func (bst *BinarySearchTree) searchRecursive(root *Node, data int) bool {
	if root == nil {
		return false
	}
	if data == root.Data {
		return true
	} else if data < root.Data {
		return bst.searchRecursive(root.Left, data)
	} else {
		return bst.searchRecursive(root.Right, data)
	}
}

//Delete, deletes a value from the BST
func (bst *BinarySearchTree) Delete(data int) {
	bst.Root = bst.deleteRecursive(bst.Root, data)
}
func (bst *BinarySearchTree) deleteRecursive(root *Node, data int) *Node {
	if root == nil {
		return root
	}
	if data < root.Data {
		root.Left = bst.deleteRecursive(root.Left, data)
	} else if data > root.Data {
		root.Right = bst.deleteRecursive(root.Right, data)
	} else {
		//Node with only one child or no child
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		//Node with two children, get the inorder successor
		//(min in the right subtree)
		temp := bst.FindMinRecusrsive(root.Right)
		root.Data = temp.Data
		root.Right = bst.deleteRecursive(root.Right, temp.Data)
	}
	return root
}

//FindMin finds the minimum value in the BST.
func (bst *BinarySearchTree) FindMin() int {
	if bst.Root == nil {
		panic("BST is empty")
	}
	return bst.FindMinRecusrsive(bst.Root).Data
}
func (bst *BinarySearchTree) FindMinRecusrsive(root *Node) *Node {
	if root.Left == nil {
		return root
	}
	return bst.FindMinRecusrsive(root.Left)
}

//FindMax, finds the Maximum value in the bst
func (bst *BinarySearchTree) FindMax() int {
	if bst.Root == nil {
		panic("BST is empty")
	}
	return bst.findMaxRecursive(bst.Root).Data
}
func (bst *BinarySearchTree) findMaxRecursive(root *Node) *Node {
	if root.Right == nil {
		return root
	}
	return bst.findMaxRecursive(root.Right)
}
