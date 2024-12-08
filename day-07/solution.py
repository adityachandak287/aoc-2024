import argparse
import sys
from dataclasses import dataclass
import enum
from itertools import product

parser = argparse.ArgumentParser(
    prog="solution",
    description="AOC 2024",
)

parser.add_argument("--part", choices=["A", "B"], required=True)

args = parser.parse_args()


@dataclass
class Equation:
    total: int
    numbers: list[int]


class Op(enum.IntEnum):
    ADD = 1
    MULTIPLY = 2
    CONCAT = 3


if args.part == "A":
    OPERATORS = [Op.ADD, Op.MULTIPLY]
else:
    OPERATORS = [Op.ADD, Op.MULTIPLY, Op.CONCAT]


equations: list[Equation] = []


def check_equation(eq: Equation) -> bool:
    num_operators = len(eq.numbers) - 1

    for perm in product(OPERATORS, repeat=num_operators):
        total = eq.numbers[0]
        for idx, num in enumerate(eq.numbers[1:]):
            match perm[idx]:
                case Op.ADD:
                    total += num
                case Op.MULTIPLY:
                    total *= num
                case Op.CONCAT:
                    total = int(str(total) + str(num))
                case _:
                    raise ValueError("Invalid operator!")

        if total == eq.total:
            print(total, [op.name for op in perm])
            return True

    return False


answer = 0
for line in sys.stdin:
    line = line.strip()

    colon_idx = line.index(":")

    eq = Equation(
        total=int(line[:colon_idx]),
        numbers=[int(n) for n in line[colon_idx + 1 :].strip().split(" ")],
    )
    print(eq)
    if check_equation(eq):
        answer += eq.total

print("Answer", answer)
