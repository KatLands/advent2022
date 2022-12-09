from dataclasses import dataclass


@dataclass()
class Elf:

    def __init__(self, items: list):
        self.items = items

    def __repr__(self):
        return str(self.items)

    def get_items(self) -> list:
        return self.items

    def item_total(self) -> int:
        total = 0
        for item in self.items:
            total += item
        return total






