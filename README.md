# vet-booking-cli

## Introduction

This project is a menu-driven CLI program for owners to book veterinary appointments for their pets. Written in Go as my first ever coding project, it demonstrates the Go features and broader engineering concepts I self-taught myself from multiple sources, beginning with the book "Introducing Go" by Caleb Doxsey around the start of the first year of my Software Engineering BSc degree, and then continuing using online blogposts and various LLMs (strictly as tutors!) to continue development at an organic, relaxed pace. I began writing this program following a month of reading and going through the exercises detailed in the aforementioned textbook during my spare time before deciding to write my own small program that I, very thoughtfully, named dogProject! (it was originally just for dogs for some reason)

Over the months it had ballooned into an approximately thousand-line chunk of code that would compile, but was now too large to return to after long breaks to pick up where I had left off. It was around this time that I decided (now about to begin the second year of my degree after summer break) to try and finish this project, first by adding the features that I had loosely envisioned, and then by cleaning it up i.e. splitting up the code into files and organising functions.

As I began taking the first steps, like figuring out where to split up the now monstrously large main.go file, I was informed by a senior mentor (who is an industry expert) that I should implement certain architectural practices into my code now, given that the current state was going to impact any future development. I was introduced to architectural concepts like decoupling the database queries from the business logic, mocking database repositories, and writing modular, maintainable code. 

If I could begin again, I would definitely begin with a structured modular codebase to make my code more scalable from the get-go, and ensure I have a clearer idea of the domain that I am modeling the application for (though still keeping it agile, as I have learnt that "Big Design Up Front" is considered an antipattern!). I would also attempt to work around a more complex domain to really showcase awareness about the domain and its unique business rules that I am building the code for.

Thank you for viewing my project!

## Brief functional overview

Pet owners can log in via using their email and a password to book, view, and delete appointments for their pets via simple multiple-choice menus with free-entry fields where necessary. It uses postgreSQL to persist data in a database and implements password hashing using the bcrypt library.

## Requirements

- Go 1.20+
- PostgreSQL

## Setup

1. Create a PostgreSQL database:
CREATE DATABASE vet_booking;

2. Connect to the database:
\c vet_booking

3. Run the schema:
 - psql vet_booking < schema.sql (Powershell)
 - \i schema.sql (Bash - make sure current directory is project root directory beforehand)

4. Create a .env file in the project root directory using .env.example as a template:
- cp .env.example .env (Linux/MacOS)
- Copy-Item .env.example .env (Windows)

5. Edit .env and replace username and password with YOUR PostgreSQL credentials:
DATABASE_URL=postgres://username:password@localhost:5432/vet_booking?sslmode=disable

6. Change directory to project root

7. Run program:
go run .

## Known limitations

 - Business logic is not fully implemented yet e.g. appointment clashing
 - Error handling is inconsistent/missing in certain parts like legacy code parts (scanner errors)
 - Tests are present but are not extensive enough and dont cover crucial parts of the application
 - Current layout of code across files is messy and lacks splitting into directories and packages
 - Some CLI functions like getUserPhone and getUserEmail are not uniform with similar get functions as they also contain the prompt instead of the prompt being passed in as a parameter
 - Email validation is not extensive enough (perhaps will replace with regex)
 - Multiple layers use the same structs in models.go, this creates some coupling as changing the fields may break code in multiple layers when it should't :(


 - Program is not containerised yet and so is tedious to install and run
 - Maybe should use explicit data transfer objects between service and repo instead of passing user object to repo from memory. Instead pass raw fields to keep it dumb?
 - ID values (ID and userID) perhaps should be integers instead of strings to better match their purpose. Also they work right now with postgreSQL but I am unsure if they would work with another database. Minor, but still noteworthy
 - When a user logs out, it would be nice to get the back to the main menu instead of just ending the program
