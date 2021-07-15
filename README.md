# Leetcode: Merge K Sorted Lists

## introduction

This repository for my personal design solution for solve the leetcode Merge K Sorted List in golang

## Problem Description

You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

## Examples:

### Example 1:
```
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
```
### Example 2:
```
Input: lists = []
Output: []
```
### Example 3:
```
Input: lists = [[]]
Output: []
```
## My First Solution
### Idea Description
Because Each list is sorted

Each time take each non-empty list first element to find the smallest 

move this smallest from the list to the merged list

until all the list element has been move the the new merge list

### Time Complexity Analysis

because each time has k element to find smallest -> Compare k times

and total need to find N elements

consume : k * N

## Second Solution
We could refine the My First Solution with Priority Queue(Min-Heap)

This will reduce the each compare time to log(k)

and total need to find N elements

consume : log(k) * N