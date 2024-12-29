import json

def to_camel_case(snake_str):
    components = snake_str.split(' ')
    return components[0].lower() + ''.join(x.title() for x in components[1:])

# Read the JSON file
with open('ball_by_ball_ipl.json', 'r') as file:
    data = json.load(file)

# Convert keys to camelCase
camel_case_data = []
for item in data:
    camel_case_item = {}
    for key, value in item.items():
        camel_case_key = to_camel_case(key)
        camel_case_item[camel_case_key] = value
    camel_case_data.append(camel_case_item)

# Write the converted JSON
with open('ball_by_ball_ipl_camel.json', 'w') as file:
    json.dump(camel_case_data, file, indent=2)