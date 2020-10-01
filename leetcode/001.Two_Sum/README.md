###Description
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have ***exactly* one solution**, and you may not use the same element twice.

You can return the answer in any order.


**Example 1:**
```markdown
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
```

**Example 2:**
```markdown
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

**Example 3:**
```markdown
Input: nums = [3,3], target = 6
Output: [0,1]
```

**Constraints:**


* 2 <= nums.length <= 105
* -109 <= nums[i] <= 109
* -109 <= target <= 109
* **Only one valid answer exists.**



###Analyze requirements

- Input: nums and target  
- Output: a slice contains two numbers. The sum of two numbers must add up to target  
- Derive from description and constraints: 
    - **2 <= nums.length <= 105** => doesn't need to check if len(nums) < 2
    - It is guaranteed we always can find the two numbers in **nums**
    - There is no guarantee that *nums* does not contain duplicate 
 
###Options:

#### 1. Bruce force 
Time Complexity: O(n^2)  
Space Complexity: O(n)
```go
func twoSum1(nums []int, target int) []int
```
- **Note:**
    - Easy to figure out 
    - O(1) space complexity is an advantage 
    - This is not a "good" solution because of time complexity, but still a correct solution
#### 2. Hashing 
Time complexity: O(n)
Space complexity: O(n)
```go
func twoSum2(nums []int, target int) []int
```
- **Note:**
    - Easy to figure out 
    - O(n) time complexity is an advantage
    - O(n) space complexity probably is a disadvantage 
    - Can use set or other data structure that take O(1) in insert and select 


###Helpful sources
-  [Two Sum](https://web.stanford.edu/class/cs9/sample_probs/TwoSum.pdf)