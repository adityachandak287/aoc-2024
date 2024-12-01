import sys
import re

space_pattern = re.compile("\s+")

l1 = []
l2 = []

for line in sys.stdin:
    line = line.strip()

    a, b = space_pattern.split(line)
    l1.append(int(a))
    l2.append(int(b))

l1.sort()
l2.sort()

answer = sum([abs(a - b) for a, b in zip(l1, l2)])

print(answer)
