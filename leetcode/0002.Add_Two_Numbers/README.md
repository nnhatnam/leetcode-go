### Description
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

```markdown
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```

### Analyze requirements

- Input: **non-empty** linked lists
- Output: a **non-empty** linked list
- Derive from description and constraints: 
    - **non-empty linked lists** => Don't need to check if linked lists is nil 
    - **non-negative** integers => don't need to worry about negative number, be careful that a node can be 0
    - Be careful about leading zero and number 0 => Edge cases: 0 + 0
    - There is no limit of linked list size => can't convert to number, add them up and convert it back to list due to overflow
    if the list is too big 
 
 ### Options
 
 #### 1. Iterative
 ```markdown
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode
```
 
 ### 2. Recursive 
```markdown
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode 
```

or 

```markdown
func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode 
```