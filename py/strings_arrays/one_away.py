

def one_away(first: str, second: str):
    if len(first) == len(second):
        return _one_edit_replace(first, second)
    elif len(first) - 1 == len(second):
        return _one_edit_insert(first, second)
    elif len(first) + 1 == len(second):
        return _one_edit_insert(second, first)
    return False


def _one_edit_replace(first: str, second: str) -> bool:
    found_difference = False
    for idx, char in enumerate(first):
        if first[idx] != second[idx]:
            if found_difference:
                return False
            found_difference = True
    return True


def _one_edit_insert(first: str, second: str) -> bool:
    first_index = 0
    second_index = 0
    while second_index < len(second) and first_index < first_index:
        if first[first_index] != second[second_index]:
            if first_index != second_index:
                return False
            second_index += 1
        else:
            first_index += 1
            second_index += 1
    return True