import requests

songs_url = "http://localhost:8081/songs/"  # URL de l'API songs (golang)


def get_songs():
    response = requests.request(method="GET", url=songs_url)
    # TODO: Ajouter les ratings dans la réponse
    return response.json(), response.status_code