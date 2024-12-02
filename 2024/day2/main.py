# It's a python kind of day.

UP = 1
UNSET = 0
DOWN = -1

def main():
	validLines = 0
	forcedLines = 0
	file = open("input.txt", "r")
	lines = file.readlines()
	for line in lines:
		vals = [int(v) for v in line.split(" ")]
		if isValid(vals):
			validLines += 1
		else:
			# brute force. works a charm.
			for v in range(len(vals)):
				if isValid(vals[:v] + vals[v+1:]):
					forcedLines += 1
					break
	
	print("lines:", len(lines))
	print("valid(1):", validLines)
	print("valid(2):", validLines + forcedLines)


def isValid(vals):
	direction = UNSET
	last = vals[0]

	for n in range(1, len(vals)):
		if direction == UNSET:
			if vals[n] > last:
				direction = UP
			elif vals[n] < last:
				direction = DOWN

		diff = abs(vals[n] - last)
		if (diff > 0 and diff <= 3) and ((direction == UP and vals[n] > last) or (direction == DOWN and vals[n] < last)):
			last = vals[n]
		else:
			return False
	return True


if __name__ == "__main__":
	main()