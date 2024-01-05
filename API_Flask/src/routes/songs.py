from flask import Blueprint
from flask_login import login_required

import src.services.songs as songs_service


# from routes import users
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
      tags:
          - songs
    """
    return songs_service.get_songs()
