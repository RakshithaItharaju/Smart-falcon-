# Smart-falcon-
# 🔗 Blockchain Asset Management using Hyperledger Fabric

This project implements a blockchain-based asset management system for a financial institution using **Hyperledger Fabric**. It includes a smart contract written in **Golang** and a **REST API** using **Node.js**.
## 📌 Problem Statement
A financial institution wants to use blockchain to manage and track assets. Each asset is an account with attributes like:
- DEALERID
- MSISDN
- MPIN
- BALANCE
- STATUS
- TRANSAMOUNT
- TRANSTYPE
- REMARKS

The system should:
- Create new assets
- Update asset balance
- Query assets
- Retrieve all assets
## 🗂️ Project Structure
project-root/
├── chaincode-go/ # Smart contract code in Go
│ └── asset_chaincode.go
│
├── rest-api/ # Node.js REST API
│ ├── server.js
│ └── Dockerfile
│
└── README.md
Start Test Network
./network.sh down
./network.sh up createChannel -ca
Deploy Chaincode bash
./network.sh deployCC -ccn assetchain -ccp ../asset-transfer-basic/chaincode-go -ccl go
REST API Endpoints
{
  "id": "ACC001",
  "dealerID": "DLR001",
  "msisdn": "9876543210",
  "mpin": "1234",
  "balance": 5000,
  "status": "ACTIVE",
  "transAmount": 1000,
  "transType": "CREDIT",
  "remarks": "Initial deposit"
}
.Build Docker Image
docker build -t asset-api .
.Run Container
docker run -p 3000:3000 asset-api
-----
Developed By
- ITHARAJU RAKSHITHA 
- Sreenidhi institute of science and technology 
