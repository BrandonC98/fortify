import requests
import pytest

@pytest.mark.order(before='test_db_retrival')
def test_db_insert(fortify_service):
    payload = '{"name":"testname", "value":"testvalue"}'
    response = requests.post(fortify_service.base_url + fortify_service.save, data=payload)
    assert response.status_code == 200
    assert response.text == "successful"

def test_db_retrival(fortify_service):
    response = requests.get(fortify_service.base_url + fortify_service.show)
    assert response.status_code == 200
    assert "testname" in response.text
    assert "testvalue" in response.text

