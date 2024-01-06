from flask import Blueprint
from flask_login import login_required

from src.models.http_exceptions import NotFound, UnprocessableEntity
from src.schemas.errors import *
import src.services.ratings as ratings_service


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


