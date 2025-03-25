# Blockchain Visualizer

This project is an interactive blockchain visualization application implemented with Go for the backend and React for the frontend. It allows users to explore the blockchain, create transactions, mine new blocks, and validate the integrity of the blockchain - all with a modern, user-friendly interface.

## Features

- **Interactive Blockchain Explorer:** Visualize the blockchain with connected blocks
- **Block Details View:** Inspect detailed information about each block
- **Transaction Management:** Create and view transactions
- **Mining Control:** Mine new blocks with visual feedback
- **Blockchain Validation:** Verify the integrity of the blockchain
- **Responsive Design:** Works on desktop and mobile devices

## Project Structure

```
blockchain-web-app/
├── backend/                # Go backend server
│   ├── blockchain/         # Blockchain implementation
│   ├── api/                # API handlers
│   ├── models/             # Data models
│   └── main.go             # Entry point
├── frontend/               # React frontend application
│   ├── public/             # Static files
│   ├── src/                # React components and logic
│   │   ├── components/     # UI components
│   │   ├── services/       # API services
│   │   ├── pages/          # Page components
│   │   └── App.js          # Main application component
│   ├── package.json        # Dependencies and scripts
│   └── README.md           # Frontend documentation
└── README.md               # This file
```

## Installation

### Prerequisites
- Go 1.16+
- Node.js 14+
- npm or yarn

### Backend Setup
```bash
# Navigate to backend directory
cd backend

# Install dependencies
go mod tidy

# Run the server
go run main.go
```

### Frontend Setup
```bash
# Navigate to frontend directory
cd frontend

# Install dependencies
npm install
# or
yarn install

# Start development server
npm start
# or
yarn start
```

## Usage

1. Start both the backend and frontend servers as described above
2. Open your browser and navigate to `http://localhost:3000`
3. The genesis block is created automatically on first load
4. Create transactions using the Transaction form
5. Mine new blocks to include pending transactions
6. Explore the blockchain by clicking on blocks to view details
7. Validate the blockchain integrity using the Validate button

## Technologies Used

- **Backend:**
  - Go
  - Gorilla Mux (API routing)
  - Crypto (for hashing)
  
- **Frontend:**
  - React
  - Redux (state management)
  - Axios (API client)
  - D3.js (visualization)
  - Material-UI (components)

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/blockchain` | GET | Get the entire blockchain |
| `/api/blocks` | GET | Get all blocks |
| `/api/blocks/:id` | GET | Get a specific block by index |
| `/api/mine` | POST | Mine a new block |
| `/api/transactions` | GET | Get all pending transactions |
| `/api/transactions` | POST | Create a new transaction |
| `/api/validate` | GET | Validate the blockchain integrity |

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.
