from flask import jsonify

# load modules
from src.routes.users import users as users_route
from src.routes.auth import auth as auth_route
from src.routes.swagger import swagger_ui_blueprint, SWAGGER_URL
from src.api_spec import spec
from src.helpers.app import config_app

# configure application and DB
app = config_app()

# register routes
app.register_blueprint(auth_route, url_prefix="/")
app.register_blueprint(users_route, url_prefix="/users")

# allows to generate Swagger doc for all documented functions
with app.test_request_context():
    for fn_name in app.view_functions:
        if fn_name == 'static':
            continue
        print(f"Loading swagger docs for function: {fn_name}")
        view_fn = app.view_functions[fn_name]
        spec.path(view=view_fn)


# specify where to get the generated doc
@app.route("/api/swagger.json")
def create_swagger_spec():
    """
    Swagger API definition.
    """
    return jsonify(spec.to_dict())


# register documentation route (see in browser at /api/docs)
app.register_blueprint(swagger_ui_blueprint, url_prefix=SWAGGER_URL)


# python main entrance program
if __name__ == "__main__":
    app.run(host='0.0.0.0', port=8888, debug=False)
