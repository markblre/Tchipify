import json

import requests
from flask_login import current_user
from marshmallow import EXCLUDE

from src.models.http_exceptions import *
from src.schemas.ratings import RatingSchema
from src.services.songs import song_exists


# URL de l'API ratings (golang)
def get_ratings_url(song_id):
    return "http://localhost:8082/songs/"+song_id+"/ratings/"


def get_ratings(song_id):
    if not song_exists(song_id):
        raise NotFound
    response = requests.request(method="GET", url=get_ratings_url(song_id))
    return response.json(), response.status_code


def create_rating(new_rating, song_id):
    if not song_exists(song_id):
        raise NotFound

    rating_schema = RatingSchema().loads(json.dumps(new_rating), unknown=EXCLUDE)

    rating_schema["user_id"] = current_user.id

    response = requests.request(method="POST", url=get_ratings_url(song_id), json=rating_schema)
    return response.json(), response.status_code

def get_rating(song_id, rating_id):
    if not song_exists(song_id):
        raise NotFound
    response = requests.request(method="GET", url=get_ratings_url(song_id)+rating_id)
    return response.json(), response.status_code