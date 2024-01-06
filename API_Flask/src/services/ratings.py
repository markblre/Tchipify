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

def delete_rating(song_id, rating_id):
    response = requests.request(method="DELETE", url=get_ratings_url(song_id)+rating_id)
    if response.status_code != 204:
        return response.json(), response.status_code
    return "", 204

def modify_rating(song_id, rating_id, rating_update):
    if not song_exists(song_id):
        raise NotFound

    r_json, r_code = get_rating(song_id, rating_id)
    if r_code != 200:
        return r_json, r_code

    if r_json["user_id"] != current_user.id:
        raise Forbidden

    rating_schema = RatingSchema().loads(json.dumps(rating_update), unknown=EXCLUDE)

    response = requests.request(method="PUT", url=get_ratings_url(song_id)+rating_id, json=rating_schema)

    return response.json(), response.status_code