import streamlit as st
from streamlit_option_menu import option_menu

def nav_sidebar(
            options:list[str]=['Login', 'Create Account', 'Forgot Password?', 'Reset Password'],
            menu_title:str="Navigation",
            icons:list[str]=['box-arrow-in-right', 'person-plus', 'x-circle','arrow-counterclockwise']
            ) -> tuple:
        """
        Creates the side navigaton bar
        """
        main_page_sidebar = st.sidebar.empty()
        with main_page_sidebar:
            selected_option = option_menu(
                menu_title = menu_title,
                menu_icon = 'list-columns-reverse',
                icons = icons,
                options = options,
                styles = {
                    "container": {"padding": "5px"},
                    "nav-link": {"font-size": "14px", "text-align": "left", "margin":"0px"}} )
        return main_page_sidebar, selected_option
