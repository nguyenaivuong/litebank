# online_banking_app.py

import streamlit as st
from pages import forgot_password, login, overview, signup, account_balance, transactions, account_settings
from components.sidebar import nav_sidebar

def logout_widget() -> None:
    """
    Creates the logout widget in the sidebar only if the user is logged in.
    """
    if st.session_state['LOGGED_IN'] == True:
        del_logout = st.sidebar.empty()
        del_logout.markdown("#")
        logout_click_check = del_logout.button("Logout")

        if logout_click_check == True:
            st.session_state['LOGOUT_BUTTON_HIT'] = True
            st.session_state['LOGGED_IN'] = False
            del_logout.empty()
            st.rerun()

def main():
    st.set_page_config(page_title="Online Banking", layout="wide")

    if 'LOGGED_IN' not in st.session_state:
        st.session_state['LOGGED_IN'] = False

    if 'LOGOUT_BUTTON_HIT' not in st.session_state:
        st.session_state['LOGOUT_BUTTON_HIT'] = False

    # If the user is not logged in, show the login page
    if not st.session_state['LOGGED_IN']:
        options = ['Login', 'Create Account', 'Forgot Password?']
        _, selected_option = nav_sidebar(options)
        if selected_option == 'Login':
            login.show()
        elif selected_option == 'Create Account':
            signup.show()
        elif selected_option == 'Forgot Password?':
            st.write('Forgot Password?')
        else:
            st.write('Choose an option from the sidebar')

    # If the user is authenticated, show the selected page and sidebar
    if st.session_state['LOGGED_IN']:
        logout_widget()
        st.sidebar.title("Lite Bank")

        options = ["Overview", "Transactions", "Change Password"]
        icons=['house', 'currency-dollar', 'gear']
        _, selected_option = nav_sidebar(options, menu_title="Lite Bank", icons=icons)

        # Show the selected page content
        if selected_option == "Overview":
            overview.show()

        elif selected_option == "Account Balance":
            account_balance.show()

        elif selected_option == "Transactions":
            transactions.show()

        elif selected_option == "Change Password":
            forgot_password.show()

if __name__ == "__main__":
    main()
