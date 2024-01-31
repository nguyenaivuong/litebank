import requests


BASE_URL = "http://localhost:8080"
LOGIN_ENDPOINT = "/login"
SIGNUP_ENDPOINT = "/signup"
RENEW_TOKEN_ENDPOINT = "/tokens/renew_access"
CREATE_ACCOUNT_ENDPOINT = "/accounts"
GET_ACCOUNT_ENDPOINT = "/accounts/{}"
UPDATE_ACCOUNT_ENDPOINT = "/accounts/{}"
DELETE_ACCOUNT_ENDPOINT = "/accounts/{}"


def login(username, password):
    data = {"username": username, "password": password}
    return requests.post(BASE_URL + LOGIN_ENDPOINT, json=data).json()


def signup(username, password):
    data = {"username": username, "password": password}
    return requests.post(BASE_URL + SIGNUP_ENDPOINT, json=data).json()


def renew_token():
    return requests.post(BASE_URL + RENEW_TOKEN_ENDPOINT).json()


def create_account(balance):
    data = {"balance": balance}
    return requests.post(BASE_URL + CREATE_ACCOUNT_ENDPOINT, json=data).json()


def get_account(account_id):
    return requests.get(BASE_URL + GET_ACCOUNT_ENDPOINT.format(account_id)).json()


def update_account(account_id, new_balance):
    data = {"balance": new_balance}
    return requests.put(BASE_URL + UPDATE_ACCOUNT_ENDPOINT.format(account_id), json=data).json()


def delete_account(account_id):
    return requests.delete(BASE_URL + DELETE_ACCOUNT_ENDPOINT.format(account_id)).json()
