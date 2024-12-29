import json

# Read the JSON file
with open('ball_by_ball_ipl.json', 'r') as file:
    data = json.load(file)

# Initialize team players
csk_players = set()
kkr_players = set()

# Collect unique players
for ball in data:
    if ball['batFirst'] == 'CSK':
        csk_players.add(ball['batter'])
        csk_players.add(ball['nonStriker'])
    else:
        kkr_players.add(ball['batter'])
        kkr_players.add(ball['nonStriker'])
    
    # Add bowlers to respective teams
    if ball['batFirst'] == 'CSK':
        kkr_players.add(ball['bowler'])
    else:
        csk_players.add(ball['bowler'])

# Create player mappings with short keys
team_players = {
    "CSK": {player: ''.join(word[0] for word in player.split()) for player in csk_players},
    "KKR": {player: ''.join(word[0] for word in player.split()) for player in kkr_players}
}

# Create formatted output
output = {
    "teams": {
        "CSK": {
            "fullName": "Chennai Super Kings",
            "players": {name: shortkey for name, shortkey in sorted(team_players["CSK"].items())}
        },
        "KKR": {
            "fullName": "Kolkata Knight Riders",
            "players": {name: shortkey for name, shortkey in sorted(team_players["KKR"].items())}
        }
    }
}

# Write to new JSON file
with open('teams_players.json', 'w') as file:
    json.dump(output, file, indent=2)