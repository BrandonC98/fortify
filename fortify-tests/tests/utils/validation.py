import json

def validate_json(json_string, expected_keys):
    try:
        data =json.loads(json_string)
    except json.JSONDecodeError:
        return False

    if not isinstance(data, dict):
        return False

    for key, expected_type in expected_keys.items():
        if key not in data or not isinstance(data[key], expected_type):
            return False

    return True
