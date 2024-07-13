from locust import HttpUser, task
import json

class FortifyUser(HttpUser):

    @task
    def ManualValueUser(self):
        payload = '{"name":"testname", "value":"testvalue"}'
        self.client.get('/')
        self.client.post('/save', data=payload)
        self.client.get('/show')

    @task
    def GeneratedValueUser(self):
        self.client.get('/')
        response = self.client.get('/generate')
        generated_value = response.json().get('message')
        payload = {"name":"testname",
                   "value": generated_value
        }
        self.client.post('/save', data=json.dumps(payload))
        self.client.get('/show')
