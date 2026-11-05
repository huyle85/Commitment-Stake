# Commitment Stake

A Web3-powered deposit system that eliminates no-shows at student events. Participants stake USDC through a smart contract to register. Checking in on site with a QR code triggers an automatic refund. If a participant fails to show up, their deposit is distributed among the attendees or sent to the event fund, as the organiser chose in advance.

Built for UNIHACK.

## The Problem

Student events and crypto meetups in Vietnam see 40-60% no-shows. Organisers who order food and seats for the registered number lose money, because free registration carries no cost of commitment.

## Key Features

- **Smart Contract Escrow:** Event deposits are locked on-chain, so no centralized middleman holds the funds.
- **Trustless Settlement:** After the event, the smart contract distributes no-show deposits automatically and transparently, either to the wallets of actual attendees or to the event fund.
- **QR-to-Chain Verification:** The organiser signs a check-in QR (EIP-712, valid for 5 minutes). Scanning it makes the contract verify the signature and refund the deposit instantly.
- **Immutable Transparency:** Registrations, check-ins and distributions are publicly verifiable on the ledger.

## How It Works

1. The organiser creates an event: deposit amount (2-5 USDC), time, and the rule for no-show deposits.
2. A participant registers by staking the deposit into the contract.
3. On site, the participant scans the organiser's QR code and submits the signature to the contract. The contract verifies it and refunds the deposit immediately.
4. When the event ends, anyone can settle it. No-show deposits are then claimed by attendees or sent to the event fund.

## Tech Stack

- **Backend:** Golang (QR signing service and API)
- **Blockchain:** Solidity smart contracts on Avalanche C-Chain, EIP-712 typed signatures, RPC integration
- **Database:** MySQL (off-chain indexing and event metadata)
- **Frontend:** Next.js mobile web (in progress)

## Project Structure

```
commitment-stake/
├── config/        chain and database configuration
├── controllers/   event and QR endpoints
├── models/        database schema
├── services/      EIP-712 check-in signer, QR content
├── contracts/     CommitmentPool, MockUSDC
├── main.go
└── .env.example
```

## Getting Started

1. Copy `.env.example` to `.env` and fill in the values.
2. Run the tests:

```
go test ./services -v
```

3. Start the server:

```
go run .
```

## Status

Backend signer and smart contracts are implemented. Event and QR endpoints, the frontend and the Fuji testnet demo are in progress.
