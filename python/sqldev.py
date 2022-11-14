import time
import urllib.parse

import requests

import sign


class Sqldev:
    def __init__(self, endpoint, app_key, app_secret):
        self.endpoint = endpoint
        self.app_key = app_key
        self.app_secret = app_secret

    def send_request(self, method, path, payload):
        params = payload
        params['app_key'] = self.app_key
        params['current_time'] = int(time.time() * 1000)
        params['sign'] = sign.get_signature(params, self.app_secret)
        if method == 'GET':
            url = self.endpoint + path + '?' + urllib.parse.urlencode(params)
            response = requests.get(url)
        elif method == 'POST':
            url = self.endpoint + path
            response = requests.post(url, data=params)
        else:
            return 'unsupported method'
        return response.text

    def query(self, project, iid, sql, owner, schema, page_index, page_size, need_total):
        payload = {
            'project': project,
            'iid': iid,
            'sql': sql,
            'owner': owner,
            'schema': schema,
            'page_index': page_index,
            'page_size': page_size,
            'need_total': need_total,
        }
        return self.send_request('POST', '/sql/query', payload)

    def get_user(self, id):
        params = {
            'id': id,
        }
        return self.send_request('POST', '/user/info', params)

    def get_user_page(self, form):
        params = {
            'page_no': form.page_no,
            'page_size': form.page_size,
            'username': form.username,
            'nickname': form.nickname,
            'state': form.state,
        }
        return self.send_request('POST', '/user/page', params)

    def add_user(self, form):
        params = {
            'username': form.username,
            'nickname': form.nickname,
            'password': form.password,
            'state': form.state,
        }
        return self.send_request('POST', '/user/add', params)

    def upd_user(self, form):
        params = {
            'id': form.id,
            'username': form.username,
            'nickname': form.nickname,
            'password': form.password,
            'state': form.state,
        }
        return self.send_request('POST', '/user/upd', params)

    def remove_user(self, id):
        params = {
            'id': id,
        }
        return self.send_request('POST', '/user/remove', params)

    def state_user(self, form):
        params = {
            'id': form.id,
            'state': form.state,
        }
        return self.send_request('POST', '/user/state', params)

    def get_instance_info(self, id):
        params = {
            "id": id
        }
        res = self.send_request("GET", "/instance/info", params)
        return res

    def get_instance_page(self, form):
        params = {
            "page_no": form.page_no,
            "page_size": form.page_size,
            "name": form.name
        }
        res = self.send_request("GET", "/instance/page", params)
        return res

    def add_instance(self, form):
        params = {
            "name": form.name,
            "host": form.host,
            "port": form.port,
            "username": form.username,
            "password": form.password,
            "db_type": form.db_type,
            "state": form.state
        }
        res = self.send_request("POST", "/instance/add", params)
        return res

    def update_instance(self, form):
        params = {
            "id": form.id,
            "name": form.name,
            "host": form.host,
            "port": form.port,
            "username": form.username,
            "password": form.password,
            "db_type": form.db_type,
            "state": form.state
        }
        res = self.send_request("POST", "/instance/upd", params)
        return res

    def remove_instance(self, id):
        params = {
            "id": id
        }
        res = self.send_request("POST", "/instance/remove", params)
        return res

    def state_instance(self, form):
        params = {
            "id": form.id,
            "state": form.state
        }
        res = self.send_request("POST", "/instance/state", params)
        return res
