# Specification

## Loan
```go
message Loan {
  uint64 id = 1;
  string amount = 2; 
  string fee = 3; 
  string collateral = 4; 
  string deadline = 5; 
  string state = 6; 
  string borrower = 7; 
  string lender = 8;  
}
```

## Borrower
The borrower will request a loan with the required information (amount, fee, collateral, borrower).

## Lender
The lender will decide whether to approve a loan. If the loan is not paid then they can liquidate it receiving the collateral and fee.