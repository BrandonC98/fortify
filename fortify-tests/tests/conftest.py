import pytest
from  dotenv import dotenv_values
import json

class generator:
    base_url = ""
    ping = "/ping"
    def __init__(self):
        config = dotenv_values('configuration/environments/test-env')
        self.base_url = config["GENERATOR_BASE_URL"]

class fortify:
    base_url = ""
    ping = "/ping"
    generate = "/generate"
    save = "/save"
    show = "/show"
    def __init__(self):
        config = dotenv_values('configuration/environments/test-env')
        self.base_url = config["FORTIFY_BASE_URL"]

@pytest.fixture
def fortify_service():
    return fortify()

@pytest.fixture
def generator_service():
    return generator()
