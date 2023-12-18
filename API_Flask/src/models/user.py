from flask_login import UserMixin
from werkzeug.security import generate_password_hash
from src.helpers import db


# modèle de donnée pour la base de donnée utilisateur
# vous pouvez lier les utilisateurs de cette API à ceux de la vôtre (Golang) avec leur ID ou leur username
class User(UserMixin, db.Model):
    __tablename__ = 'users'

    id = db.Column(db.String(255), primary_key=True)
    username = db.Column(db.String(255), unique=True, nullable=False)
    encrypted_password = db.Column(db.String(255), nullable=False)

    def __init__(self, uuid, username, encrypted_password):
        self.id = uuid
        self.username = username
        self.encrypted_password = encrypted_password

    def is_empty(self):
        return (not self.id or self.id == "") and \
               (not self.username or self.username == "") and \
               (not self.encrypted_password or self.encrypted_password == "")

    @staticmethod
    def from_dict_with_clear_password(obj):
        username = obj.get("username") if obj.get("username") != "" else None
        password = generate_password_hash(obj.get("password")) if obj.get("password") != "" else None
        return User(None, username, password)
