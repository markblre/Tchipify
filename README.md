# Tchipify
Projet réalisé dans le cadre du master 1 informatique à l'Université Clermont Auvergne.
## Auteurs
[Ballereau Mark](https://github.com/markblre)  
[Combronde Léna](https://github.com/lena-cbrd)

## Versions utilisées
- Go 1.21.4
- Python 3.12.0
- Node 20.11.0
- NPM 10.2.4

## Installation

Dans ```/API_Flask``` :
```bash
pip install -r requirements.txt
```

## Comment utiliser ?
### Lancement des APIs

#### API Users :
```bash
cd Users && go run cmd/main.go
```

#### API Songs :
```bash
cd Songs && go run cmd/main.go
```

#### API Ratings :
```bash
cd Ratings && go run cmd/main.go
```

#### API Flask (IntelliJ) :
- Définir le répertoire ```/API_Flask``` comme source root
- Ajouter une configuration Python -> Script python, puis sélectionner le chemin du script dans src/app.py.
- Lancer la configuration.

### Lancement de l'interface web
- Rendez-vous dans le répertoire ```/front```

```bash
npm install
npm run dev
```

## Documentations
- [API Users](Users/api/swagger.json)
- [API Songs](Songs/api/swagger.json)
- [API Ratings](Ratings/api/swagger.json)
- [API Flask](http://localhost:8888/api/docs/) (Lancement de l'API Flask requis)


