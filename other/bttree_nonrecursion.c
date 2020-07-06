typedef struct BTNode
{

    int data;
    struct BTNode *left;
    struct BTNode *right;

} BTNode;

void preorderWithNonorecursion(BTNode *bt)
{
    if (bt != NULL)
    {
        int top = -1;
        BTNode *Stack[size];
        BTNode *p;
        Stack[++top] = bt;
        while (top != -1)
        {
            p = Stack[top--];
            VisitNode(p);
            if (p->right != NULL)
            {
                Stack[++top] = p->right;
            }
            if (p->left != NULL)
            {
                Stack[++top] = p->left;
            }
        }
    }
}
