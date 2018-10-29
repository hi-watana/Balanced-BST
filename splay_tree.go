package main

import "fmt"

type Node struct {
    key int
    left, right *Node
}

func newNode(key int) *Node {
    node := new(Node)
    node.key = key
    node.left = nil
    node.right = nil
    return node
}

func rightRotate(x *Node) *Node {
    y := x.left
    x.left = y.right
    y.right = x
    return y
}

func leftRotate(y *Node) *Node {
    x := y.right
    y.right = x.left
    x.left = y
    return x
}

func splay(root *Node, key int) *Node {
    if root == nil || key == root.key {
        return root
    }

    if key < root.key {
        if root.left == nil {
            return root
        }

        if key < root.left.key {
            root.left.left = splay(root.left.left, key)
            root = rightRotate(root)
        } else if key > root.left.key {
            root.left.right = splay(root.left.right, key)
            if root.left.right != nil {
                root.left = leftRotate(root.left)
            }
        }

        if root.left == nil {
            return root
        }
        return rightRotate(root)
    } else { // key > root.key
        if root.right == nil {
            return root
        }

        if key > root.right.key {
            root.right.right = splay(root.right.right, key)
            root = leftRotate(root)
        } else if key < root.right.key {
            root.right.left = splay(root.right.left, key)
            if root.right.left != nil {
                root.right = rightRotate(root.right)
            }
        }

        if root.right == nil {
            return root
        }
        return leftRotate(root)
    }
}

func search(root *Node, key int) *Node {
    node := splay(root, key)
    if node.key == key {
        return node
    }
    return nil
}

func insert(root *Node, key int) *Node {
    if root == nil {
        return newNode(key)
    }

    root = splay(root, key)

    if key == root.key {
        return root
    }

    node := newNode(key)
    if key < root.key {
        node.left = root.left
        node.right = root
        root.left = nil
    } else { // key > root.key
        node.right = root.right
        node.left = root
        root.right = nil
    }

    return node
}

func erase(root *Node, key int) *Node {
    if root == nil {
        return root
    }

    root = splay(root, key)

    if root.key != key {
        return root
    }

    if root.left == nil {
        return root.right
    }
    leftNode := root.left
    rightNode := root.right
    leftNode = splay(leftNode, key)
    leftNode.right = rightNode
    return leftNode
}

func preOrder(root *Node) {
    if root != nil {
        fmt.Print(root.key, " ")
        preOrder(root.left)
        preOrder(root.right)
    }
}

func main() {
    var root *Node = nil
    var n, m int
    fmt.Scan(&n, &m)
    for i := 0; i < n; i++ {
        var key int
        fmt.Scan(&key)
        root = insert(root, key)
        fmt.Print(key, " ")
    }
    fmt.Println()

    preOrder(root)
    fmt.Println()
    fmt.Println()

    for i := 0; i < m; i++ {
        var key int
        fmt.Scan(&key)
        root = erase(root, key)
        fmt.Print(key, " ")
    }
    fmt.Println()

    preOrder(root)
    fmt.Println()
}
