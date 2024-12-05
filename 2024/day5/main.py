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

# Process page lists
while pos < len(data):
	rulesOk = True
	pages = [int(p) for p in data[pos].split(",")]
	for (l, r) in rules:
		try:
			# Get the positions of the pages in the current rule...
			lpos = pages.index(l)
			rpos = pages.index(r)

			if lpos > rpos:
				# Rule violation - continue to the next page.
				rulesOk = False
				break

		except ValueError:
			# Dispose of ValueError here by continuing to the next rule
			# - if a rule contains a page number not contained in the set, ignore it.
			continue
	
	if rulesOk:
		total += pages[int(len(pages)/2)]
	
	pos += 1

print("total:", total)