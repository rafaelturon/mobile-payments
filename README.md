## Decred wallet ledger services

### Features
* Blockchain services and wallet
  - Connect to network
  - Automatically starts wallet services
  - Enables access to more critical commands using 2FA
    * Device shutdown
    * Ticket buying
* Exposes a secure API service
  - Balance
    * Available Balance
  - Tickets stats
    * Own Mempool
    * Immature
    * Live
    * Total Subsidy (in DCR)


### Architecture
* Wallet agent (docker)
  - Security and connectivity health check
  - Private key generation
    * Exodus process (12+ seed words, password, recovery link)
    https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki
  - Address management
  - Transactions
    * Send transaction
    * Ticket buying
* Ledger services (cloud)
  - Accounts
  - Payments
    * Invoices
    * Fiat and Crypto Transactions History
    * Third party payment processors
* Pricing services (cloud)
