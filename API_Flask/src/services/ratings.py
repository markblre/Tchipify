import requests

from src.models.http_exceptions import *
from src.services.songs import song_exists


# URL de l'API ratings (golang)
def get_ratings_url(song_id):
    return "http://localhost:8082/songs/"+song_id+"/ratings/"


def get_ratings(song_id):
    if not song_exists(song_id):
        raise NotFound
    response = requests.request(method="GET", url=get_ratings_url(song_id))
    return response.json(), response.status_code
