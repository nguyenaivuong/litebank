import streamlit as st
from components.utils import check_unique_email, check_valid_email, check_valid_name, check_unique_usr
from api_client.auth import signup

def show() -> None:
    """
    Creates the sign-up widget and stores the user info in a secure way in the _secret_auth_.json file.
    """
    with st.form("Sign Up Form"):
        name_sign_up = st.text_input("Name *", placeholder = 'Please enter your name')
        email_sign_up = st.text_input("Email *", placeholder = 'Please enter your email')
        username_sign_up = st.text_input("Username *", placeholder = 'Enter a unique username')
        password_sign_up = st.text_input("Password *", placeholder = 'Create a strong password', type = 'password')

        st.markdown("###")
        sign_up_submit_button = st.form_submit_button(label = 'Register')

        if sign_up_submit_button:
            error_message = validate_form(name_sign_up, email_sign_up, username_sign_up)
            if error_message:
                st.error(error_message)
            else:
                signup(username_sign_up, password_sign_up, email_sign_up, name_sign_up)


def validate_form(name_sign_up, email_sign_up, username_sign_up):
    if not check_valid_name(name_sign_up):
        return "Please enter a valid name!"
    if not check_valid_email(email_sign_up):
        return "Please enter a valid Email!"
    if not check_unique_email(email_sign_up):
        return "Email already exists!"
    if not check_unique_usr(username_sign_up):
        return f'Sorry, username {username_sign_up} already exists!'
    if username_sign_up.strip() == '':
        return 'Please enter a non - empty Username!'
    return None
