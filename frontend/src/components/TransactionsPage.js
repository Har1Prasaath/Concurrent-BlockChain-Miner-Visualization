import React, { useMemo } from 'react';
import { useBlockchain } from '../context/BlockchainContext';
import '../styles/components/TransactionsPage.css';

function TransactionsPage() {
  const { blockchain, loading, error } = useBlockchain();
  
  // Extract and flatten all transactions from the blockchain
  const allTransactions = useMemo(() => {
    if (!blockchain.chain || blockchain.chain.length === 0) {
      return [];
    }
    
    return blockchain.chain
      .flatMap(block => block.Transactions)
      .sort((a, b) => b.Timestamp - a.Timestamp); // Sort by timestamp, newest first
  }, [blockchain.chain]);
  
  if (loading) return <div className="loading-container">Loading transactions...</div>;
  if (error) return <div className="error-container">{error}</div>;
  
  return (
    <div className="transactions-page">
      <div className="transactions-header">
        <h2>All Blockchain Transactions</h2>
        <div className="transaction-count">
          {allTransactions.length} transactions found
        </div>
      </div>
      
      {allTransactions.length === 0 ? (
        <div className="no-transactions-message">
          No transactions found in the blockchain.
        </div>
      ) : (
        <div className="transactions-list">
          {allTransactions.map((tx, index) => (
            <div key={index} className="transaction-item">
              <div className="transaction-header">
                <span className="transaction-id">ID: {tx.ID}</span>
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
              
              <div className="transaction-block-info">
                <span className="label">Included in Block:</span>
                <span className="value">
                  {blockchain.chain.findIndex(block => 
                    block.Transactions.some(t => t.ID === tx.ID)
                  )}
                </span>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default TransactionsPage;