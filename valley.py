#!/usr/bin/python3
# https://www.hackerrank.com/challenges/counting-valleys/problem?


def countingValleys(n, s):
    # n = 8
    # s = UDDDUDUU
    """
    _/\      _
       \    /
        \/\/
    """
    valley = 0
    height = 0
    gone_down = False
    for direction in s:
        if height < 0:
            gone_down = True
        if direction == "U":
            height += 1
        elif direction == "D":
            height -= 1
        if height == 0 and gone_down:
            valley += 1
            gone_down = False
    print(valley)
    return valley


countingValleys(8, "UDDDUDUU")
