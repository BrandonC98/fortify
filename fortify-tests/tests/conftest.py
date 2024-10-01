import dotenv
import pytest
from  dotenv import dotenv_values
import json
import os

def env_config(val):
    return 'env/' + val

class generator:
    base_url = ""
    ping = "/ping"
    def __init__(self):
        env = env_config(os.getenv('CONFIG'))
        config = dotenv_values(env)
        self.base_url = config["GENERATOR_BASE_URL"]

class fortify:
    base_url = ""
    ping = "/ping"
    generate = "/generate"
    save = "/save"
    show = "/show"
    nameFieldId = "#nameField"
    valueFieldId = "#valueField"
    secretsFieldId = "#secretList"
    saveBtn = "#saveBtn"
    showBtn = "#showBtn"
    generateBtn = "#generateBtn"
    def __init__(self):
        env = env_config(os.getenv('CONFIG'))
        config = dotenv_values(env)
        self.base_url = config["FORTIFY_BASE_URL"]

@pytest.fixture
def fortify_service():
    return fortify()

@pytest.fixture
def generator_service():
    return generator()
