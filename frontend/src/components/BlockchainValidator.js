import React, { useState } from 'react';
import { useBlockchain } from '../context/BlockchainContext';
import '../styles/components/BlockchainValidator.css';

function BlockchainValidator() {
  const { blockchain, isValid, validateBlockchain } = useBlockchain();
  const [validationDetails, setValidationDetails] = useState(null);
  
  const runFullValidation = () => {
    if (!blockchain.chain || blockchain.chain.length <= 1) {
      setValidationDetails({
        valid: true,
        message: 'Blockchain only has one block (Genesis block).',
        issues: []
      });
      return;
    }
    
    const issues = [];
    let valid = true;
    
    // Perform validation checks
    for (let i = 1; i < blockchain.chain.length; i++) {
      const currentBlock = blockchain.chain[i];
      const previousBlock = blockchain.chain[i-1];
      
      // Check previous hash
      if (currentBlock.PreviousHash !== previousBlock.Hash) {
        issues.push({
          block: currentBlock.Index,
          issue: 'Previous hash does not match the hash of the previous block'
        });
        valid = false;
      }
      
      // Check hash format (starts with zeros based on difficulty)
      if (!currentBlock.Hash.startsWith('0'.repeat(currentBlock.Difficulty))) {
        issues.push({
          block: currentBlock.Index,
          issue: `Hash does not meet difficulty requirement of ${currentBlock.Difficulty} leading zeros`
        });
        valid = false;
      }
    }
    
    setValidationDetails({
      valid,
      message: valid 
        ? 'Blockchain is valid. All blocks are properly linked and have valid hashes.' 
        : 'Blockchain validation failed. See issues below.',
      issues
    });
    
    // Update the global validation state
    validateBlockchain(blockchain.chain);
  };
  
  return (
    <div className="blockchain-validator">
      <h2>Blockchain Validator</h2>
      
      <div className={`validation-status ${isValid ? 'valid' : 'invalid'}`}>
        <div className="status-indicator"></div>
        <span className="status-text">
          {isValid ? 'Blockchain is currently valid' : 'Blockchain integrity issues detected'}
        </span>
      </div>
      
      <button 
        className="validate-button" 
        onClick={runFullValidation}
      >
        Run Full Validation
      </button>
      
      {validationDetails && (
        <div className="validation-details">
          <h3>Validation Results</h3>
          <p className={validationDetails.valid ? 'valid-message' : 'invalid-message'}>
            {validationDetails.message}
          </p>
          
          {validationDetails.issues.length > 0 && (
            <>
              <h4>Issues Found:</h4>
              <ul className="issues-list">
                {validationDetails.issues.map((issue, index) => (
                  <li key={index} className="issue-item">
                    <span className="issue-block">Block #{issue.block}:</span> {issue.issue}
                  </li>
                ))}
              </ul>
            </>
          )}
        </div>
      )}
    </div>
  );
}

export default BlockchainValidator;