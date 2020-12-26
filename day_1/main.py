
def main():
    f = open("input.txt", "r")
    lines = []
    for single_line in f:
        lines.append(int(single_line))

    for i in range(len(lines)):
        for j in range(i+1, len(lines)):
            for k in range(j+1, len(lines)):
                if lines[i] + lines[j] + lines[k] == 2020:
                    print(lines[i] * lines[j] * lines[k])
                    quit()

if __name__ == "__main__":
    main()
