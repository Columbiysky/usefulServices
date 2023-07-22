
from deepdiff import DeepDiff
from fastapi.testclient import TestClient

from .main import app

client = TestClient(app)

def test_rootHello():
    response = client.get("/")
    assert response.status_code == 200
    assert response.json() == {"response":"Hello World"}

def test_rootPost():
    response = client.post("/",headers={"X-Token": "coneofsilence"}, json={"id":"1","name":"test"})
    assert response.status_code == 200
    diff = DeepDiff(response.json(), {"id":"1","name":"test"}, ignore_order=True)
    assert diff, f"difference in response: {diff}"