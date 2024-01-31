import streamlit as st
import pandas as pd

# Sample user information
user_info = {
    'AccountHolder': 'John Doe',
    'AccountNumber': '1234567890',
    'CurrentBalance': 1500
}

# Sample user list for demonstrating transfer
user_list = {
    'John Doe': '1234567890',
    'Jane Doe': '0987654321',
    'Bob Smith': '5678901234'
}

# Set page title and icon
st.set_page_config(page_title='Money Transfer', page_icon='💸')

# Sidebar
st.sidebar.title('User Information')
st.sidebar.table(pd.DataFrame([user_info]))

# Function to add a new recipient
def add_recipient():
    new_recipient_name = st.text_input('Enter Recipient Name')
    new_recipient_account = st.text_input('Enter Recipient Account Number')
    if st.button('Add Recipient') and new_recipient_name and new_recipient_account:
        user_list[new_recipient_name] = new_recipient_account
        st.success(f'New recipient {new_recipient_name} added successfully!')
        return True
    return False

# Main content
st.title('Money Transfer')

# Show the list of recipients
recipient_name = st.selectbox('Select Recipient', ['', *user_list.keys()])

# Form to perform money transfer
transfer_amount = st.number_input('Enter Transfer Amount', min_value=0.01, format='%f', step=0.01)
transfer_notes = st.text_area('Add Notes (Optional)', max_chars=200)

# Perform transfer when the user clicks the 'Transfer' button
if st.button('Transfer') and recipient_name:
    recipient_account = user_list.get(recipient_name, '')
    new_balance = user_info['CurrentBalance'] - transfer_amount
    st.success(f'Transfer of ${transfer_amount:.2f} to {recipient_name} successful!')
    st.info(f'New Balance: ${new_balance:.2f}')

# Display the current list of recipients
st.subheader('Recipient List')
st.table(pd.DataFrame(list(user_list.items()), columns=['Recipient Name', 'Account Number']))

# Button to add a new recipient
if add_recipient():
    # Update frontend immediately
    st.experimental_rerun()
