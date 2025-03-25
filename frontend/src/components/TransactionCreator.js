import React, { useState } from 'react';
import { useBlockchain } from '../context/BlockchainContext';
import '../styles/components/TransactionCreator.css';

function TransactionCreator() {
  const { createTransaction } = useBlockchain();
  const [formData, setFormData] = useState({
    sender: '',
    recipient: '',
    amount: ''
  });
  const [status, setStatus] = useState(null);
  const [isSubmitting, setIsSubmitting] = useState(false);
  
  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: name === 'amount' ? parseFloat(value) || '' : value
    });
  };
  
  const handleSubmit = async (e) => {
    e.preventDefault();
    
    // Validate form
    if (!formData.sender || !formData.recipient || !formData.amount) {
      setStatus({
        type: 'error',
        message: 'All fields are required'
      });
      return;
    }
    
    setIsSubmitting(true);
    setStatus(null);
    
    try {
      const result = await createTransaction({
        sender: formData.sender,
        recipient: formData.recipient,
        amount: parseFloat(formData.amount)
      });
      
      if (result.success) {
        setStatus({
          type: 'success',
          message: 'Transaction created successfully!'
        });
        // Reset form
        setFormData({
          sender: '',
          recipient: '',
          amount: ''
        });
      } else {
        setStatus({
          type: 'error',
          message: result.error || 'Failed to create transaction'
        });
      }
    } catch (error) {
      setStatus({
        type: 'error',
        message: error.message || 'Failed to create transaction'
      });
    } finally {
      setIsSubmitting(false);
    }
  };
  
  return (
    <div className="transaction-creator">
      <h2>Create Transaction</h2>
      
      {status && (
        <div className={`status-message ${status.type}`}>
          {status.message}
        </div>
      )}
      
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="sender">From (Sender)</label>
          <input
            type="text"
            id="sender"
            name="sender"
            value={formData.sender}
            onChange={handleChange}
            disabled={isSubmitting}
            placeholder="Sender's address"
          />
        </div>
        
        <div className="form-group">
          <label htmlFor="recipient">To (Recipient)</label>
          <input
            type="text"
            id="recipient"
            name="recipient"
            value={formData.recipient}
            onChange={handleChange}
            disabled={isSubmitting}
            placeholder="Recipient's address"
          />
        </div>
        
        <div className="form-group">
          <label htmlFor="amount">Amount</label>
          <div className="amount-input">
            <input
              type="number"
              id="amount"
              name="amount"
              value={formData.amount}
              onChange={handleChange}
              disabled={isSubmitting}
              placeholder="0.00"
              min="0.01"
              step="0.01"
            />
            <span className="currency">COINS</span>
          </div>
        </div>
        
        <button 
          type="submit" 
          disabled={isSubmitting}
          className="submit-button"
        >
          {isSubmitting ? 'Processing...' : 'Create Transaction'}
        </button>
      </form>
    </div>
  );
}

export default TransactionCreator;