import sys
import re

fn_pattern = re.compile(r"(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don't\(\))")
args_pattern = re.compile(r"\d{1,3}")

ans = 0
enabled = True
for line in sys.stdin:
    line = line.strip()

    matches = fn_pattern.findall(line)
    functions = [fn for grp in matches for fn in grp if fn != ""]

    assert len(matches) == len(functions)

    for fn in functions:
        if fn in ["do()", "don't()"]:
            enabled = fn == "do()"
            continue

        if not enabled:
            continue

        args = [int(arg) for arg in args_pattern.findall(fn)]
        assert len(args) == 2

        ans += args[0] * args[1]

print("Answer", ans)
