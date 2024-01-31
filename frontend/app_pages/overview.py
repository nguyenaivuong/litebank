import streamlit as st
import pandas as pd
import plotly.express as px
from datetime import datetime

def show():
    # Sample user and financial information
    user_info = {
        'CardNumber': '85810854',
        'AccountHolder': 'NGUYEN AI VUONG',
        'TotalBalance': 300,
        'TotalIncome': 400,
        'TotalExpense': 100
    }

    # Sample data for demonstration purposes
    transactions_data = {
        'Name': ['Alice', 'Bob', 'Hana', 'John', 'Alice', 'John', 'Alice'],
        'Category': ['Groceries', 'Entertainment', 'Utilities', 'Dining', 'Dining', 'Entertainment', 'Utilities'],
        'Date': ['2022-01-04', '2022-01-04', '2022-01-04', '2022-01-03', '2022-01-02','2022-01-02', '2022-01-01'],
        'Amount': [100, -50, 150, -20, 100, -30, 50],
        'Type': ['Income', 'Outcome', 'Income', 'Outcome', 'Income', 'Outcome', 'Income']
    }
    df_transactions = pd.DataFrame(transactions_data)

    # Sidebar for date range selection
    st.sidebar.title('Date Range Selection')
    start_date = st.sidebar.date_input('Start Date', datetime.strptime(min(df_transactions['Date']), '%Y-%m-%d'))
    end_date = st.sidebar.date_input('End Date', datetime.strptime(max(df_transactions['Date']), '%Y-%m-%d'))

    # Filter transactions based on the selected date range
    df_transactions = df_transactions[(df_transactions['Date'] >= start_date.strftime('%Y-%m-%d')) & (df_transactions['Date'] <= end_date.strftime('%Y-%m-%d'))]
    if df_transactions.empty:
        st.warning('No transactions to display.')

    balances = calculate_daily_balance(df_transactions, start_date, end_date)
    if balances.empty:
        st.warning('No transactions to display.')

    overview(user_info)
    transaction_table_with_pagination(df_transactions)
    statistics_graph(df_transactions, balances)



def overview(user_info):
    st.title('Overview')

    # Use st.columns to create two side-by-side columns
    left_column, right_column = st.columns(2)

    # Left column for card information
    with left_column:
        st.markdown(
            f"""
            <div style='border-radius: 25px; background-color: rgb(41, 44, 51); padding: 25px; color: white; height: 200px;'>
                <h3 style='margin: 0;'>Card Information</h3>
                <p style='margin: 0; line-height: 1.5; margin-top: 25px'>Card Number: {user_info['CardNumber']}</p>
                <p style='margin: 0; line-height: 1.5; margin-top: 15px'>Account Holder: {user_info['AccountHolder']}</p>
            </div>
            """,
            unsafe_allow_html=True
        )

    financial_info = {
        'TotalBalance': 300,
        'TotalIncome': 400,
        'TotalExpense': 100
    }

    # Right column for financial information
    with right_column:
        # Container for the main rectangle and two smaller rectangles
        with st.container():
            st.markdown(
                f"""
                <div style='border-radius: 25px; background-color: rgb(41, 44, 54); padding: 25px; color: white; height: 200px;'>
                    <h3 style='font-size: 18px; margin: 0;'>Total Balance</h3>
                    <p style='font-size: 46px; font-weight: bold; margin-top: -10px; margin: 0;'>${financial_info['TotalBalance']}</p>
                    <div style='display: flex; justify-content: space-between; align-items: stretch;'>
                        <div style='width: 48%;'>
                            <div style='border-radius: 10px; background-color: #3366ff; padding: 10px; height: 25px; margin-top: 20px'>
                                <h4 style='font-size: 12px; margin-left: 10px; margin-top: -17px'>Total Income: ${financial_info['TotalIncome']}</h4>
                            </div>
                        </div>
                        <div style='width: 48%;'>
                            <div style='border-radius: 10px; background-color: #ff3333; padding: 10px; height: 25px; margin-top: 20px'>
                                <h4 style='font-size:12px; margin-left: 10px; margin-top: -17px'>Total Expense: ${financial_info['TotalExpense']}</h4>
                            </div>
                        </div>
                    </div>
                </div>
                """,
                unsafe_allow_html=True
            )

def transaction_table(df_transactions):
    # Add space between titles
    st.markdown("<br><br><br>", unsafe_allow_html=True)
    # Main content
    st.title('Recent Transactions')

    # Recent Transactions table with search input at the top-right corner
    search_query = st.text_input('', key='search_transactions', max_chars=20, help='Input a name', placeholder='🔍 Search Transactions by name')

    # Filter transactions based on the search query
    filtered_transactions = df_transactions[df_transactions['Name'].str.contains(search_query, case=False)]
    st.table(filtered_transactions[['Name', 'Category', 'Date', 'Amount', 'Type']])

def statistics_graph(df_transactions, balances):
    # Add space between titles
    st.markdown("<br><br><br>", unsafe_allow_html=True)

    # Main content
    st.title('Statistics')
    fig = None

    # Create checkboxes for selecting options
    show_balance = st.checkbox('Balance')

    # If Balance option is selected, calculate daily balance
    if show_balance:
        chart_option = st.selectbox('Select Chart Option', ['Category', 'Date'])
        if chart_option == 'Category':
            filtered_transactions = df_transactions.groupby(['Date', 'Category']).agg({'Amount': 'sum'}).reset_index()
            fig = px.bar(
                filtered_transactions,
                x='Date',
                y='Amount',
                color='Category',
                barmode='group',  # This groups bars for each Date
                title='Transaction History by Category',
                labels={'Amount': 'Transaction Amount ($)'},
                height=400,
            )
        elif chart_option == 'Date':
            fig = px.line(
                balances,
                x='Date',
                y='Cumulative Balance',
                title='Daily Balance',
                labels={'Cumulative Balance': 'Balance ($)'},
                height=400,
            )
    else:
        chart_option = st.selectbox('Select Chart Option', ['Category', 'Type'])
        if chart_option == 'Category':
            show_income = st.checkbox('Show Income', True)
            show_outcome = st.checkbox('Show Outcome', True)

            if show_income and show_outcome:
                filtered_transactions = df_transactions
            elif show_income:
                filtered_transactions = df_transactions[df_transactions['Type'] == 'Income']
            elif show_outcome:
                filtered_transactions = df_transactions[df_transactions['Type'] == 'Outcome']
            else:
                filtered_transactions = pd.DataFrame()

            if not filtered_transactions.empty:
                fig = px.bar(
                    filtered_transactions,
                    x='Date',
                    y='Amount',
                    color='Category',
                    barmode='group',  # This groups bars for each Date
                    title='Transaction History by Category',
                    labels={'Amount': 'Transaction Amount ($)'},
                    height=400,
                )

        elif chart_option == 'Type':
            fig = px.bar(
                df_transactions,
                x='Date',
                y='Amount',
                color='Type',
                title='Transaction History by Type',
                labels={'Amount': 'Transaction Amount ($)'},
                height=400,
            )

    try:
        st.plotly_chart(fig)
    except Exception as e:
        st.error(f"Error: {str(e)}")

def calculate_daily_balance(df_transactions, start_date, end_date):
    # Calculate the cumulative balance amount for each day
    balance_amounts = []
    cumulative_balance = 0
    filtered_transactions = df_transactions.groupby('Date').agg({'Amount': 'sum'}).reset_index()

    for balance in filtered_transactions['Amount']:
        cumulative_balance += balance
        balance_amounts.append(cumulative_balance)

    filtered_transactions['Cumulative Balance'] = balance_amounts
    filtered_transactions = filtered_transactions[(filtered_transactions['Date'] >= start_date.strftime('%Y-%m-%d')) & (filtered_transactions['Date'] <= end_date.strftime('%Y-%m-%d'))]

    return filtered_transactions

def transaction_table_with_pagination(df_transactions, page_size=5):
    # Add space between titles
    st.markdown("<br><br><br>", unsafe_allow_html=True)
    # Main content
    st.title('Recent Transactions')

    # Recent Transactions table with search input at the top-right corner
    search_query = st.text_input('', key='search_transactions', max_chars=20, help='Input a name', placeholder='🔍 Search Transactions by name')

    # Filter transactions based on the search query
    filtered_transactions = df_transactions[df_transactions['Name'].str.contains(search_query, case=False)]

    # Calculate the number of pages
    num_pages = len(filtered_transactions) // page_size + (len(filtered_transactions) % page_size > 0)

    # Sidebar for page selection
    page_number = 1

    # Add next and previous page buttons with appropriate styling
    col1, col2 = st.columns([6, 1])
    with col1:
        if col1.button('Previous Page', key='prev_page', help='Go to the previous page') and page_number > 1:
            page_number -= 1

    with col2:
        if col2.button('Next Page', key='next_page', help='Go to the next page') and page_number < num_pages:
            page_number += 1

    # Calculate the start and end indices for the selected page
    start_idx = (page_number - 1) * page_size
    end_idx = min(start_idx + page_size, len(filtered_transactions))

    # Display the selected page of data
    st.table(filtered_transactions[['Name', 'Category', 'Date', 'Amount', 'Type']][start_idx:end_idx])
