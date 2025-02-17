"""The tests module."""

import pytest

from .main import Solution, get_from_array


@pytest.mark.parametrize("data,expected_data", [
    (
        [(5, "apple"), (2, "banana"), (9, "cherry")],
        [
            [(5, "apple"), (2, "banana"), (9, "cherry")],
            [(2, "banana"), (5, "apple"), (9, "cherry")],
            [(2, "banana"), (5, "apple"), (9, "cherry")],
        ],
    ),
    (
        [(3, "cat"), (3, "bird"), (2, "dog")],
        [
            [(3, "cat"), (3, "bird"), (2, "dog")],
            [(3, "cat"), (3, "bird"), (2, "dog")],
            [(2, "dog"), (3, "cat"), (3, "bird")],
        ],
    ),
])
def test_main(
    data: list[tuple[int, str]],
    expected_data: list[list[tuple[int, str]]],
):
    """Test the main module."""
    solution = Solution()
    pairs = get_from_array(data)
    result = solution.insertion_sort(pairs)
    expected = [get_from_array(i) for i in expected_data]

    assert result == expected
