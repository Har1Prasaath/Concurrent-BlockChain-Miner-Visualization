const API_URL = 'http://localhost:8080';

// Fetch the blockchain
export async function fetchBlockchain() {
  try {
    const response = await fetch(`${API_URL}/chain`);
    if (!response.ok) {
      throw new Error(`Server responded with ${response.status}`);
    }
    return await response.json();
  } catch (error) {
    console.error('API error:', error);
    throw new Error(`Failed to fetch blockchain: ${error.message}`);
  }
}

// Mine a new block
export async function mineBlock() {
  try {
    const response = await fetch(`${API_URL}/mine`);
    if (!response.ok) {
      throw new Error(`Server responded with ${response.status}`);
    }
    return await response.json();
  } catch (error) {
    console.error('API error:', error);
    throw new Error(`Failed to mine block: ${error.message}`);
  }
}

// Create a new transaction
export async function createTransaction(transactionData) {
  try {
    const response = await fetch(`${API_URL}/transactions/new`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(transactionData)
    });
    
    if (!response.ok) {
      throw new Error(`Server responded with ${response.status}`);
    }
    
    return await response.json();
  } catch (error) {
    console.error('API error:', error);
    throw new Error(`Failed to create transaction: ${error.message}`);
  }
}