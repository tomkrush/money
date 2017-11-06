## Overview

The money program is designed to help families with their finances. At the moment, this program only accepts CSV files with transactin data. The CSV parser right now only works with Capital One 360 CSV files. It may work with other csv files, but this has not been tested.

## Motivation

I struggle to use any mainstream personal finance tool. I have used Mint and plenty of iOS apps. Each tool has it's own problem. My big problem with all of these tools is getting the proper visibility.

## Goals

1. I want to see how much money I have available to me for each day. This amount should take into account future planned expenses for each month.
2. I want to create saving goals that are small buckets that allow me to save for small things like an oil change or large things like a house.
3. I want to share all this information with my wife. No finance tool is designed for families. This must be family friendly.

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

`money --path finances.csv > finances.md`

## Roadmap

- [x] Parse CSV file from Capital One 360
- [x] Create transactions
- [x] Create accounts
- [x] Calculate starting balance
- [x] Calculate sum of transactions & accounts
- [x] Create transaction mapping rules
    - [x] Create want & need
    - [x] Create categories
    - [ ] Create location based categories
    - [x] Create transaction name re-mapping
    - [ ] Create bills
- [ ] Restructure the transaction system. I got carried away on the concept of accounts. This is not my intent. By design, I want a single stream of transactions.
- [ ] Create bill reports
- [ ] Think about web interface
- [ ] Create SMS notification for families

## Interface

Phase 1:
Create command line tool that outputs reports in Markdown. This will allow me to keep snapshots of my finanial history in apps like Day One or ia Writer.

Phase 2:
Create web interface.

Phase 3:
Integrate with SMS. The idea of this will allow for the 50,000 foot view on finances. This way both my wife and I have a good understanding of where we are at today.
- The sms should display our net worth.
- How many bills have been paid this month
- How many bills are left to be paid this month
- What my daily allowance is today.
- What my remaining allowance for the week is.