import json
from flask import Blueprint, request
from marshmallow import ValidationError
from flask_login import login_user, logout_user, login_required, current_user

from src.models.http_exceptions import *
from src.schemas.errors import *
from src.schemas.user_auth import UserLoginSchema, UserRegisterSchema
import src.services.users as users_service
import src.services.auth as auth_service


auth = Blueprint(name="login", import_name=__name__)


@auth.route('/login', methods=['POST'])
def login():
    """
    ---
    post:
      description: Login
      requestBody:
        required: true
        content:
            application/json:
                schema: UserLogin
      responses:
        '200':
          description: Ok
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '403':
          description: Already logged in
          content:
            application/json:
              schema: Forbidden
            application/yaml:
              schema: Forbidden
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - auth
          - users
    """
    if current_user.is_authenticated:
        error = ForbiddenSchema().loads(json.dumps({"message": "Already logged in"}))
        return error, error.get("code")

    # parser le body
    try:
        # it is possible to use marshmallow Schemas validation (used also for doc) or custom classes
        user_login = UserLoginSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    # logger l'utilisateur
    try:
        user = auth_service.login(user_login)
    except (NotFound, Unauthorized):
        error = UnauthorizedSchema().loads("{}")
        return error, error.get("code")
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")

    login_user(user, remember=True)
    return "", 200


@auth.route('/logout', methods=['POST'])
@login_required
def logout():
    """
    ---
    post:
      description: Logout
      responses:
        '200':
          description: Ok
        '401':
          description: Unauthorized
      tags:
          - auth
          - users
    """
    logout_user()
    return "", 200


@auth.route('/register', methods=['POST'])
def register():
    """
    ---
    post:
      description: Register
      requestBody:
        required: true
        content:
            application/json:
                schema: UserRegister
      responses:
        '201':
          description: Created
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
        '403':
          description: Already logged in
          content:
            application/json:
              schema: Forbidden
            application/yaml:
              schema: Forbidden
        '409':
          description: User already exists
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
          - auth
          - users
    """
    if current_user.is_authenticated:
        error = ForbiddenSchema().loads(json.dumps({"message": "Already logged in"}))
        return error, error.get("code")

    # parser le body
    try:
        user_register = UserRegisterSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    # enregistrer l'utilisateur
    try:
        return auth_service.register(user_register)
    except Conflict:
        error = ConflictSchema().loads(json.dumps({"message": "User already exists"}))
        return error, error.get("code")
    except SomethingWentWrong:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")


#pour obtenir l'utilisateur connecté
@auth.route('/introspect', methods=["GET"])
@login_required
def introspect():
    """
    ---
    get:
      description: Getting authenticated user
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
      tags:
          - auth
          - users
    """
    return users_service.get_user(current_user.id)
