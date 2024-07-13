import requests
from utils.validation import validate_json
import pytest

@pytest.mark.it
def test_generation( fortify_service):
    expected_Keys = {
        "message": str
    }
    response = requests.get(fortify_service.base_url + fortify_service.generate)
    assert response.status_code == 200
    assert validate_json(response.text, expected_Keys) == True
