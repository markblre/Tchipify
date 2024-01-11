from marshmallow import Schema, fields, post_load


# Schemas marschmallow pour les erreurs HTTP
class BaseHTTPError(Schema):
    message = fields.String(description="Error message")

    # permet après avoir loadé le contenu, de mettre comme message par défaut la fonction __str__ du schéma
    @post_load
    def load_default_message(self, in_data, **kwargs):
        if "message" not in in_data:
            in_data["message"] = self.__str__()
        return in_data


class UnprocessableEntitySchema(BaseHTTPError):
    code = fields.Constant(422, description="HTTP code")

    def __str__(self):
        return "Unprocessable Entity"


class NotFoundSchema(BaseHTTPError):
    code = fields.Constant(404, description="HTTP code")

    def __str__(self):
        return "Not Found"


class ConflictSchema(BaseHTTPError):
    code = fields.Constant(409, description="HTTP code")

    def __str__(self):
        return "Conflict"


class UnauthorizedSchema(BaseHTTPError):
    code = fields.Constant(401, description="HTTP code")

    def __str__(self):
        return "Unauthorized"


class ForbiddenSchema(BaseHTTPError):
    code = fields.Constant(403, description="HTTP code")

    def __str__(self):
        return "Forbidden"


class SomethingWentWrongSchema(BaseHTTPError):
    code = fields.Constant(500, description="HTTP code")

    def __str__(self):
        return "Something went wrong"
