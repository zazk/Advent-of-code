from functools import reduce
with open("input.txt") as file: 
    instructions = file.read().split('\n')
    res = reduce(lambda res, value: res + int(value), instructions, 0)
print(f'Result: {res}')