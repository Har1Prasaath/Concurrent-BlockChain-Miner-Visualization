import React from 'react';
import '../styles/components/TransactionList.css';

function TransactionList({ transactions }) {
  if (!transactions || transactions.length === 0) {
    return <div className="no-transactions">No transactions in this block</div>;
  }
  
  return (
    <div className="transaction-list">
      {transactions.map((tx, index) => (
        <div key={index} className="transaction-item">
          <div className="transaction-header">
            <span className="transaction-id">ID: {tx.ID.substring(0, 8)}...</span>
            <span className="transaction-time">
              {new Date(tx.Timestamp * 1000).toLocaleString()}
            </span>
          </div>
          
          <div className="transaction-details">
            <div className="transaction-addresses">
              <div className="transaction-from">
                <span className="label">From:</span>
                <span className="value">{tx.Sender}</span>
              </div>
              <div className="transaction-arrow">→</div>
              <div className="transaction-to">
                <span className="label">To:</span>
                <span className="value">{tx.Recipient}</span>
              </div>
            </div>
            
            <div className="transaction-amount">
              <span className="amount-value">{tx.Amount}</span>
              <span className="amount-currency">Coins</span>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}

export default TransactionList;