---
tags:
  - finance
time: 20:28
date: 2026-01-01
---
I want to create a personal finance system. This will be a continuous project. For now, I just want to lay the very basic like data saving foundation of it. At the moment, I'll work on the income and spending flow. Very basic.

I think these are the rows that I want to build the foundation on.
- id
- transaction date
- details - basically you just put there what you bought or the name of the item, doesn't matter but it is required
- transaction type
	- income
	- expense
- currency
- amount - how much money is involved in this transaction
- wallet - need a better name for this. basically this is either where the money enters or where the money comes from. So like let's say it's cash on hand, from an e-wallet, a bank account, etc.

So, I think there will be another table for like users, and then the opened accounts of the users. This is now where I fall in a pit for designing the database of this system.

I need to be able to use this for multiple users.