import yaml
from flask import request


def content_negociation(out_body, out_code):
    accepted_type = request.accept_mimetypes.best_match(['application/json', 'application/yaml'])
    if accepted_type == 'application/yaml':
        return yaml.dump(out_body, default_flow_style=False), out_code, {'Content-Type': 'application/yaml'}
    else:
        return out_body, out_code, {'Content-Type': 'application/json'}
