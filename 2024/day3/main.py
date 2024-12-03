import re

# It's another python kind of day.

def main():
	total = 0
	newtotal = 0
	enable = True
	file = open("input.txt", "r")
	data = file.read()

	# and then I had two problems.
	mul = re.compile(r'(?P<func>do(?:n\'t)?)\(\)|mul\((?P<l>\d+),(?P<r>\d+)\)')

	for res in re.finditer(mul, data):
		match res.group("func"):
			case "do":
				enable = True
				continue
			case "don't":
				enable = False
				continue
		
		if enable:
			newtotal += int(res.group("l")) * int(res.group("r"))
		
		total += int(res.group("l")) * int(res.group("r"))
	
	print("total(1):", total)
	print("total(2):", newtotal)

if __name__ == "__main__":
	main()