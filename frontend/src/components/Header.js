import React, { useState } from 'react';
import { useBlockchain } from '../context/BlockchainContext';
import '../styles/components/Header.css';

function Header() {
  const { blockchain, isValid } = useBlockchain();
  const [activeTab, setActiveTab] = useState('blockchain');
  
  return (
    <header className="app-header">
      <div className="header-title">
        <h1>Blockchain Visualizer</h1>
        <div className={`blockchain-status ${isValid ? 'valid' : 'invalid'}`}>
          {isValid ? 'Blockchain Valid' : 'Blockchain Invalid'}
        </div>
      </div>
      
      <div className="blockchain-info">
        <div className="info-item">
          <span className="info-label">Blocks:</span>
          <span className="info-value">{blockchain.length}</span>
        </div>
        <div className="info-item">
          <span className="info-label">Latest Hash:</span>
          <span className="info-value hash">
            {blockchain.chain && blockchain.chain.length > 0 
              ? `${blockchain.chain[blockchain.chain.length-1].Hash.substring(0, 10)}...` 
              : 'None'}
          </span>
        </div>
      </div>
      
      <nav className="header-nav">
        <ul>
          <li 
            className={activeTab === 'blockchain' ? 'active' : ''}
            onClick={() => setActiveTab('blockchain')}
          >
            Blockchain
          </li>
          <li 
            className={activeTab === 'transactions' ? 'active' : ''}
            onClick={() => setActiveTab('transactions')}
          >
            Transactions
          </li>
          <li 
            className={activeTab === 'mining' ? 'active' : ''}
            onClick={() => setActiveTab('mining')}
          >
            Mining
          </li>
        </ul>
      </nav>
    </header>
  );
}

export default Header;