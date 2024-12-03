import re

# It's another python kind of day.

def main():
	total = 0
	file = open("input.txt", "r")
	data = file.read()

	mul = re.compile(r'mul\((\d+),(\d+)\)')

	for res in [int(x.group(1)) * int(x.group(2)) for x in re.finditer(mul, data)]:
		total += res
	
	print("total:", total)

if __name__ == "__main__":
	main()