from werkzeug.security import check_password_hash

from src.models.http_exceptions import *
import src.services.users as users_service
"""communique avec le repository ou les apis go """

def login(user_login):
    # on récupère l'utilisateur depuis la base de donnée
    existing_user = users_service.get_user_from_db(user_login.get("username"))
    if existing_user:
        # on vérifie son mot de passe
        if not check_password_hash(existing_user.encrypted_password, user_login.get("password")):
            raise Unauthorized
    else:
        raise NotFound

    return existing_user


def register(user_register):
    # on vérifie que l'utilisateur n'existe pas déjà
    if users_service.user_exists(user_register.get("username")):
        raise Conflict

    # on crée l'utilisateur
    try:
        return users_service.create_user(user_register)
    except Exception:
        raise SomethingWentWrong
