# 🧱 Basic Blockchain in Golang

A simple yet extendable blockchain implementation in Go. This project includes block creation, node synchronization via HTTP/WebSocket, a minimal React dashboard, and local chain persistence.

## ✨ Features

- Full block and blockchain structure
- Simple Proof-of-Work algorithm
- RESTful API for blockchain access
- Real-time WebSocket broadcasting
- Block mining and syncing
- Persistent storage to `chain.json`
- Frontend dashboard with React

---

## 🗂 Project Structure

```sh

foundational-blockchain/
├── cmd/                    # Main application entry point
│   └── main.go
├── internal/
│   ├── api/                # HTTP and WebSocket logic
│   ├── blockchain/         # Core blockchain logic
│   └── frontend/           # React frontend dashboard
├── pkg/logger/             # Simple logger utility
├── data/                   # Persisted blockchain data
├── go.mod / go.sum         # Dependencies
└── README.md               # This file

```

---

## ⚙️ Getting Started

### 1. Prerequisites

- Go 1.23 or higher  
- Node.js (for frontend dashboard)

### 2. Clone the Repository

```bash
git clone https://github.com/yourusername/basic-blockchain.git
cd basic-blockchain
```

### 3. Run the Node

```bash
go run ./cmd/main.go
```

### 4. Mine a New Block

```bash
curl -X POST http://localhost:8080/mine \
  -H "Content-Type: application/json" \
  -d '{"data":"your block data"}'
```

### 5. Get the Blockchain

```bash
curl http://localhost:8080/chain
```

---

## 💻 React Dashboard (Optional)

To build and serve the frontend dashboard:

```bash
cd internal/frontend
npm install
npm run build
```

Once built, the dashboard will be accessible at:

http://localhost:8080/dashboard

---

## 🔗 API Endpoints

| Endpoint     | Method | Description                           |
|--------------|--------|---------------------------------------|
| `/mine`      | POST   | Mine a new block                      |
| `/chain`     | GET    | Get the current blockchain            |
| `/block`     | POST   | Receive a block from a peer node      |
| `/peers`     | GET    | List connected peers                  |
| `/peers`     | POST   | Add a new peer node                   |
| `/ws`        | GET    | WebSocket for live updates            |

---

## 🔧 TODO

- Add advanced consensus algorithms (PoS, DPoS, Raft)  
- Implement public/private key digital signatures  
- Add data encryption & transaction validation  
- Visualize blockchains in the frontend  

---

## 👨‍💻 Author

Developed by [Alireza](https://github.com/astrica1)  
Designed with clean architecture principles in mind using Go.

---

## 🪪 License

This project is licensed under the MIT License.
