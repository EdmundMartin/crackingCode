from collections import defaultdict


def palindrome_permutation_naive(target: str) -> bool:
    frequencies = defaultdict(int)
    for char in target:
        frequencies[char] += 1
    found_odd = False
    for val in frequencies.values():
        if val % 2 == 1:
            if found_odd:
                return False
            found_odd = True
    return True


def palindrom_permutation_optimised(target: str) -> bool:
    frequencies = defaultdict(int)
    odd_count = 0
    for char in target:
        frequencies[char] += 1
        if frequencies[char] % 2 == 1:
            odd_count += 1
        else:
            odd_count -= 1
    return odd_count <= 1
