# TP middleware example

## Run

Download modules :
```
pip install -r requirements.txt
```
On Debian : `pip install -r requirements.txt --break-system-packages`  

Run (command line) from base directory (*flask_base*) :
```
PYTHONPATH=$PYTHONPATH:$(pwd) python3 src/app.py
```

If a warning about your PATH appears :  
```
export PATH=$PATH:$HOME/.local/bin
```

## Documentation

Documentation is visible in **/api/docs** when running the app.

There is also final [API docs for front compatibility](api_documentations/flask_api_for_front_compatibility/swagger.json)
and Go [Ratings API docs](api_documentations/go_api_ratings/swagger.json).