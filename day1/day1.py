with open("input.txt") as file: 
    result = 0;
    instructions = file.read().split('\n')
    for step in instructions:
        result = result + int(step)
print(f'Result:{result}')