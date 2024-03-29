import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from src.helpers.content_negotiation import content_negotiation
from src.models.http_exceptions import *
from src.schemas.user import UserUpdateSchema
from src.schemas.errors import *
import src.services.users as users_service

# from routes import users
users = Blueprint(name="users", import_name=__name__)


@users.route('/<id>', methods=['GET'])
@login_required
def get_user(id):
    """
    ---
    get:
      description: Getting a user
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of user id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
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
          - users
    """
    return content_negotiation(*users_service.get_user(id))


@users.route('/<id>', methods=['PUT'])
@login_required
def put_user(id):
    """
    ---
    put:
      description: Updating a user
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of user id
      requestBody:
        required: true
        content:
            application/json:
                schema: UserUpdate
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
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
        '409':
          description: Conflict
          content:
            application/json:
              schema: Conflict
            application/yaml:
              schema: Conflict
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
          - users
    """
    # parser le body
    try:
        user_update = UserUpdateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return content_negotiation(error, error.get("code"))

    # modification de l'utilisateur (username, nom, mot de passe, etc.)
    try:
        return content_negotiation(*users_service.modify_user(id, user_update))
    except Conflict:
        error = ConflictSchema().loads(json.dumps({"message": "User already exists"}))
        return content_negotiation(error, error.get("code"))
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
        return content_negotiation(error, error.get("code"))
    except Forbidden:
        error = ForbiddenSchema().loads(json.dumps({"message": "Can't manage other users"}))
        return content_negotiation(error, error.get("code"))
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return content_negotiation(error, error.get("code"))


@users.route('/<id>', methods=['DELETE'])
@login_required
def delete_user(id):
    """
    ---
    delete:
      description: Delete a user
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of user id
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
        '403':
          description: Forbidden
          content:
            application/json:
              schema: Forbidden
            application/yaml:
              schema: Forbidden
        '500':
          description: Something went wrong
          content:
            application/json:
              schema: SomethingWentWrong
            application/yaml:
              schema: SomethingWentWrong
      tags:
          - users
    """
    try:
        return content_negotiation(*users_service.delete_user(id))
    except Forbidden:
        error = ForbiddenSchema().loads(json.dumps({"message": "Can't delete other users"}))
        return content_negotiation(error, error.get("code"))
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return content_negotiation(error, error.get("code"))


@users.route('/', methods=['GET'])
def get_users():
    """
    ---
    get:
      description: Getting users
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: User
            application/yaml:
              schema: User
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
          - users
    """
    return content_negotiation(*users_service.get_users())
