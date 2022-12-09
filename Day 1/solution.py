from elf import Elf
from itertools import groupby

elf_list = []


def read_file_input():
    items = []
    with open("input.txt", "r") as file:
        for line in file:
            items.append(line.strip())
        split_list = [list(group) for key, group in groupby(items, key=bool) if key]
        for item in split_list:
            int_list = [eval(i) for i in item]
            create_elf(int_list)


def create_elf(items):
    elf = Elf(items)
    elf_list.append(elf)


if __name__ == '__main__':
    read_file_input()
    target_elf1 = sorted(elf_list, key=lambda e: sum(e.get_items()))[-1]
    target_elf2 = sorted(elf_list, key=lambda e: sum(e.get_items()))[-2]
    target_elf3 = sorted(elf_list, key=lambda e: sum(e.get_items()))[-3]
    print(target_elf1.item_total() + target_elf2.item_total() + target_elf3.item_total())
