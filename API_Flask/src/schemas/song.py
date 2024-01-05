from marshmallow import Schema, fields


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
