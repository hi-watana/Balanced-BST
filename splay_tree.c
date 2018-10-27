#include <stdio.h>
#include <stdlib.h>

struct Node {
    int key;
    struct Node *left, *right;
};

struct Node *newNode(int key) {
    struct Node *node = (struct Node*) malloc(sizeof(struct Node));
    node->key = key;
    node->left = node->right = NULL;
    return node;
}

struct Node *rightRotate(struct Node *y) {
    struct Node *x = y->left;
    y->left = x->right;
    x->right = y;
    return x;
}

struct Node *leftRotate(struct Node *x) {
    struct Node *y = x->right;
    x->right = y->left;
    y->left = x;
    return y;
}

struct Node *splay(struct Node *root, int key) {
    if (root == NULL || root->key == key)
        return root;

    if (key < root->key) {
        if (root->left == NULL)
            return root;

        if (key < root->left->key) {
            root->left->left = splay(root->left->left, key);
            root = rightRotate(root);
        } else if (key > root->left->key) {
            root->left->right = splay(root->left->right, key);

            if (root->left->right != NULL)
                root->left = leftRotate(root->left);
        }

        return (root->left == NULL) ? root : rightRotate(root);
    } else {
        if (root->right == NULL)
            return root;

        if (key > root->right->key) {
            root->right->right = splay(root->right->right, key);
            root = leftRotate(root);
        } else if (key < root->right->key) {
            root->right->left = splay(root->right->left, key);
            if (root->right->left != NULL)
                root->right = rightRotate(root->right);
        }

        return (root->right == NULL) ? root : leftRotate(root);
    }
}

struct Node *search(struct Node *root, int key) {
    return splay(root, key);
}

void preOrder(struct Node *root) {
    if (root != NULL) {
        printf("%d ", root->key);
        preOrder(root->left);
        preOrder(root->right);
    }
}

int main() {
    //struct Node *root = NULL;

    //int n, m;
    //scanf("%d %d", &n, &m);
    //for (int i = 0; i < n; i++) {
    //    int a;
    //    scanf(" %d", &a);
    //    root = insert(root, a);
    //}

    //preOrder(root);
    //printf("\n");

    //for (int i = 0; i < m; i++) {
    //    int a;
    //    scanf(" %d", &a);
    //    root = delete(root, a);
    //}

    //preOrder(root);
    //printf("\n");
    struct Node *root = newNode(100);
    root->left = newNode(50);
    root->right = newNode(200);
    root->left->left = newNode(40);
    root->left->left->left = newNode(30);
    root->left->left->left->left = newNode(20);

    root = search(root, 20);
    printf("Preorder traversal of the modified Splay tree is \n");
    preOrder(root);
    return 0;
}
