# https://leetcode.com/problems/add-two-numbers/description/
# Example 2:
# Input: l1 = [0], l2 = [0]
# Output: [0]

# Example 3:
# Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
# Output: [8,9,9,9,0,0,0,1]


def solution(l1: list, l2: list) -> list:
    j = lambda l: int("".join([str(n) for n in l[::-1]]))
    return [int(n) for n in str(j(l1) + j(l2))[::-1]]


if __name__ == "__main__":
    solution([0], [0])
    l1 = [9, 9, 9, 9, 9, 9, 9]
    l2 = [9, 9, 9, 9]
    print(solution(l1, l2))
