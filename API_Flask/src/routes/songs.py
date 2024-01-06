import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from src.schemas.errors import *
import src.services.songs as songs_service
from src.schemas.song import NewSongSchema, SongUpdateSchema


songs = Blueprint(name="songs", import_name=__name__)


@songs.route('/', methods=['GET'])
@login_required
def get_songs():
    """
    ---
    get:
      description: Getting songs
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - songs
    """
    try:
        return songs_service.get_songs()
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")

@songs.route('/', methods=['POST'])
@login_required
def post_song():
    """
    ---
    post:
      description: Posting a song
      requestBody:
        required: true
        content:
            application/json:
                schema: NewSong
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
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
    """

    try:
        new_song = NewSongSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    try:
        return songs_service.create_song(new_song)
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")

@songs.route('/<id>', methods=['GET'])
@login_required
def get_song(id):
    """
    ---
    get:
      description: Getting a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
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
    """
    try:
        return songs_service.get_song(id)
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")

@songs.route('/<id>', methods=['DELETE'])
@login_required
def delete_song(id):
    """
    ---
    delete:
      description: Delete a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      responses:
        '204':
          description: No content
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
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
    """
    try:
        return songs_service.delete_song(id)
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")

@songs.route('/<id>', methods=['PUT'])
@login_required
def put_song(id):
    """
    ---
    put:
      description: Updating a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      requestBody:
        required: true
        content:
            application/json:
                schema: SongUpdate
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
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
    """
    try:
        song_update = SongUpdateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    try:
        return songs_service.modify_song(id, song_update)
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")