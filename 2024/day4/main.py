# I swear I'll switch away from python again... eventually.
# It's just easy for challenges like this.

def main():
	total = 0
	file = open("input.txt", "r")
	data = file.readlines()

	for x in range(len(data)):
		for y in range(len(data[0])):
			if data[x][y] == "X":
				total += findXmas(data, x, y)

	print("total:", total)

# Check Adjacent Cells for XMAS
def findXmas(arr, x, y) -> int:
	found = 0
	if arr[x][y] != "X":
		return 0
	for i in range(-1, 2):
		for j in range(-1, 2):
			# print("%d,%d @ %d, %d" % (x, y, i, j))
			# Don't check the current cell
			if i == 0 and j == 0:
				continue
			# Edge cases (...get it?)
			if x+(3*i) not in range(len(arr)) or y+(3*j) not in range(len(arr[0])):
				continue
			if arr[x+i][y+j] == "M" and arr[x+(2*i)][y+(2*j)] == "A" and arr[x+(3*i)][y+(3*j)] == "S":
				# print("Found XMAS: %d,%d to %d,%d (%d, %d)" % (x, y, x+(3*i), y+(3*j), i, j))
				found += 1
	return found


if __name__ == "__main__":
	main()