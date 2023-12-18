
"""erreur http"""
class UnprocessableEntity(Exception):
    pass


class NotFound(Exception):
    pass


class Conflict(Exception):
    pass


class Unauthorized(Exception):
    pass


class Forbidden(Exception):
    pass


class SomethingWentWrong(Exception):
    pass
