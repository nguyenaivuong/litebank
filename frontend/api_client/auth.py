import requests
import streamlit as st

BASE_URL = "http://localhost:8080"
LOGIN_ENDPOINT = "/login"
SIGNUP_ENDPOINT = "/signup"
FORGOT_PASSWORD_ENDPOINT = "/forgot-password"

def login(username, password):
    api_url = BASE_URL + LOGIN_ENDPOINT
    payload = {"username": username, "password": password}

    try:
        response = requests.post(api_url, json=payload)
        if response.status_code == 200:
            return True
        else:
            st.error(f"Login failed: {response.json().get('error')}")
    except requests.exceptions.RequestException as e:
        st.error(f"Error during login request: {e}")

    return False


def signup(username, password, email, full_name):
    payload = {"username": username, "password": password, "email": email, "full_name": full_name}
    api_url = BASE_URL + SIGNUP_ENDPOINT

    try:
        response = requests.post(api_url, json=payload)
        if response.status_code == 201:
            st.success('Account created successfully! Please login to continue.')
            return True
        else:
            st.error(f"Login failed: {response.json().get('error')}")
    except requests.exceptions.RequestException as e:
        st.error(f"Error during login request: {e}")

    return False

def forgot_password(username):
    data = {"username": username}
    return requests.post(BASE_URL + FORGOT_PASSWORD_ENDPOINT, json=data).json()
