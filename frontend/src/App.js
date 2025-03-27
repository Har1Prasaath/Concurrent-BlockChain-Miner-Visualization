import React, { useState } from 'react';
import { BlockchainProvider } from './context/BlockchainContext';
import Header from './components/Header';
import BlockchainDisplay from './components/BlockchainDisplay';
import TransactionsPage from './components/TransactionsPage';
import TransactionCreator from './components/TransactionCreator';
import MiningControl from './components/MiningControl';
import BlockchainValidator from './components/BlockchainValidator';
import './App.css';

function App() {
  const [activeTab, setActiveTab] = useState('blockchain');

  return (
    <BlockchainProvider>
      <div className="app-container">
        <Header activeTab={activeTab} setActiveTab={setActiveTab} />
        
        <main className="main-content">
          <div className="sidebar">
            <TransactionCreator />
            <MiningControl />
            <BlockchainValidator />
          </div>
          
          <div className="blockchain-container">
            {activeTab === 'blockchain' && <BlockchainDisplay />}
            {activeTab === 'transactions' && <TransactionsPage />}
            {activeTab === 'mining' && <div className="coming-soon">Mining Dashboard Coming Soon</div>}
          </div>
        </main>
        
        <footer className="footer">
          <p>Blockchain Visualizer &copy; {new Date().getFullYear()} - Educational Tool</p>
        </footer>
      </div>
    </BlockchainProvider>
  );
}

export default App;