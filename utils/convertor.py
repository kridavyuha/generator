import csv
import json

# Read the CSV file
csv_file_path = 'ball_by_ball_ipl.csv'
json_file_path = 'ball_by_ball_ipl.json'

data = []
with open(csv_file_path, mode='r') as csv_file:
    csv_reader = csv.DictReader(csv_file)
    for row in csv_reader:
        data.append(row)

# Write the JSON file
with open(json_file_path, mode='w') as json_file:
    json.dump(data, json_file, indent=4)