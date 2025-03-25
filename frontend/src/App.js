import React from 'react';
import { BlockchainProvider } from './context/BlockchainContext';
import Header from './components/Header';
import BlockchainDisplay from './components/BlockchainDisplay';
import TransactionCreator from './components/TransactionCreator';
import MiningControl from './components/MiningControl';
import BlockchainValidator from './components/BlockchainValidator';
import './App.css';

function App() {
  return (
    <BlockchainProvider>
      <div className="app-container">
        <Header />
        
        <main className="main-content">
          <div className="sidebar">
            <TransactionCreator />
            <MiningControl />
            <BlockchainValidator />
          </div>
          
          <div className="blockchain-container">
            <BlockchainDisplay />
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