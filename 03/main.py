def max_joltage(s: str) -> int:
    if len(s) < 12:
        return 0
    if len(s) == 12:
        return int(s)

    num = 0
    for i in range(min(13, len(s))):
        ns = s[:i] + s[i + 1 :]
        if int(ns) > num:
            num = int(ns)

    return max_joltage(str(num))


with open("input.txt") as f:
    total = 0
    txt = f.read()
    for line in txt.split("\n"):
        if line == "":
            continue
        total += max_joltage(line)
    print(total)
