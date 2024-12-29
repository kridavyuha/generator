import json

# Read the JSON file
with open('ball_by_ball_ipl.json', 'r') as file:
    data = json.load(file)

# Replace team names
for ball in data:
    if ball['batFirst'] == 'Chennai Super Kings':
        ball['batFirst'] = 'CSK'
    if ball['batSecond'] == 'Kolkata Knight Riders':
        ball['batSecond'] = 'KKR'
    if ball['winner'] == 'Chennai Super Kings':
        ball['winner'] = 'CSK'
    if ball['winner'] == 'Kolkata Knight Riders':
        ball['winner'] = 'KKR'

# Write back to file
with open('ball_by_ball_ipl.json', 'w') as file:
    json.dump(data, file, indent=2)