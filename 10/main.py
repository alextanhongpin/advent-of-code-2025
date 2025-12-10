import numpy as np
from scipy.optimize import milp, LinearConstraint
import re

with open("10/input.txt") as f:
    total = 0
    for line in f.readlines():
        line = re.sub(r"[\[\]\{\}\(\)]", "", line)
        _, *raw_buttons, joltage = line.split(" ")
        joltage = list(map(int, joltage.split(",")))
        buttons = []
        for b in raw_buttons:
            bi = [0] * len(joltage)
            for i in map(int, b.split(",")):
                bi[i] = 1
            buttons.append(bi)
        A = np.array(buttons)
        b = np.array(joltage)
        c = np.array([1 for _ in A])
        A = A.T
        constraints = LinearConstraint(A, b, b)
        integrality = np.ones_like(c)  # array of 1's with same shape as c
        res = milp(c=c, constraints=constraints, integrality=integrality)
        total += res.fun
    print(int(total))  # 18559
