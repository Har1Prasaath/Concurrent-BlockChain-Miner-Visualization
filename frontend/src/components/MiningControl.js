import React, { useState } from 'react';
import { useBlockchain } from '../context/BlockchainContext';
import '../styles/components/MiningControl.css';

function MiningControl() {
  const { mineBlock, blockchain } = useBlockchain();
  const [isMining, setIsMining] = useState(false);
  const [miningStatus, setMiningStatus] = useState(null);
  const [miningAnimation, setMiningAnimation] = useState(false);
  
  const handleMineBlock = async () => {
    setIsMining(true);
    setMiningStatus(null);
    setMiningAnimation(true);
    
    setTimeout(async () => {
      try {
        const result = await mineBlock();
        
        if (result.success) {
          setMiningStatus({
            type: 'success',
            message: 'Block mined successfully!'
          });
        } else {
          setMiningStatus({
            type: 'error',
            message: result.error || 'Failed to mine block'
          });
        }
      } catch (error) {
        setMiningStatus({
          type: 'error',
          message: error.message || 'Error during mining process'
        });
      } finally {
        setIsMining(false);
        setTimeout(() => setMiningAnimation(false), 1000);
      }
    }, 1500); // Simulate mining delay for effect
  };
  
  const latestBlock = blockchain.chain && blockchain.chain.length > 0 
    ? blockchain.chain[blockchain.chain.length - 1] 
    : null;
  
  return (
    <div className="mining-control">
      <h2>Mine New Block</h2>
      
      <div className="mining-info">
        <div className="info-row">
          <span className="info-label">Latest Block:</span>
          <span className="info-value">
            {latestBlock ? `#${latestBlock.Index}` : 'None'}
          </span>
        </div>
        <div className="info-row">
          <span className="info-label">Current Difficulty:</span>
          <span className="info-value">
            {latestBlock ? latestBlock.Difficulty : 'N/A'}
          </span>
        </div>
      </div>
      
      {miningStatus && (
        <div className={`mining-status ${miningStatus.type}`}>
          {miningStatus.message}
        </div>
      )}
      
      <div className={`mining-animation-container ${miningAnimation ? 'active' : ''}`}>
        <div className="mining-animation">
          <div className="mining-block"></div>
          <div className="mining-pickaxe">⛏️</div>
        </div>
      </div>
      
      <button
        className="mine-button"
        onClick={handleMineBlock}
        disabled={isMining}
      >
        {isMining ? 'Mining...' : 'Mine New Block'}
      </button>
    </div>
  );
}

export default MiningControl;