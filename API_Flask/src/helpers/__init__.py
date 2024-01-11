from flask_sqlalchemy import SQLAlchemy
from flask import Flask


# db and app are here to avoid circular dependencies
db = SQLAlchemy()
app = Flask(__name__)
