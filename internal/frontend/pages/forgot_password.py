import streamlit as st
from api_client.auth import forgot_password

def show() -> None:
    """
    Creates the forgot password widget and after user authentication (email), triggers an email to the user
    containing a random password.
    """
    with st.form("Forgot Password Form"):
        email_forgot_passwd = st.text_input("Email", placeholder='Please enter your email')
        st.markdown("###")
        forgot_passwd_submit_button = st.form_submit_button(label='Get Password')

        if forgot_passwd_submit_button:
            response = forgot_password(email_forgot_passwd)
            handle_response(response)

def handle_response(response):
    print(response)
    if response.status_code == 200:
        st.session_state["token"] = response.json().get("token")
        st.success("Secure Password Sent Successfully!")
    elif response.status_code == 404:
        st.error("Email ID not registered with us!")
    else:
        st.error("Something went wrong!")
