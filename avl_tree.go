package main

import "fmt"

type Node struct {
    key int
    left, right *Node
    height int
}

func newNode(key int) *Node {
    node := new(Node)
    node.key = key
    node.left = nil
    node.right = nil
    node.height = 1
    return node
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func minValueNode(root *Node) *Node {
    if root.left == nil {
        return root
    }
    return minValueNode(root.left)
}

func height(root *Node) int {
    if root == nil {
        return 0
    }
    return root.height
}

func getBalance(root *Node) int {
    if root == nil {
        return 0
    }
    return height(root.left) - height(root.right)
}

func rightRotate(x *Node) *Node {
    y := x.left
    t2 := y.right

    x.left = t2
    y.right = x

    x.height = max(height(x.left), height(x.right)) + 1
    y.height = max(height(y.left), height(y.right)) + 1

    return y
}

func leftRotate(y *Node) *Node {
    x := y.right
    t2 := x.left

    y.right = t2
    x.left = y

    y.height = max(height(y.left), height(y.right)) + 1
    x.height = max(height(x.left), height(x.right)) + 1

    return x
}

func insert(root *Node, key int) *Node {
    if root == nil {
        return newNode(key)
    }

    if root.key == key {
        return root
    }

    if key < root.key {
        root.left = insert(root.left, key)
    } else {
        root.right = insert(root.right, key)
    }

    root.height = max(height(root.left), height(root.right)) + 1

    balance := getBalance(root)

    if balance > 1 {
        if key < root.left.key {
            return rightRotate(root)
        }
        if key > root.left.key {
            root.left = leftRotate(root.left)
            return rightRotate(root)
        }
    } else if balance < -1 {
        if key > root.right.key {
            return leftRotate(root)
        }
        if key < root.right.key {
            root.right = rightRotate(root.right)
            return leftRotate(root)
        }
    }

    return root
}

func erase(root *Node, key int) *Node {
    if root == nil {
        return root
    }

    if key < root.key {
        root.left = erase(root.left, key)
    } else if key > root.key {
        root.right = erase(root.right, key)
    } else {
        if root.left == nil {
            if root.right == nil {
                root = nil
            } else {
                root = root.right
            }
        } else {
            if root.left == nil {
                root = root.left
            } else {
                rightMinNode := minValueNode(root.right)
                root.key = rightMinNode.key
                root.right = erase(root.right, root.key)
            }
        }
    }

    if root == nil {
        return root
    }

    root.height = max(height(root.right), height(root.left)) + 1

    balance := getBalance(root)

    if balance > 1 {
        if getBalance(root.left) < 0 {
            root.left = leftRotate(root.left)
        }
        return rightRotate(root)
    } else if balance < -1 {
        if getBalance(root.right) > 0 {
            root.right = rightRotate(root.right)
        }
        return leftRotate(root)
    }

    return root
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
    }

    preOrder(root)
    fmt.Println()

    for i := 0; i < m; i++ {
        var key int
        fmt.Scan(&key)
        root = erase(root, key)
    }

    preOrder(root)
    fmt.Println()
}
