import React, { createContext, useState, useContext, useEffect } from 'react';
import { fetchBlockchain, mineBlock, createTransaction } from '../services/api';

const BlockchainContext = createContext();

export function useBlockchain() {
  return useContext(BlockchainContext);
}

export function BlockchainProvider({ children }) {
  const [blockchain, setBlockchain] = useState({ chain: [], length: 0 });
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [isValid, setIsValid] = useState(true);

  // Load blockchain data
  const loadBlockchain = async () => {
    setLoading(true);
    try {
      const data = await fetchBlockchain();
      setBlockchain(data);
      setError(null);
      validateBlockchain(data.chain);
    } catch (err) {
      console.error('Error loading blockchain:', err);
      setError('Failed to load blockchain data');
    } finally {
      setLoading(false);
    }
  };

  // Mine a new block
  const handleMineBlock = async () => {
    try {
      setLoading(true);
      await mineBlock();
      await loadBlockchain();
      return { success: true };
    } catch (err) {
      console.error('Mining error:', err);
      return { success: false, error: err.message };
    } finally {
      setLoading(false);
    }
  };

  // Create a new transaction
  const handleCreateTransaction = async (transactionData) => {
    try {
      setLoading(true);
      await createTransaction(transactionData);
      await loadBlockchain();
      return { success: true };
    } catch (err) {
      console.error('Transaction error:', err);
      return { success: false, error: err.message };
    } finally {
      setLoading(false);
    }
  };

  // Validate the blockchain
  const validateBlockchain = (chain) => {
    if (!chain || chain.length <= 1) {
      setIsValid(true);
      return;
    }

    let valid = true;
    
    // Simple validation logic - in a real app, this would be more complex
    for (let i = 1; i < chain.length; i++) {
      const currentBlock = chain[i];
      const previousBlock = chain[i-1];
      
      // Check if previous hash matches
      if (currentBlock.PreviousHash !== previousBlock.Hash) {
        valid = false;
        break;
      }
      
      // Check block hash format
      if (!currentBlock.Hash.startsWith('0'.repeat(currentBlock.Difficulty))) {
        valid = false;
        break;
      }
    }
    
    setIsValid(valid);
  };

  // Load blockchain on component mount
  useEffect(() => {
    loadBlockchain();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const value = {
    blockchain,
    loading,
    error,
    isValid,
    refreshBlockchain: loadBlockchain,
    mineBlock: handleMineBlock,
    createTransaction: handleCreateTransaction,
    validateBlockchain
  };

  return (
    <BlockchainContext.Provider value={value}>
      {children}
    </BlockchainContext.Provider>
  );
}