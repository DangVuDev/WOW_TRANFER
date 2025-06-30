WoWToken Web API Documentation
Font: Times New Roman
This document outlines the RESTful API endpoints for interacting with the WoWToken smart contract, designed for a decentralized banking dApp tailored for non-technical users. The APIs are divided into two sections: those directly interacting with the smart contract and user management APIs for a custodial wallet-based banking application. Detailed descriptions are provided to clarify the purpose, prerequisites, processing flow, and error cases for each API.
Base URL
https://api.wowtoken.example.com/v1
Authentication

APIs require an API key passed in the X-API-Key header or JWT token for authenticated endpoints.
User management APIs use email/password-based authentication for non-technical users. The backend manages wallet private keys securely (custodial model).
Ensure TLS is used for all requests.

Error Handling
Errors return HTTP status codes (400, 401, 403, 500, etc.) with a JSON response:
{
  "error": "Description of the error"
}

API Endpoints
1. APIs Directly Interacting with WoWToken Smart Contract
These APIs correspond to the functions and data structures in the WoWToken smart contract, including token management, multi-signature events, and contract state queries.



Endpoint
Method
Description
Parameters
Response
Example
Detailed Description



/balances/{address}
GET
Retrieve the token balance of a given address.
Path: address (string, Ethereum address)
JSON: {"address": string, "balance": string}
GET /balances/0x123... → {"address": "0x123...", "balance": "1000000000000000000"}
Purpose: Allows users or the system to check the WoWToken balance of a specific Ethereum address, useful for account management or transaction validation.  Prerequisites: The address must be a valid Ethereum address. The contract must not be paused.  Processing Flow: 1. Validate the input address format. 2. Query the balances mapping in the WoWToken contract using a blockchain provider (e.g., Infura). 3. Return the balance in wei.  Error Cases: - 400: Invalid address format. - 500: Blockchain query failure (e.g., provider downtime).


/transfer
POST
Transfer tokens from one address to another (used by backend for user-initiated transfers).
Body: {"from": string, "to": string, "amount": string, "jwt": string}
JSON: {"transactionHash": string, "status": string}
POST /transfer with {"from": "0x123...", "to": "0x456...", "amount": "1000000000000000000", "jwt": "ey..."} → {"transactionHash": "0x789...", "status": "success"}
Purpose: Enables token transfers between addresses, typically triggered by user actions (e.g., sending money to another user).  Prerequisites: Valid JWT token, from address must have sufficient balance, to address must be valid, contract not paused.  Processing Flow: 1. Validate JWT to authenticate the user. 2. Verify from and to address formats and amount validity. 3. Retrieve the from address's encrypted private key from the database, decrypt it, and sign a transferToken transaction. 4. Submit the transaction to the blockchain and wait for confirmation. 5. Return the transaction hash.  Error Cases: - 401: Invalid or expired JWT. - 400: Invalid addresses or amount. - 403: Insufficient balance or contract paused. - 500: Transaction failure (e.g., gas limit exceeded).


/approve
POST
Approve a spender to transfer tokens on behalf of an address (used by backend for user approvals).
Body: {"owner": string, "spender": string, "amount": string, "jwt": string}
JSON: {"transactionHash": string, "status": string}
POST /approve with {"owner": "0x123...", "spender": "0x456...", "amount": "500000000000000000", "jwt": "ey..."} → {"transactionHash": "0x789...", "status": "success"}
Purpose: Allows a user to authorize another address to spend tokens on their behalf, useful for delegated payments or smart contract interactions.  Prerequisites: Valid JWT, owner and spender must be valid addresses, contract not paused.  Processing Flow: 1. Validate JWT and input parameters. 2. Retrieve the owner address's encrypted private key, decrypt it, and sign an approve transaction. 3. Submit the transaction to the blockchain. 4. Return the transaction hash.  Error Cases: - 401: Invalid JWT. - 400: Invalid addresses or amount. - 403: Contract paused. - 500: Transaction failure.


/transferFrom
POST
Transfer tokens from one address to another by an approved spender (used by backend).
Body: {"from": string, "to": string, "amount": string, "jwt": string}
JSON: {"transactionHash": string, "status": string}
POST /transferFrom with {"from": "0x123...", "to": "0x456...", "amount": "1000000000000000000", "jwt": "ey..."} → {"transactionHash": "0x789...", "status": "success"}
Purpose: Enables an approved spender to transfer tokens from one address to another, used for automated payments or third-party services.  Prerequisites: Valid JWT, from address must have sufficient balance and allowance for the spender, to address valid, contract not paused.  Processing Flow: 1. Validate JWT and input parameters. 2. Verify allowance via the allowance mapping. 3. Retrieve the spender's encrypted private key, decrypt, and sign a transferFrom transaction. 4. Submit the transaction and return the hash.  Error Cases: - 401: Invalid JWT. - 400: Invalid addresses or amount. - 403: Insufficient balance, allowance, or contract paused. - 500: Transaction failure.


/allowance/{owner}/{spender}
GET
Check the allowance of a spender for a given owner.
Path: owner (string), spender (string)
JSON: {"owner": string, "spender": string, "allowance": string}
GET /allowance/0x123.../0x456... → {"owner": "0x123...", "spender": "0x456...", "allowance": "500000000000000000"}
Purpose: Retrieves the amount of tokens a spender is allowed to transfer on behalf of an owner, useful for verifying delegated permissions.  Prerequisites: Valid owner and spender addresses.  Processing Flow: 1. Validate address formats. 2. Query the allowance mapping in the contract. 3. Return the allowance value.  Error Cases: - 400: Invalid address format. - 500: Blockchain query failure.


/events/mint
POST
Create a multi-signature event to mint new tokens (restricted to co-owners).
Body: {"amount": string, "jwt": string}
JSON: {"eventId": string, "transactionHash": string, "status": string}
POST /events/mint with {"amount": "1000000000000000000", "jwt": "ey..."} → {"eventId": "mint_1698765432", "transactionHash": "0x789...", "status": "success"}
Purpose: Initiates a multi-signature process to mint new tokens to the bank reserve, ensuring governance control.  Prerequisites: Valid JWT, caller must be a co-owner, contract not paused, amount within LIMIT_TOKEN.  Processing Flow: 1. Validate JWT and co-owner status via co_token_owner. 2. Verify amount against LIMIT_TOKEN - current_total_token. 3. Sign a mintToken transaction using the co-owner's private key. 4. Create a multi-signature event in event_requireMultiSignature. 5. Return event ID and transaction hash.  Error Cases: - 401: Invalid JWT. - 403: Not a co-owner or contract paused. - 400: Invalid amount or exceeds limit. - 500: Transaction failure.


/events/burn
POST
Create a multi-signature event to burn tokens from the reserve (restricted to co-owners).
Body: {"amount": string, "jwt": string}
JSON: {"eventId": string, "transactionHash": string, "status": string}
POST /events/burn with {"amount": "1000000000000000000", "jwt": "ey..."} → {"eventId": "burn_1698765432", "transactionHash": "0x789...", "status": "success"}
Purpose: Initiates a multi-signature process to burn tokens from the bank reserve, reducing total supply.  Prerequisites: Valid JWT, caller is a co-owner, contract not paused, sufficient reserve balance.  Processing Flow: 1. Validate JWT and co-owner status. 2. Verify amount against balances[bank_reserve]. 3. Sign a generateBurnTokenEvent transaction. 4. Create a multi-signature event. 5. Return event ID and transaction hash.  Error Cases: - 401: Invalid JWT. - 403: Not a co-owner or contract paused. - 400: Invalid amount or insufficient reserve balance. - 500: Transaction failure.


/events/add-co-owner
POST
Create a multi-signature event to add a new co-owner (restricted to co-owners).
Body: {"newCoOwner": string, "jwt": string}
JSON: {"eventId": string, "transactionHash": string, "status": string}
POST /events/add-co-owner with {"newCoOwner": "0x456...", "jwt": "ey..."} → {"eventId": "add_co_1698765432", "transactionHash": "0x789...", "status": "success"}
Purpose: Initiates a multi-signature process to add a new co-owner for contract governance.  Prerequisites: Valid JWT, caller is a co-owner, newCoOwner not already a co-owner, contract not paused.  Processing Flow: 1. Validate JWT and co-owner status. 2. Verify newCoOwner address and status. 3. Sign an addCoOwner transaction. 4. Create a multi-signature event. 5. Return event ID and transaction hash.  Error Cases: - 401: Invalid JWT. - 403: Not a co-owner or contract paused. - 400: Invalid or existing co-owner address. - 500: Transaction failure.


/events/pause
POST
Create a multi-signature event to pause the contract (restricted to co-owners).
Body: {"jwt": string}
JSON: {"eventId": string, "transactionHash": string, "status": string}
POST /events/pause with {"jwt": "ey..."} → {"eventId": "pause_1698765432", "transactionHash": "0x789...", "status": "success"}
Purpose: Initiates a multi-signature process to pause contract operations for security or maintenance.  Prerequisites: Valid JWT, caller is a co-owner, contract not already paused.  Processing Flow: 1. Validate JWT and co-owner status. 2. Check contract pause status. 3. Sign a generatePauseEvent transaction. 4. Create a multi-signature event. 5. Return event ID and transaction hash.  Error Cases: - 401: Invalid JWT. - 403: Not a co-owner or contract already paused. - 500: Transaction failure.


/events/unpause
POST
Create a multi-signature event to unpause the contract (restricted to co-owners).
Body: {"jwt": string}
JSON: {"eventId": string, "transactionHash": string, "status": string}
POST /events/unpause with {"jwt": "ey..."} → {"eventId": "unpause_1698765432", "transactionHash": "0x789...", "status": "success"}
Purpose: Initiates a multi-signature process to resume contract operations.  Prerequisites: Valid JWT, caller is a co-owner, contract is paused.  Processing Flow: 1. Validate JWT and co-owner status. 2. Check contract pause status. 3. Sign a generateUnpauseEvent transaction. 4. Create a multi-signature event. 5. Return event ID and transaction hash.  Error Cases: - 401: Invalid JWT. - 403: Not a co-owner or contract not paused. - 500: Transaction failure.


/events/sign
POST
Sign a multi-signature event (approve or reject, restricted to co-owners).
Body: {"eventId": string, "approve": boolean, "jwt": string}
JSON: {"eventId": string, "transactionHash": string, "status": string}
POST /events/sign with {"eventId": "mint_1698765432", "approve": true, "jwt": "ey..."} → {"eventId": "mint_1698765432", "transactionHash": "0x789...", "status": "success"}
Purpose: Allows a co-owner to approve or reject a multi-signature event, ensuring decentralized governance.  Prerequisites: Valid JWT, caller is a co-owner, event exists and is not completed, caller has not signed.  Processing Flow: 1. Validate JWT and co-owner status. 2. Verify event existence and status. 3. Sign a signEvent transaction with the co-owner's private key. 4. Update eventSigners and signature_count. 5. Return event ID and transaction hash.  Error Cases: - 401: Invalid JWT. - 403: Not a co-owner, event completed, or already signed. - 400: Invalid event ID. - 500: Transaction failure.


/events/execute/{eventType}
POST
Execute a multi-signature event (mint, burn, add-co-owner, pause, unpause, restricted to co-owners).
Path: eventType (string: mint, burn, add-co-owner, pause, unpause)  Body: {"eventId": string, "jwt": string}
JSON: {"eventId": string, "transactionHash": string, "status": string}
POST /events/execute/mint with {"eventId": "mint_1698765432", "jwt": "ey..."} → {"eventId": "mint_1698765432", "transactionHash": "0x789...", "status": "success"}
Purpose: Finalizes a multi-signature event after sufficient approvals, executing the corresponding action.  Prerequisites: Valid JWT, caller is a co-owner, event exists, not completed, has >50% signatures, contract not paused (except for unpause).  Processing Flow: 1. Validate JWT, co-owner status, and event type. 2. Verify event and signature count. 3. Sign the corresponding execute* transaction (e.g., executeMintTokenEvent). 4. Update contract state and mark event as completed. 5. Return event ID and transaction hash.  Error Cases: - 401: Invalid JWT. - 403: Not a co-owner, insufficient signatures, or event completed. - 400: Invalid event ID or type. - 500: Transaction failure.


/events/{eventId}
GET
Retrieve details of a multi-signature event.
Path: eventId (string)
JSON: {"id_event": string, "isCompleted": boolean, "event_name": string, "description": string, "signature_count": number, "signers": string[], "createdAt": number}
GET /events/mint_1698765432 → {"id_event": "mint_1698765432", "isCompleted": false, "event_name": "Mint Token", "description": "Mint token", "signature_count": 2, "signers": ["0x123...", "0x456..."], "createdAt": 1698765432}
Purpose: Provides details of a multi-signature event for transparency or governance tracking.  Prerequisites: Valid event ID.  Processing Flow: 1. Validate event ID format. 2. Query event_requireMultiSignature[eventId] in the contract. 3. Return event details.  Error Cases: - 400: Invalid event ID. - 404: Event does not exist. - 500: Blockchain query failure.


/contract/state
GET
Retrieve the contract's state.
None
JSON: {"token_owner": string, "co_owner_count": number, "bank_reserve": string, "current_total_token": string, "paused": boolean}
GET /contract/state → {"token_owner": "0x123...", "co_owner_count": 3, "bank_reserve": "0x456...", "current_total_token": "1000000000000000000", "paused": false}
Purpose: Provides an overview of the contract's current state for monitoring or auditing purposes.  Prerequisites: None.  Processing Flow: 1. Query contract variables (token_owner, co_owner_count, bank_reserve, current_total_token, paused). 2. Return the state in JSON format.  Error Cases: - 500: Blockchain query failure.


/co-owners
GET
Retrieve the list of co-owners.
None
JSON: {"co_owners": string[]}
GET /co-owners → {"co_owners": ["0x123...", "0x456..."]}
Purpose: Lists all co-owners for governance transparency.  Prerequisites: None.  Processing Flow: 1. Query co_token_owner mapping to retrieve addresses with true status. 2. Return the list of co-owner addresses.  Error Cases: - 500: Blockchain query failure.


/burn
POST
Burn tokens from an address (used by backend for user-initiated burns).
Body: {"address": string, "amount": string, "jwt": string}
JSON: {"transactionHash": string, "status": string}
POST /burn with {"address": "0x123...", "amount": "1000000000000000000", "jwt": "ey..."} → {"transactionHash": "0x789...", "status": "success"}
Purpose: Allows users to burn tokens from their balance, reducing their holdings.  Prerequisites: Valid JWT, address has sufficient balance, contract not paused.  Processing Flow: 1. Validate JWT and input parameters. 2. Retrieve the address's encrypted private key, decrypt, and sign a burnToken transaction. 3. Submit the transaction and return the hash.  Error Cases: - 401: Invalid JWT. - 400: Invalid address or amount. - 403: Insufficient balance or contract paused. - 500: Transaction failure.


2. User Management APIs for Banking dApp (Custodial Wallet)
These APIs support user management and banking functionalities for non-technical users, using a custodial wallet model where the backend manages Ethereum wallets. Users interact with the dApp via email/password, and the backend handles blockchain interactions.



Endpoint
Method
Description
Parameters
Response
Example
Detailed Description



/users/register
POST
Register a new user, creating a custodial Ethereum wallet and storing off-chain metadata.
Body: {"email": string, "password": string, "name": string}
JSON: {"userId": string, "address": string, "status": string}
POST /users/register with {"email": "john@example.com", "password": "pass123", "name": "John Doe"} → {"userId": "uuid123", "address": "0x123...", "status": "success"}
Purpose: Enables non-technical users to create an account with email/password, automatically generating a custodial Ethereum wallet for blockchain interactions.  Prerequisites: Unique email, valid password (e.g., minimum 8 characters), valid name.  Processing Flow: 1. Validate email, password, and name. 2. Hash the password (e.g., using bcrypt). 3. Generate a new Ethereum wallet using ethers.Wallet.createRandom(). 4. Encrypt the private key (e.g., using AES-256 or AWS KMS). 5. Store {"userId", "email", "hashedPassword", "name", "address", "encryptedPrivateKey"} in the database. 6. Return userId and address.  Error Cases: - 400: Invalid email, password, or name format. - 409: Email already registered. - 500: Database or wallet creation failure.


/users/login
POST
Authenticate a user and return a JWT token for subsequent requests.
Body: {"email": string, "password": string}
JSON: {"jwt": string, "address": string, "status": string}
POST /users/login with {"email": "john@example.com", "password": "pass123"} → {"jwt": "ey...", "address": "0x123...", "status": "success"}
Purpose: Authenticates users to access protected APIs, providing a JWT token for session management.  Prerequisites: User must be registered with valid email/password.  Processing Flow: 1. Validate email and password. 2. Compare password with stored hash. 3. Generate a JWT token containing user data (e.g., address). 4. Return JWT and associated address.  Error Cases: - 401: Invalid email or password. - 400: Missing email/password. - 500: Database or JWT generation failure.


/users/{address}
GET
Retrieve user profile information.
Path: address (string)  Header: Authorization: Bearer <jwt>
JSON: {"address": string, "name": string, "email": string, "registeredAt": number}
GET /users/0x123... → {"address": "0x123...", "name": "John Doe", "email": "john@example.com", "registeredAt": 1698765432}
Purpose: Retrieves user profile details for display in the dApp (e.g., account settings).  Prerequisites: Valid JWT, address matches the JWT's user address.  Processing Flow: 1. Validate JWT and ensure address matches the JWT's user. 2. Query the database for user data by address. 3. Return user profile information.  Error Cases: - 401: Invalid or expired JWT. - 403: Address does not match JWT user. - 404: User not found. - 500: Database query failure.


/users/{address}/update
PUT
Update user profile information (off-chain).
Path: address (string)  Body: {"name": string, "email": string, "password": string, "jwt": string}
JSON: {"address": string, "status": string}
PUT /users/0x123.../update with {"name": "Jane Doe", "email": "jane@example.com", "password": "newpass123", "jwt": "ey..."} → {"address": "0x123...", "status": "success"}
Purpose: Allows users to update their profile information (e.g., name, email, password) for account management.  Prerequisites: Valid JWT, address matches JWT user, new email not already registered.  Processing Flow: 1. Validate JWT and ensure address matches. 2. Validate new email, name, and password. 3. Hash new password if provided. 4. Update user data in the database. 5. Return updated status.  Error Cases: - 401: Invalid JWT. - 403: Address does not match JWT user. - 400: Invalid email, name, or password format. - 409: New email already registered. - 500: Database update failure.


/users/{address}/transactions
GET
Retrieve transaction history for a user.
Path: address (string)  Query: limit (number, optional), offset (number, optional)  Header: Authorization: Bearer <jwt>
JSON: {"address": string, "transactions": [{"type": string, "from": string, "to": string, "amount": string, "timestamp": number, "transactionHash": string}]}
GET /users/0x123.../transactions?limit=10&offset=0 → {"address": "0x123...", "transactions": [{"type": "TransferToken", "from": "0x123...", "to": "0x456...", "amount": "1000000000000000000", "timestamp": 1698765432, "transactionHash": "0x789..."}]}
Purpose: Provides a history of user transactions (e.g., transfers, burns) for transparency and account tracking.  Prerequisites: Valid JWT, address matches JWT user.  Processing Flow: 1. Validate JWT and address. 2. Query blockchain events (e.g., TransferToken, BurnToken) for the address using a service like Moralis or Alchemy. 3. Apply pagination with limit and offset. 4. Return transaction list.  Error Cases: - 401: Invalid JWT. - 403: Address does not match JWT user. - 400: Invalid pagination parameters. - 500: Blockchain query failure.


/users/deposit
POST
Deposit tokens to the bank reserve from the user's custodial wallet.
Body: {"amount": string, "jwt": string}
JSON: {"transactionHash": string, "status": string}
POST /users/deposit with {"amount": "1000000000000000000", "jwt": "ey..."} → {"transactionHash": "0x789...", "status": "success"}
Purpose: Allows users to deposit tokens into the bank reserve, similar to depositing money into a bank account.  Prerequisites: Valid JWT, user has sufficient balance, contract not paused.  Processing Flow: 1. Validate JWT and extract user address. 2. Verify amount and balance via balances[userAddress]. 3. Retrieve and decrypt the user's private key. 4. Sign a transferToken(bank_reserve, amount) transaction. 5. Submit the transaction and return the hash.  Error Cases: - 401: Invalid JWT. - 400: Invalid amount. - 403: Insufficient balance or contract paused. - 500: Transaction failure.


/users/withdraw
POST
Create a multi-signature event to withdraw tokens from the bank reserve to the user's custodial wallet.
Body: {"amount": string, "jwt": string}
JSON: {"eventId": string, "transactionHash": string, "status": string}
POST /users/withdraw with {"amount": "1000000000000000000", "jwt": "ey..."} → {"eventId": "withdraw_1698765432", "transactionHash": "0x789...", "status": "success"}
Purpose: Initiates a multi-signature process to withdraw tokens from the bank reserve to the user's wallet, ensuring secure fund release.  Prerequisites: Valid JWT, sufficient reserve balance, contract not paused.  Processing Flow: 1. Validate JWT and extract user address. 2. Verify amount against balances[bank_reserve]. 3. Create a multi-signature event with ID withdraw_<timestamp>. 4. Store event in event_requireMultiSignature and event_requireMultiSignatureAmount. 5. Sign the event creation transaction using a co-owner's private key (if the caller is a co-owner). 6. Return event ID and transaction hash.  Error Cases: - 401: Invalid JWT. - 400: Invalid amount. - 403: Insufficient reserve balance or contract paused. - 500: Transaction failure.


Notes

Custodial Wallet Model:
Upon registration (/users/register), the backend creates an Ethereum wallet using ethers.Wallet.createRandom().
The private key is encrypted (e.g., using AES-256 or AWS KMS) and stored securely in a database, linked to the user's email and address.
Example wallet creation in Node.js with ethers.js:const { ethers } = require('ethers');
const wallet = ethers.Wallet.createRandom();
const address = wallet.address; // e.g., 0x123...
const encryptedPrivateKey = encrypt(wallet.privateKey); // Store securely


The backend signs transactions on behalf of users using the stored private key, eliminating the need for users to interact with wallets like MetaMask.


Authentication:
Users authenticate via /users/login with email/password, receiving a JWT token for subsequent requests.
JWT tokens are validated in the Authorization header for protected endpoints.


User Data Storage:
User metadata (email, name, encrypted private key) is stored off-chain in a database (e.g., MongoDB, PostgreSQL) for scalability, linked to the Ethereum address.


Transaction History:
The /users/{address}/transactions endpoint queries blockchain events (e.g., TransferToken, BurnToken, Approval) using services like Moralis or Alchemy.


Deposit/Withdraw Logic:
Deposit: The backend calls transferToken(bank_reserve, amount) using the user's custodial private key.
Withdraw: Creates a multi-signature event to transfer tokens from bank_reserve to the user's address, requiring approval from more than 50% of co-owners.


Blockchain Interaction:
APIs interact with the Ethereum blockchain via providers like Infura or Alchemy using libraries such as web3.js or ethers.js.


Security:
Private keys are encrypted and stored securely (e.g., using AWS KMS or HSM).
Implement rate-limiting, TLS, and JWT validation to prevent abuse.
Use two-factor authentication (2FA) for sensitive actions like withdrawals.


Caching:
Use caching (e.g., Redis) for read-heavy endpoints like /balances/{address}, /users/{address}, and /events/{eventId} to reduce blockchain queries.



For further details or support, contact the WoWToken development team at support@wowtoken.example.com.