import itertools as it


f = lambda l: [(min(a, c), min(b, d), max(a, c), max(b, d)) for (a, b), (c, d) in l]

red = list(map(eval, open("09/input.txt")))
green = f(it.pairwise(red))

a = b = 0

for x, y, u, v in f(it.combinations(red, 2)):
    size = (u - x + 1) * (v - y + 1)

    if size > b:
        for p, q, r, s in green:
            if x < r and y < s and u > p and v > q:
                break

        else:
            b = max(b, size)

    a = max(a, size)

print(a, b)
