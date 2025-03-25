import React, { useState } from 'react';
import TransactionList from './TransactionList';
import '../styles/components/BlockDetails.css';

function BlockDetails({ block }) {
  const [activeTab, setActiveTab] = useState('overview');
  
  if (!block) return null;
  
  return (
    <div className="block-details">
      <div className="block-details-header">
        <h3>Block #{block.Index} Details</h3>
        <div className="block-details-tabs">
          <button 
            className={`tab ${activeTab === 'overview' ? 'active' : ''}`}
            onClick={() => setActiveTab('overview')}
          >
            Overview
          </button>
          <button 
            className={`tab ${activeTab === 'transactions' ? 'active' : ''}`}
            onClick={() => setActiveTab('transactions')}
          >
            Transactions ({block.Transactions.length})
          </button>
          <button 
            className={`tab ${activeTab === 'raw' ? 'active' : ''}`}
            onClick={() => setActiveTab('raw')}
          >
            Raw Data
          </button>
        </div>
      </div>
      
      <div className="block-details-content">
        {activeTab === 'overview' && (
          <div className="block-overview">
            <div className="detail-row">
              <span className="detail-label">Block Index:</span>
              <span className="detail-value">{block.Index}</span>
            </div>
            <div className="detail-row">
              <span className="detail-label">Timestamp:</span>
              <span className="detail-value">{new Date(block.Timestamp * 1000).toLocaleString()}</span>
            </div>
            <div className="detail-row">
              <span className="detail-label">Nonce:</span>
              <span className="detail-value">{block.Nonce}</span>
            </div>
            <div className="detail-row">
              <span className="detail-label">Difficulty:</span>
              <span className="detail-value">{block.Difficulty}</span>
            </div>
            <div className="detail-row">
              <span className="detail-label">Hash:</span>
              <span className="detail-value hash">{block.Hash}</span>
            </div>
            <div className="detail-row">
              <span className="detail-label">Previous Hash:</span>
              <span className="detail-value hash">
                {block.PreviousHash ? block.PreviousHash : 'Genesis Block'}
              </span>
            </div>
            <div className="detail-row">
              <span className="detail-label">Transaction Count:</span>
              <span className="detail-value">{block.Transactions.length}</span>
            </div>
          </div>
        )}
        
        {activeTab === 'transactions' && (
          <TransactionList transactions={block.Transactions} />
        )}
        
        {activeTab === 'raw' && (
          <div className="block-raw">
            <pre>{JSON.stringify(block, null, 2)}</pre>
          </div>
        )}
      </div>
    </div>
  );
}

export default BlockDetails;