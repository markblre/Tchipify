import json

from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from src.models.http_exceptions import NotFound, UnprocessableEntity
from src.schemas.errors import *
import src.services.ratings as ratings_service
from src.schemas.ratings import NewRatingSchema

ratings = Blueprint(name="ratings", import_name=__name__)


@ratings.route('/', methods=['GET'])
@login_required
def get_ratings(song_id):
    """
    ---
    get:
      description: Getting ratings of a song
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Song not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - songs
          - ratings
    """
    try:
        return ratings_service.get_ratings(song_id)
    except NotFound:
        error = NotFoundSchema().loads("{}")
        return error, error.get("code")
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads("{}")
        return error, error.get("code")
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")

@ratings.route('/', methods=['POST'])
@login_required
def post_rating(song_id):
    """
    ---
    post:
      description: Posting a rating
      requestBody:
        required: true
        content:
            application/json:
                schema: NewRating
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '422':
          description: Rating must be between 0 and 5
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - songs
          - ratings
    """

    try:
        new_rating = NewRatingSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    try:
        return ratings_service.create_rating(new_rating, song_id)
    except NotFound:
        error = NotFoundSchema().loads("{}")
        return error, error.get("code")
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads("{}")
        return error, error.get("code")
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")

@ratings.route('/<rating_id>', methods=['GET'])
@login_required
def get_song(song_id, rating_id):
    """
    ---
    get:
      description: Getting a rating of a song
      parameters:
        - in: path
          name: song_id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
        - in: path
          name: rating_id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - songs
          - ratings
    """
    try:
        return ratings_service.get_rating(song_id, rating_id)
    except NotFound:
        error = NotFoundSchema().loads("{}")
        return error, error.get("code")
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads("{}")
        return error, error.get("code")
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")