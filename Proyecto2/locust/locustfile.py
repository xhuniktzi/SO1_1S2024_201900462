from locust import HttpUser, task, between
from json import load
from random import choice

class VotingMusic(HttpUser):

    wait_time = between(1, 5)

    def on_start(self):
        with open("music.json", "r", encoding="utf-8") as f:
            self.music = load(f)

    @task
    def vote(self):
        music = choice(self.music)
        self.client.post("/insert", json=music, headers={"Content-Type": "application/json"})