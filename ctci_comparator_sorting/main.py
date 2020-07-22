"""The main module."""
from functools import cmp_to_key
from typing import List


class Player:
    """The player comparator class."""

    name: str
    score: int

    def __init__(self, name: str, score: int):
        """Construct the object."""
        self.name = name
        self.score = score

    def __eq__(self, other: object):
        """Check if the objects is match."""
        if not isinstance(other, Player):
            return NotImplemented
        return self.name == other.name and self.score == other.score

    def __repr__(self):
        """Represent the object."""
        return f"name: {self.name}; score: {self.score}"

    @staticmethod
    def comparator(one: "Player", two: "Player"):
        """Compare the provided players the object."""
        if one.score > two.score:
            return -1
        elif one.score < two.score:
            return 1
        if one.name > two.name:
            return 1
        elif one.name < two.name:
            return -1
        return 0


def read_players_from_input() -> List[Player]:
    """Read players data from the input."""
    n = int(input())
    players = []
    for i in range(n):
        name, score = input().split()
        player = Player(name, int(score))
        players.append(player)
    return players


def sort_players(players: List[Player]) -> List[Player]:
    """Return the sorted players list."""
    return sorted(players, key=cmp_to_key(Player.comparator))


if __name__ == "__main__":
    for player in sort_players(read_players_from_input()):
        print(player.name, player.score)
