import streamlit as st
import pandas as pd

def show():
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

    # Button to add a new recipient
    with st.expander('Add New Recipient'):
        with st.form(key='add_recipient_form'):
            new_recipient_name = st.text_input('Enter Recipient Name')
            new_recipient_account = st.text_input('Enter Recipient Account Number')
            if st.form_submit_button('Add Recipient') and new_recipient_name and new_recipient_account:
                user_list[new_recipient_name] = new_recipient_account
                st.success(f'New recipient {new_recipient_name} added successfully!')

    # Main content
    st.title('Money Transfer')

    # Show the list of recipients
    recipient_name = st.selectbox('Select Recipient', ['', *user_list.keys()])

    # Form to perform money transfer
    with st.form(key='money_transfer_form'):
        recipient_account = user_list.get(recipient_name, '')
        transfer_amount = st.number_input('Enter Transfer Amount', min_value=0.01, format='%f', step=0.01)
        transfer_notes = st.text_area('Add Notes (Optional)', max_chars=200)
        print(transfer_notes)
        # Perform transfer when the user clicks the 'Transfer' button
        if st.form_submit_button('Transfer') and recipient_account:
            new_balance = user_info['CurrentBalance'] - transfer_amount
            st.success(f'Transfer of ${transfer_amount:.2f} to {recipient_name} successful!')
            st.info(f'New Balance: ${new_balance:.2f}')

    # Display the current list of recipients
    st.subheader('Recipient List')
    st.table(pd.DataFrame(list(user_list.items()), columns=['Recipient Name', 'Account Number']))
