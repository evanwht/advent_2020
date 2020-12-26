
def main():
    f = open("input.txt", "r")
    lines = []
    for single_line in f:
        lines.append(single_line)

    valid = 0
    for line in lines:
        parts = line.split(": ")
        if is_password_valid(parts[0], parts[1]):
            valid += 1
    print(valid)

def is_password_valid(rule, password):
    amount, letter = rule.split(" ")
    low, high = amount.split("-")
    low = int(low) - 1
    high = int(high) - 1 

    if password[low] == letter or password[high] == letter:
        if password[low] != password[high]:
            return True
    else:
        return False

if __name__ == "__main__":
    main()
