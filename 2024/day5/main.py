# no main function this time since it's kinda redundant

file = open("input.txt", "r")
data = [l.strip() for l in file.readlines()]
file.close()

pos = 0
rules = []
total = 0

# Parse Rules
while data[pos] != "":
	[l, r] = [int(i) for i in data[pos].split("|")][:2]
	rules.append((l,r))
	pos += 1

print("rules:", len(rules))
pos += 1

# refactor: pulled this logic out into a function for reuse
def validate(pages, rules):
	for (l, r) in rules:
		try:
			# Get the positions of the pages in the current rule...
			lpos = pages.index(l)
			rpos = pages.index(r)

			if lpos > rpos:
				return False, (l, r), (lpos, rpos)

		except ValueError:
			# Dispose of ValueError here by continuing to the next rule
			# - if a rule contains a page number not contained in the set, ignore it.
			continue
	return True, None, None

violating: list[list[int]] = []

# Process page lists
while pos < len(data):
	pages = [int(p) for p in data[pos].split(",")]
	
	ok, _, _ = validate(pages, rules)
	if ok:
		total += pages[int(len(pages)/2)]
	else:
		violating.append(pages)
	
	pos += 1

print("total(1):", total)

newTotal = 0

# ALL OF THE BRUTE FORCE.
for pages in violating:
	ok = False
	while not ok:
		ok, rule, pos = validate(pages, rules)
		if rule is not None:
			# print(pages, rule, pos)
			page = pages.pop(pos[1])
			pages.insert(pos[0], page)
	newTotal += pages[int(len(pages)/2)]

print("total(2):", newTotal)