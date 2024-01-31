import streamlit as st
from api_client.auth import login

def show():
    if st.session_state['LOGGED_IN'] == False:
        st.session_state['LOGOUT_BUTTON_HIT'] = False

        del_login = st.empty()
        with del_login.form("Login Form"):
            username = st.text_input("Username", placeholder = 'Your unique username')
            password = st.text_input("Password", placeholder = 'Your password', type = 'password')

            st.markdown("###")
            login_submit_button = st.form_submit_button(label = 'Login')

            if login_submit_button == True:
                ok = login(username, password)
                if ok:
                    st.success("Login successful!")
                    st.session_state['LOGGED_IN'] = True
                    del_login.empty()
                    st.rerun()
