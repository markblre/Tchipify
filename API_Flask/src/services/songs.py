import json

import requests
from marshmallow import EXCLUDE

from src.models.http_exceptions import UnprocessableEntity
from src.schemas.song import SongSchema

songs_url = "http://localhost:8081/songs/"  # URL de l'API songs (golang)


def song_exists(song_id):
    response = requests.request(method="GET", url=songs_url+song_id)
    if response.status_code == 200:
        return True
    elif response.status_code == 422:
        raise UnprocessableEntity
    return False

def get_songs():
    from src.services.ratings import get_ratings
    response = requests.request(method="GET", url=songs_url)
    if response.status_code != 200:
        return response.json(), response.status_code

    songs_json_with_ratings = response.json()

    for song in songs_json_with_ratings:
        song["ratings"], _ = get_ratings(song["id"])

    return songs_json_with_ratings, response.status_code

def create_song(new_song):
    song_schema = SongSchema().loads(json.dumps(new_song), unknown=EXCLUDE)

    response = requests.request(method="POST", url=songs_url, json=song_schema)
    if response.status_code != 201:
        return response.json(), response.status_code

    song_json_with_ratings = response.json()
    song_json_with_ratings["ratings"], _ = ([], 200)

    return song_json_with_ratings, response.status_code

def get_song(id):
    from src.services.ratings import get_ratings

    response = requests.request(method="GET", url=songs_url+id)
    if response.status_code != 200:
        return response.json(), response.status_code

    song_json_with_ratings = response.json()
    song_json_with_ratings["ratings"], _ = get_ratings(id)

    return song_json_with_ratings, response.status_code

def delete_song(id):
    response = requests.request(method="DELETE", url=songs_url+id)
    if response.status_code != 204:
        return response.json(), response.status_code
    return "", 204

def modify_song(id, song_update):
    from src.services.ratings import get_ratings
    song_schema = SongSchema().loads(json.dumps(song_update), unknown=EXCLUDE)

    response = requests.request(method="PUT", url=songs_url+id, json=song_schema)
    if response.status_code != 200:
        return response.json(), response.status_code

    song_json_with_ratings = response.json()
    song_json_with_ratings["ratings"], _ = get_ratings(id)

    return song_json_with_ratings, response.status_code