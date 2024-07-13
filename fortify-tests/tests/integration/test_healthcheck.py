import requests
import pytest

@pytest.mark.order(1)
class TestHealthcheck:

    @pytest.mark.it
    @pytest.mark.healthcheck
    def test_generator_ping(self, generator_service):
        response = requests.get(generator_service.base_url + generator_service.ping)
        assert response.status_code == 200
        assert response.text == '{"message":"pong"}'

    @pytest.mark.it
    @pytest.mark.healthcheck
    def test_fortify_ping(self, fortify_service):
        response = requests.get(fortify_service.base_url + fortify_service.ping)
        assert response.status_code == 200
        assert response.text == '{"message":"pong"}'
