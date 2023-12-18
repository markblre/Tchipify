"""OpenAPI v3 Specification"""
# Swagger documentation
# apispec via OpenAPI
from apispec import APISpec
from apispec.ext.marshmallow import MarshmallowPlugin
from apispec_webframeworks.flask import FlaskPlugin
from src.schemas.user import *
from src.schemas.user_auth import *
from src.schemas.errors import *


# Create an APISpec
spec = APISpec(
    title="Tchipify",
    version="1.0.0",
    openapi_version="3.0.2",
    plugins=[FlaskPlugin(), MarshmallowPlugin()],
)


# register used schemas with spec
spec.components.schema("User", schema=UserSchema)
spec.components.schema("UserLogin", schema=UserLoginSchema)
spec.components.schema("UserRegister", schema=UserRegisterSchema)
spec.components.schema("UserUpdate", schema=UserUpdateSchema)
spec.components.schema("Unauthorized", schema=UnauthorizedSchema)
spec.components.schema("Forbidden", schema=ForbiddenSchema)
spec.components.schema("NotFound", schema=NotFoundSchema)
spec.components.schema("Conflict", schema=ConflictSchema)
spec.components.schema("UnprocessableEntity", schema=UnprocessableEntitySchema)
spec.components.schema("SomethingWentWrong", schema=SomethingWentWrongSchema)


# add swagger tags that are used for endpoint annotation
tags = [
    {
        "name": "users",
        "description": "Managing users"
    },
    {
        "name": "auth",
        "description": "Managing authentication"
    }
]

for tag in tags:
    print(f"Adding tag: {tag['name']}")
    spec.tag(tag)
