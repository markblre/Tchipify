from marshmallow import Schema, fields


class RatingSchema(Schema):
    comment = fields.String(description="Comment of the rating")
    id = fields.String(description="UUID")
    rating = fields.Integer(description="Rating")
    rating_date = fields.DateTime(description="Rating date")
    song_id = fields.String(description="UUID")
    user_id = fields.String(description="UUID")

    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
            (not obj.get("comment") or obj.get("comment") == "") and \
            (not obj.get("rating") or obj.get("rating") == "") and \
            (not obj.get("rating_date") or obj.get("rating_date") == "") and \
            (not obj.get("song_id") or obj.get("song_id") == "") and \
            (not obj.get("user_id") or obj.get("user_id") == "")