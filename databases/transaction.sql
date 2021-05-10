CREATE TABLE IF NOT EXISTS TB_TRANSACTION (
    TRANSACTION_ID SERIAL PRIMARY KEY,
    TRANSACTION_AMOUNT BIGINT,
    TRANSACTION_DATE DATE,

    ACCOUNT_ID INT,
    TRANSACTION_TYPE INT,

    FOREIGN KEY (ACCOUNT_ID)
        REFERENCES TB_ACCOUNT(ACCOUNT_ID),
    FOREIGN KEY (TRANSACTION_TYPE)
        REFERENCES TB_TRANSACTION_TYPE(TRANSACTION_TYPE_ID)
);