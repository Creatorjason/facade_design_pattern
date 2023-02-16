package main


import "fmt"

type walletFacade struct{
	account *account
	wallet *wallet
	securityCode *securityCode
	notification *notification
	ledger *ledger
}

func newWalletFacade(accountID string, code int) *walletFacade{
	fmt.Println("Starting create account")
	walletFacade := &walletFacade{
		account : newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet : newWallet(),
		notification : &notification{},
		ledger : &ledger{},
	}
	fmt.Println("Account created")
	return walletFacade
}

func(wf *walletFacade) addMoneyToWallet(accountID string, code, amount int) error{
	fmt.Println("Starting add money to wallet")
	err := wf.account.checkAccount(accountID)
	if err != nil{
		return err
	}
	err = wf.securityCode.checkSecurityCode(code)
	if err != nil{
		return err
	}

	wf.wallet.creditBalance(amount)
	wf.notification.sendWalletCreditNotification()
	wf.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (wf *walletFacade) deductMoneyFromWallet(accountID string, code, amount int) error {
	fmt.Println("Starting debit from wallet")
	err := wf.account.checkAccount(accountID)
	if err != nil{
		return err
	}
	err = wf.securityCode.checkSecurityCode(code)
	if err != nil{
		return err
	}
	err = wf.wallet.debitBalance(amount)
	if err != nil{
		return err
	}
	wf.notification.sendWalletDebitNotification()
	wf.ledger.makeEntry(accountID, "debit", amount)
	return nil
}