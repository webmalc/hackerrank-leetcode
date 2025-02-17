"""The main module."""
from dataclasses import dataclass
from pprint import pprint


@dataclass
class Pair:
    """The pair class"""
    key: int
    value: str


def get_from_array(data: list[tuple[int, str]]) -> list[Pair]:
    """Create pairs from the array"""
    return [Pair(i[0], i[1]) for i in data]


class Solution:
    """The solution class."""

    def insertion_sort(self, pairs: list[Pair]) -> list[list[Pair]]:
        # pylint: disable=no-self-use
        """Sort the list of pairs."""
        data = []
        for i, _ in enumerate(pairs):
            j = i - 1
            while j >= 0 and pairs[j].key > pairs[j + 1].key:
                pairs[j], pairs[j + 1] = pairs[j + 1], pairs[j]
                j -= 1
            data.append(pairs.copy())
        return data


if __name__ == "__main__":
    solution = Solution()

    for k in [
        [(5, "apple"), (2, "banana"), (9, "cherry")],
        [(3, "cat"), (3, "bird"), (2, "dog")],
    ]:
        result = solution.insertion_sort(get_from_array(k))
        pprint(result)
