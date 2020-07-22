"""The tests module."""

from .main import Player, sort_players


def test_main():
    """Test the main module."""
    players = [
        Player("Smith", 20),
        Player("Jones", 15),
        Player("Jones", 20),
    ]
    assert sort_players(players) == [
        Player("Jones", 20),
        Player("Smith", 20),
        Player("Jones", 15),
    ]
