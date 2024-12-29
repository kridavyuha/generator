import json

# Read the JSON file
with open('ball_by_ball_ipl_camel.json', 'r') as file:
    data = json.load(file)

# Increment ball numbers
for ball in data:
    ball['ballNo'] = int(ball['ballNo']) + 1

# Write back to file
with open('ball_by_ball_ipl.json', 'w') as file:
    json.dump(data, file, indent=2)