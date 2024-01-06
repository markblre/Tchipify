from marshmallow import Schema, fields, validates_schema, ValidationError


class SongSchema(Schema):
    artist = fields.String(description="Artist of the song")
    file_name = fields.String(description="Song file name")
    id = fields.String(description="UUID")
    published_date = fields.DateTime(description="Published date")
    ratings = fields.List(fields.String, description="Ratings of the song")  # TODO: renvoyer un tableau avec les notations en JSON
    title = fields.String(description="Title")

    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
                (not obj.get("artist") or obj.get("artist") == "") and \
                (not obj.get("file_name") or obj.get("file_name") == "") and \
                (not obj.get("published_date") or obj.get("published_date") == "") and \
                (not obj.get("ratings") or obj.get("ratings") == "") and \
                (not obj.get("title") or obj.get("title") == "")

class BaseSongSchema(Schema):
    artist = fields.String(description="Artist of the song")
    file_name = fields.String(description="Song file name")
    title = fields.String(description="Title")

class NewSongSchema(BaseSongSchema):
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if "artist" not in data or data["artist"] == "" or \
                "file_name" not in data or data["file_name"] == "" or \
                "title" not in data or data["title"] == "":
            raise ValidationError("['artist','file_name','title'] must all be specified")

class SongUpdateSchema(BaseSongSchema):
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("artist" in data and data["artist"] != "") or
                ("file_name" in data and data["file_name"] != "") or
                ("title" in data and data["title"] != "")):
            raise ValidationError("at least one of ['artist','file_name','title'] must be specified")

