
import unittest

from deepdiff import DeepDiff
from fastapi.testclient import TestClient
from main import app

client = TestClient(app)



class TestMainApi(unittest.TestCase):
    def setUp(self):
        pass

    def test_rootHello(self):
        headers = {'Authorization' : 'Bearer CI_TEST'}
        response = client.get("/", headers=headers)
        assert response.status_code == 200
        assert response.json() == {"response": "Hello World"}

    def test_rootPost(self):
        headers = {'Authorization' : 'Bearer CI_TEST'}
        response = client.get("/",headers=headers)
        response = client.post(
            "/", headers={"X-Token": "coneofsilence", 'Authorization' : 'Bearer CI_TEST'}, json={"id": "1", "name": "test"})
        assert response.status_code == 200
        diff = DeepDiff(response.json(), {
                        "id": "1", "name": "test"}, ignore_order=True)
        assert diff, f"difference in response: {diff}"


if __name__ == "__main__":
    unittest.main()
