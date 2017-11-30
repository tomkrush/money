[![Travis CI](https://travis-ci.org/tomkrush/money.svg?branch=master)](https://travis-ci.org/tomkrush/money)
[![Documentation](https://godoc.org/github.com/tomkrush/money?status.svg)](https://godoc.org/github.com/tomkrush/money)
[![codecov](https://codecov.io/gh/tomkrush/money/branch/master/graph/badge.svg)](https://codecov.io/gh/tomkrush/money)


## Overview

The money program is designed to help families with their finances. At the moment, this program only displays a summary of the current months financial history. The CSV parser right now only works with Capital One 360 CSV files. It may work with other csv files, but this has not been tested.

## Motivation

I struggle to use any mainstream personal finance tool. I have used Mint and plenty of iOS apps. Each tool has it's own problem. My big problem with all of these tools is getting the proper visibility.

## Goals

1. I want to see how much money I have available to me for each day. This amount should take into account future planned expenses for each month.
2. I want to create saving goals that are small buckets that allow me to save for small things like an oil change or large things like a house.
3. I want to share all this information with my wife. No finance tool is designed for families. This must be family friendly.

## High Level Idea's

### Allowance
Allowance represents the amount of money this month left over after your projected income minus your projected bills minus your expenses so far.

### Bills
Bills are identified on a transaction rule. Bills contain a projected amount. Once a bill is identified as paid, the bill will use the actual amount.

### Planned Expenses
Planned expenses is a checklist of expected expenses in the current month. A planned expense is composed of a description and amount. Planned expenses are now explicitly associated with transactions. Marking paid will only reduce your projected planned expenses.

### Wants vs. Needs
Learning how to save money requires you to differeniate between needs and wants. Transaction rules apply need/want status to each transaction.

### Transaction Rules
Transaction rules are stored in JSON. Rules that match a transaction description can manipulate the following attributes.
- Bill
- Income
- Category
- Description

## Installation
- go get https://github.com/tomkrush/money-go

## Tests
- go test

## Install
- go install

## Usage
If your GOPATH is setup correctly than you should be able to use the `money` command line app from any folder.

Right now this app only accepts csv files outputted by Capital One. I personally output the csv file for my entire year.

The usefulness of this program will grow as items are completed on the roadmap.

`money --path finances.csv --rules rules.json`

## Roadmap

- [x] Parse CSV file from Capital One 360
- [x] Create transactions
- [x] Create accounts
- [x] Calculate starting balance
- [x] Calculate sum of transactions & accounts
- [x] Create transaction mapping rules
    - [x] Create want & need
    - [x] Create categories
    - [-] Create location based categories
    - [x] Create transaction name re-mapping
    - [x] Create bills
    - [ ] Create planned transactions
- [x] Restructure the transaction system. I got carried away on the concept of accounts. This is not my intent. By design, I want a single stream of transactions.
- [-] Create bill reports
- [x] Think about web interface
- [-] Create SMS notification for families

## User Interface

The user interface right now is designed specifically for the mobile screen. I am specifically choosing to not create Android and iOS apps. Maybe later I will.

The mobile web interface is meant to be added to the user's phone screen.
