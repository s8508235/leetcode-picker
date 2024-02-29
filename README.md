## Overview
Fetch leetcode problems with rating(likes/likes+dislikes) and difficulty

## Install
```
go install github.com/s8508235/leetcode-picker/cmd/leetcode-picker@latest
```
## Argument
```
-r, --rating int
0(Negative): rating > 0%
1(MostlyNegative): rating > 20%
2(Mixed): rating > 40%
3(MostlyPositive): rating > 70%
4(Positive): rating > 80%
5(OverwhelminglyPositive): rating > 95%
```
```
-l, --level string
all,a
easy,e
medium,m
normal,n
hard,h
```
## TODO
- optimiztion for speed
- user login to skip submitted problems
- [Hyperlinks in terminal](https://gist.github.com/egmontkob/eb114294efbcd5adb1944c9f3cb5feda)

## Example
```
leetcode-picker -l e -r 1 #difficulty easy and rating over 20%
leetcode-picker -l m -r 4 #difficulty medium and rating over 80%
leetcode-picker -l a -r 5 #difficulty easy,medium,hard and rating over 95%
```
