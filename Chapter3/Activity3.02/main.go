package main

import (
	"errors"
	"fmt"
)

const (
	goodCreditScore int = 450
	lowCreditRatio float64 = 0.1
	goodCreditRatio float64 = 0.2
)

var (
	creditScoreError = errors.New("Bad Credit Score!")
	monthlyIncomeError = errors.New("Monthly Income not Valid!")
	loanAmtError = errors.New("Invalid Loan Amount")
	loanTermError = errors.New("Invalid Loan Term")
)

func checkLoan(creditScore int, income float64, loanAmount float64, loanTerm int) error {
	interest := 20
	if creditScore >= 450 {
		interest = 15
	}

	if creditScore < 1{
		return creditScoreError
	}
	if income < 1 {
		return monthlyIncomeError
	}
	if loanAmount < 1 {
		return loanAmtError
	}
	if loanTerm < 1 || loanTerm % 12 != 0 {
		return loanTermError
	}

	rate := float64(interest) / 100
	monthlyPayment := ((loanAmount * rate)/float64(loanTerm)) + (loanAmount/float64(loanTerm)) 
	totalInterest := (monthlyPayment * float64(loanTerm)) - loanAmount

	approved := false

	if income > float64(monthlyPayment) {
		var ratio float64 = (float64(monthlyPayment)/income) * 100
		if creditScore > goodCreditScore && ratio > (goodCreditRatio) {
			approved = true
		} else if ratio < lowCreditRatio {
			approved = false
		}
	}

	fmt.Println("Credit Score    :", creditScore)
	fmt.Println("Income          :", income)
	fmt.Println("Loan Amount     :", loanAmount)
	fmt.Println("Loan Term       :", loanTerm)
	fmt.Println("Monthly Payment :", monthlyPayment)
	fmt.Println("Rate            :", interest)
	fmt.Println("Total Cost      :", totalInterest)
	fmt.Println("Approved        :", approved)
	fmt.Println("")

	return nil
}

func main() {
	// Approved
	fmt.Println("Applicant 1")
	fmt.Println("-----------")
	err := checkLoan(500, 1000, 1000, 24)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Denied
	fmt.Println("Applicant 2")
	fmt.Println("-----------")
	err = checkLoan(350, 1000, 10000, 12)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}