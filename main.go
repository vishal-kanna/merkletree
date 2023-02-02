package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type Tnode struct {
	left  *Tnode
	right *Tnode
	hash  string
}

//this is store the merkleroot of the tree
type MerkleTree struct {
	head *Tnode
}

func MerkleTreeNode(data []string) *MerkleTree {

	var nodes []*Tnode
	if len(data)%2 != 0 {
		val := data[len(data)-1]
		data = append(data, val)
	}
	for _, val := range data {
		node := &Tnode{
			left:  nil,
			right: nil,
			hash:  hashvalue(val),
		}
		nodes = append(nodes, node)
	}
	for i := 0; i < len(nodes); i++ {
		fmt.Print(nodes[i].hash, " ")
	}
	fmt.Println()
	for i := 0; i < len(data)/2; i++ {
		var nextlevel []*Tnode
		for j := 0; j < len(nodes); j += 2 {
			node := &Tnode{
				left:  nodes[j],
				right: nodes[j+1],
				hash:  hashvalue(nodes[j].hash + nodes[j+1].hash),
			}
			nextlevel = append(nextlevel, node)
		}
		nodes = nextlevel
	}
	return &MerkleTree{head: nodes[0]}
}
func hashvalue(msg string) string {
	h := sha1.New()
	h.Write([]byte(msg))
	return hex.EncodeToString(h.Sum(nil))
}
func (node *Tnode) Find(s string) bool {
	if node.hash == s {
		fmt.Println(node.hash == s)
		return true
	}
	var lc bool
	var rc bool
	if node.left != nil {
		lc = node.left.Find(s)
	}
	if node.right != nil {
		rc = node.right.Find(s)
	}
	return rc || lc
}
func (node *Tnode) printtree(level int) {
	format := ""
	for i := 0; i < level; i++ {
		format += "       "
	}
	format += "---[ "
	level++

	fmt.Println("the hash is ", format, node.hash)
	// fmt.Println("node left add", node.left)
	if node.left != nil {
		node.left.printtree(level)
	}
	if node.right != nil {
		node.right.printtree(level)
	}
}
func Delete(data []string, s string) []string {
	i := -1
	for key, val := range data {
		if val == s {
			i = key
		}
	}
	if i == -1 {
		fmt.Println("the string not match")
		return data
	} else {
		data = append(data[:i], data[i+1:]...)
	}
	fmt.Println("the len of the data is ", len(data))
	return data

}
func main() {
	data := []string{}
Loop:
	for {
		fmt.Println(" 1.ADD the value \n 2.Delete the String \n 3.Verify \n 4.Exit")
		var op int
		fmt.Scan(&op)
		switch op {
		case 1:
			fmt.Println("enter the value to be added")
			var n string
			fmt.Scan(&n)
			data = append(data, n)
			ans := MerkleTreeNode(data)
			ans.head.printtree(0)
		case 2:
			fmt.Println("enter the string to be deleted")
			var s string
			fmt.Scan(&s)
			data = Delete(data, s)
			ans := MerkleTreeNode(data)
			ans.head.printtree(0)
		case 3:
			ans := MerkleTreeNode(data)
			fmt.Println("enter the string to be searched")
			var searched string
			fmt.Scan(&searched)
			b := ans.head.Find(hashvalue(searched))
			if b == true {
				fmt.Println("The string found in the merkle tree")
			} else {
				fmt.Println("the string did not found")
			}
		case 4:
			break Loop
		}
	}
}
