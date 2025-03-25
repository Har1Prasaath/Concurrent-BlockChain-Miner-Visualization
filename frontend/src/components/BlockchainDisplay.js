import React, { useState } from 'react';
import { useBlockchain } from '../context/BlockchainContext';
import BlockDetails from './BlockDetails';
import '../styles/components/BlockchainDisplay.css';

function BlockchainDisplay() {
  const { blockchain, loading, error, refreshBlockchain } = useBlockchain();
  const [selectedBlock, setSelectedBlock] = useState(null);
  
  if (loading) return <div className="loading-container">Loading blockchain data...</div>;
  if (error) return <div className="error-container">{error}</div>;
  
  return (
    <div className="blockchain-display">
      <div className="blockchain-header">
        <h2>Blockchain Explorer</h2>
        <button className="refresh-button" onClick={refreshBlockchain}>
          Refresh
        </button>
      </div>
      
      {/* Visual blockchain representation */}
      <div className="blockchain-visualization">
        {blockchain.chain && blockchain.chain.map((block, index) => (
          <div key={index} className="block-container">
            {index > 0 && <div className="block-connector"></div>}
            <div 
              className={`block-card ${selectedBlock === index ? 'selected' : ''}`}
              onClick={() => setSelectedBlock(index)}
            >
              <div className="block-header">
                <h3>Block #{block.Index}</h3>
                <span className="block-timestamp">
                  {new Date(block.Timestamp * 1000).toLocaleString()}
                </span>
              </div>
              <div className="block-hash">
                <span className="label">Hash:</span>
                <span className="value">{block.Hash.substring(0, 15)}...</span>
              </div>
              <div className="block-transactions">
                <span className="label">Transactions:</span>
                <span className="value">{block.Transactions.length}</span>
              </div>
            </div>
          </div>
        ))}
      </div>
      
      {/* Selected block details */}
      {selectedBlock !== null && blockchain.chain && blockchain.chain[selectedBlock] && (
        <BlockDetails block={blockchain.chain[selectedBlock]} />
      )}
    </div>
  );
}

export default BlockchainDisplay;