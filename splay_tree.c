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

struct Node *insert(struct Node *root, int key) {
    if (root == NULL)
        return newNode(key);

    root = splay(root, key);

    if (key < root->key) {
        struct Node *node = newNode(key);
        node->right = root;
        node->left = root->left;
        root->left = NULL;
        return node;
    }
    if (key > root->key) {
        struct Node *node = newNode(key);
        node->left = root;
        node->right = root->right;
        root->right = NULL;
        return node;
    }
    return root;
}

struct Node *maxValueNode(struct Node *root) {
    if (root->right == NULL)
        return root;
    return maxValueNode(root->right);
}

struct Node *delete(struct Node *root, int key) {
    if (root == NULL)
        return root;

    root = splay(root, key);

    if (root->key != key)
        return root;

    struct Node *t1 = root->left, *t2 = root->right;
    free(root);

    if (t1 == NULL)
        return t2;

    t1 = splay(t1, maxValueNode(t1)->key);
    t1->right = t2;
    return t1;
}

void preOrder(struct Node *root) {
    if (root != NULL) {
        printf("%d ", root->key);
        preOrder(root->left);
        preOrder(root->right);
    }
}

int main() {
    struct Node *root = NULL;

    int n, m;
    scanf("%d %d", &n, &m);
    for (int i = 0; i < n; i++) {
        int a;
        scanf(" %d", &a);
        printf("%d ", a);
        root = insert(root, a);
    }

    printf("\n");
    preOrder(root);
    printf("\n\n");

    for (int i = 0; i < m; i++) {
        int a;
        scanf(" %d", &a);
        printf("%d ", a);
        root = delete(root, a);
    }

    printf("\n");
    preOrder(root);
    printf("\n");
    return 0;
}
