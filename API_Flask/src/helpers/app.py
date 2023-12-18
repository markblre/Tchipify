import os

from flask_login import LoginManager

from src.helpers import db, app
from src.models.user import User
from src.schemas.errors import UnauthorizedSchema


def config_app():
    # db localisation et nom
    uri = 'sqlite:///./users.db'
    app.app_context().push()

    # os.urandom permet de générer un nouveau secret de session (notamment authentification)
    # si vous souhaitez gérer une seule session pour vos tests, remplacez par "secret"
    app.config['SECRET_KEY'] = os.urandom(12)
    app.config['SQLALCHEMY_DATABASE_URI'] = uri

    db.init_app(app)

    with app.app_context():
        db.create_all()
        db.session.commit()

    login_manager = LoginManager()
    login_manager.init_app(app)

    # Vous pouvez commenter ce callback si vous ne voulez pas de body à vos réponses Unauthorized
    def unauthorized_response():
        error = UnauthorizedSchema().loads("{}")
        return error, error.get("code")
    login_manager.unauthorized_callback = unauthorized_response

    @login_manager.user_loader
    def load_user(user_id):
        return User.query.get(user_id)

    return app
