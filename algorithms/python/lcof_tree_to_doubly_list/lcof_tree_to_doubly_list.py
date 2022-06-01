"""
# Definition for a Node.
class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
"""

class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def __init__(self):
        self.head = None
        self.cur = None
        self.prev = None

    def treeToDoublyList(self, root: 'Node') -> 'Node':
        self.dfs(root)

        if self.head:
            self.head.left, self.prev.right = self.prev, self.head

        return self.head

    def dfs(self, node: 'Node'):
        if not node:
            return

        self.cur = node
        self.dfs(node.left)
        if not self.prev:
            self.head = node
        else:
            self.prev.right, node.left = node, self.prev

        self.prev = node

        self.dfs(node.right)
