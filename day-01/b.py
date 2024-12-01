import sys
import re
from collections import Counter

space_pattern = re.compile("\s+")

l1 = []
l2 = []

for line in sys.stdin:
    line = line.strip()

    a, b = space_pattern.split(line)
    l1.append(int(a))
    l2.append(int(b))

l2_counts = Counter(l2)

answer = sum([a * l2_counts[a] for a in l1])

print(answer)
