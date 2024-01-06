import json

import requests
from marshmallow import EXCLUDE

from src.schemas.song import SongSchema

songs_url = "http://localhost:8081/songs/"  # URL de l'API songs (golang)


def get_songs():
    response = requests.request(method="GET", url=songs_url)
    # TODO: Ajouter les ratings dans la réponse
    return response.json(), response.status_code

def create_song(new_song):
    song_schema = SongSchema().loads(json.dumps(new_song), unknown=EXCLUDE)

    response = requests.request(method="POST", url=songs_url, json=song_schema)
    if response.status_code != 201:
        return response.json(), response.status_code

    return response.json(), response.status_code

def get_song(id):
    response = requests.request(method="GET", url=songs_url+id)
    return response.json(), response.status_code

def delete_song(id):
    response = requests.request(method="DELETE", url=songs_url+id)
    if response.status_code != 204:
        return response.json(), response.status_code
    return response.json(), 204