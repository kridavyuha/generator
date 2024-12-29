import json

def build_squad_lookup(squads_path: str) -> dict:
    """Create a lookup of {teamName: {playerName: playerID}} from squads.json."""
    with open(squads_path, 'r') as f:
        squads = json.load(f)

    squad_lookup = {}
    for squad in squads:
        team = squad["team"]
        if team not in squad_lookup:
            squad_lookup[team] = {}
        for p in squad["players"]:
            squad_lookup[team][p["name"]] = p["id"]
    return squad_lookup

def add_player_ids(ball_data_path: str, squad_lookup: dict, output_path: str) -> None:
    """Append new ID fields without removing original name fields."""
    with open(ball_data_path, 'r') as f:
        balls = json.load(f)

    updated_balls = []
    for ball in balls:
        # Determine batting/bowling teams based on innings
        if ball["innings"] == "1":
            batting_team = ball["batFirst"]
            bowling_team = ball["batSecond"]
        else:
            batting_team = ball["batSecond"]
            bowling_team = ball["batFirst"]

        # Original fields (names)
        batter_name = ball.get("batter", "")
        non_striker_name = ball.get("nonStriker", "")
        bowler_name = ball.get("bowler", "")
        caught_by = ball.get("caughtBy", "")
        player_out = ball.get("playerOut", "")

        # New ID fields:
        # Batter, non-striker, playerOut come from batting team
        # Bowler, caughtBy come from bowling team
        ball["batterId"] = squad_lookup.get(batting_team, {}).get(batter_name, "")
        ball["nonStrikerId"] = squad_lookup.get(batting_team, {}).get(non_striker_name, "")
        ball["bowlerId"] = squad_lookup.get(bowling_team, {}).get(bowler_name, "")
        ball["caughtById"] = squad_lookup.get(bowling_team, {}).get(caught_by, "")
        ball["playerOutId"] = squad_lookup.get(batting_team, {}).get(player_out, "")

        updated_balls.append(ball)

    # Write updated data to new file
    with open(output_path, 'w') as f:
        json.dump(updated_balls, f, indent=2)

if __name__ == "__main__":
    squads_path = "squads.json"
    ball_data_path = "ball_by_ball_ipl.json"
    output_path = "ball_by_ball_ipl_with_ids.json"

    squad_lookup = build_squad_lookup(squads_path)
    add_player_ids(ball_data_path, squad_lookup, output_path)
    print(f"Updated ball-by-ball data saved to {output_path}")