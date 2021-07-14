from collections import defaultdict


def permutations_with_sort(first: str, second: str) -> bool:
    if len(first) != len(second):
        return False
    return sorted(first) == sorted(second)


def permutations_count(first: str, second: str) -> bool:
    if len(first) != len(second):
        return False
    letters = defaultdict(int)

    for ch in first:
        letters[ch] += 1

    for ch in second:
        letters[ch] -= 1
        if letters[ch] < 0:
            return False
    return True
