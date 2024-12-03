import sys
import re

fn_pattern = re.compile(r"mul\(\d{1,3},\d{1,3}\)")
args_pattern = re.compile(r"\d{1,3}")

ans = 0
for line in sys.stdin:
    line = line.strip()

    for fn in fn_pattern.findall(line):
        args = [int(arg) for arg in args_pattern.findall(fn)]
        assert len(args) == 2

        ans += args[0] * args[1]

print("Answer", ans)
