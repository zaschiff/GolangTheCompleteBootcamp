# GolangTheCompleteBootcamp

This repo holds the code for the Udemy Class Go (Golang): The Complete Bootcamp by Jose Portilla and Inanc Gumus

## Table of Contents

- [GolangTheCompleteBootcamp](#golangthecompletebootcamp)
  - [Table of Contents](#table-of-contents)
    - [learngo](#learngo)
    - [MyWork](#mywork)
      - [WriteYourFirstGoProgram](#writeyourfirstgoprogram)
        - [Exercise1Section2](#exercise1section2)
        - [Exercise2Section2](#exercise2section2)
        - [first](#first)
        - [printer](#printer)
        - [scopes](#scopes)
      - [MastertheTypeSystemofGo](#masterthetypesystemofgo)
        - [Greeter](#greeter)
        - [CelsiusToFahrenheit](#celsiustofahrenheit)
        - [Banger](#banger)
        - [Exercise1Section3](#exercise1section3)
        - [Exercise2Section3](#exercise2section3)
        - [Exercise3Section3](#exercise3section3)
        - [Exercise4Section3](#exercise4section3)
        - [Exercise5Section3](#exercise5section3)
        - [Exercise6Section3](#exercise6section3)
        - [Exercise7Section3](#exercise7section3)
        - [Exercise8Section3](#exercise8section3)
        - [Exercise9Section3](#exercise9section3)
        - [Exercise10Section3](#exercise10section3)
      - [ControlFlowandErrorHandling](#controlflowanderrorhandling)
        - [Passme](#passme)
        - [PassMulti](#passmulti)
        - [PartOfDays](#partofdays)
        - [Exercise1Section4](#exercise1section4)
        - [Exercise2Section4](#exercise2section4)
          - [Movies](#movies)
          - [OddOrEven](#oddoreven)
          - [LeapYear](#leapyear)
          - [DaysOfMonth](#daysofmonth)
        - [Exercise3Section4](#exercise3section4)
          - [richterScale1](#richterscale1)
          - [richterScale2](#richterscale2)
          - [convert](#convert)
          - [stringManipulation](#stringmanipulation)
          - [DaysOfMonth2](#daysofmonth2)
        - [Exercise4Section4](#exercise4section4)
          - [multiplicationTable](#multiplicationtable)
          - [mathTable](#mathtable)
        - [Exercise5Section4](#exercise5section4)
          - [sum](#sum)
          - [sumVerbos](#sumverbos)
          - [sumN](#sumn)
          - [evenOnly](#evenonly)
          - [breakUp](#breakup)
          - [infiniteKill](#infinitekill)
      - [ProjectforBeginners](#projectforbeginners)
        - [luckyNum](#luckynum)
        - [Exercise1Section5](#exercise1section5)
          - [luckyNumVerbose](#luckynumverbose)
          - [luckyNumIncrGuess](#luckynumincrguess)
          - [luckyNumDynDif](#luckynumdyndif)
        - [Exercise2Section5](#exercise2section5)
          - [crunchPrimes](#crunchprimes)
          - [pathSearch](#pathsearch)
          - [wordSearch](#wordsearch)
      - [ArraysandMemory](#arraysandmemory)
        - [moodly](#moodly)
        - [Exercise1Section6](#exercise1section6)
          - [assingArray](#assingarray)
          - [compareArrays](#comparearrays)
          - [emptyArray](#emptyarray)
          - [fix](#fix)
          - [getSetArray](#getsetarray)
        - [Exercise2Setion6](#exercise2setion6)
          - [averages](#averages)
          - [bookSearch](#booksearch)
          - [currencyConverter](#currencyconverter)
          - [numberSort](#numbersort)
          - [wizardPrinter](#wizardprinter)
          - [wordFinder](#wordfinder)

### learngo

this folder is the course work provided by the instructor. I have cloned his repo from <https://github.com/inangumus/learngo/> in order to keep course materials close at hand.

### MyWork

This folder is the location for my work and notes. I will keep the work I do and code I type in here to separate it from course materials. The _Notes.txt._ file is where I will write up my notes on the subject. It will always remain in the MyWork folder. Since this folder holds my work done, each section will be broken down into it's respective folder. Inside each section, there will be folders for work done while following along and folders for exercises. __Exercise__ folders will be labeled accoridng to exercise number in order. IE: if the folder is labeled _Exercise1_ then the folder is the first exercise in the section. It will be commented accrodingly for documentation.

#### WriteYourFirstGoProgram

Code pertaining to the course section 'Write Your First Go Program'

##### Exercise1Section2

Code pertaining to the first exercise of the section.

##### Exercise2Section2

Code pertaining to the second exercise of the section.

##### first

This is the first program written in the course and is a simple 'Hello World' style file.

##### printer

First library package written in the course. It holds one function to show the difference between an exported object and an unexported object. The cmd folder has one main.go file to demonstrate the errors that are thrown when trying to call exported and unexported object.

##### scopes

Simple go file to demonstrate scope of an object inside of go program.

#### MastertheTypeSystemofGo

Code pertaining to the course section Master the Type System of Go

##### Greeter

A samll program that takes in a command line argument and greets the user using it. Used to demonstrate user input, the os package, and the os.Args slice.

##### CelsiusToFahrenheit

This is a simple program written during the section to give practice working with arithmetic operators, precedence order, string conversion, and formatted printing. The program takes in a command line argument and converts it to a fahrenheit temperature with the assumption that the input is a temperature in celsius. It then rints the result to the screen with proper formatting.

##### Banger

This a is a simple program written during the seciton to give practice working with strigns and manipulating them. The program takes a singl command line argument, counts the number of characters in the resulting sting, and adds that number of exclamation points to the beginning and end of the string.

##### Exercise1Section3

Code pertaining to the first exercise in section 3. It shows examples of printing types.

##### Exercise2Section3

Code pertaining to the second exercise in section 3. SHows the various declaration method and the zero values for the specific types.

##### Exercise3Section3

Code pertaining to the third exercise of section 3. Shows the use of short variable declaration and multi-declaration and assignment.

##### Exercise4Section3

Code pertaining to the fourth exercise of section 3. Shows the use of variable assignment and variable swapping.

##### Exercise5Section3

Code pertaining to the fifth exercise of section 3. Shows the use of type conversion and correction.

##### Exercise6Section3

Code pertaining to the sixth exercise of section 3. Shows the use of command line arguments in a program and basic information of slices.

##### Exercise7Section3

Code pertaining to the seventh exercise of section 3. Shows use of the Printf function and formatting values. Also demonstrates using the backslash (\\) to use escape sequences or escape symbols suchs as double quotes.

##### Exercise8Section3

Code pertaining to the eighth exercise of section 3. Shows use of basic arithmetic and precedence order. Also shows how to change the order using parentheses.

##### Exercise9Section3

Code pertaining to the ninth exercise of section 3. Shows type safety, custom types, and the basic type sizes in go. This axercise shows how we can manipulate the types size to adjust accodingly.

##### Exercise10Section3

Code pertaining to the tenth exercise of section 3. Shows the use of constants. How to declare them, make them typed or typeless, and use of iota.

#### ControlFlowandErrorHandling

Code pertaining to the course section Control Flow and Error Handling. At this point in the cours exercise problems are getting more robust and longer. They will now be with in the exercisefolder and each individual problem will be in its own sub folder.

##### Passme

A simple code project in go to show if branching and logical operators. It will take command line input and "validate" agianst some conditions and print out various messages. It is a very simplistic and unsecure password checking tool. This code will only chekc for one user and its correct password.

##### PassMulti

An a more robust version of the Passme project. While this project ufnction in identical fashion, it has the ability to validate two different users, rather than a single user.

##### PartOfDays

A simple execution of a switch statement to part the days into a specific greeting message. It takes the curent hour of the computer and prints the appropriate response.

##### Exercise1Section4

Code pertaining to first exercise of section 4. Shows examples of the if statement and else if branches.

##### Exercise2Section4

Code pertaining to the second exercise in section 4. Shows examples of more if statments and error handling.

###### Movies

First problem in exercsie 2. A small program that checks a person's age and prints what movie ratings they are permitted to see.

###### OddOrEven

Second Problem in exercise 2. The program takes in an argument and checks if it is eve or odd and if it is divisble by 8.

###### LeapYear

Third and fourth problem in exercise 2. The program will take a command line argument of a year and print out whether it is a leap year or not. The fourth exercise is to refine the third exercise.

###### DaysOfMonth

The fifth problem of the exercise. The program takes in a command line argument of a onth name and notifies the user how many days are in the month. It accomodates whether it is a leap year of the current day.

##### Exercise3Section4

The third exercise in seciton 4. Code pertinaing to the use of switch statements.

###### richterScale1

First problem of the exercise. Takes in a command line argument, an earthquake magnitude, and returns a message of magnitude using a switch statement.

###### richterScale2

Second problem of the exercise.  Takes in a command line message and prints the approapriate magnitude. It is essential the opposite of the first problem while still using switch statements.

###### convert

Third problem in the exercise. The problem is to convert the commented if statement into a switch statement.

###### stringManipulation

Fourth problem of the exercise. The program manipulates a string by a chosen command. Both the command and string are entered at the command line. My variaiton is more advanced as I chose to research how to loop through an array or slice and build the string if neceesary. The protects to a degree around whether the user has remembered or not the quotation marks at the command line.

###### DaysOfMonth2

Fifth and final problem of exercise 3. This program is a refactored version of the days in month problem from the last exercise. The refactored version utilizes a swtich statement rather than if statement.

##### Exercise4Section4

this is the fourth exercie in the section. It focuses on the use of for loops to repeat a body of code in go

###### multiplicationTable

The first problem in the exercise. To build a multiplication table after getting the size of the table from the user in the form of a comman line argument.

###### mathTable

The second problem of the exercise. This problem is similar to the multiplication table, except instead of just a size being entered at the command line, so to is the operation for the table to be built on.

##### Exercise5Section4

The fifth exercise in the section. It also governs around the use fo a for loop and different structures of it.

###### sum

uses a for loop to sum the numbers 1 to 10.

###### sumVerbos

takes sum and adds a verbos print out of the actual addition.

###### sumN

sum but the starting and stopping point are given at the command line as arguments passed by the user.

###### evenOnly

takes sumN and only prints th sum of the ven numbers only still in verbose style and usign arguments given by the user at the command line.

###### breakUp

this is the sumN program but rather than using the for statment to stop the loop, we use the break keyword.

###### infiniteKill

a small program that continues to run printing the same thing over and over untilthe keyboard intercup of Ctrl-C.

#### ProjectforBeginners

Code pertaining to small projects designed to teach some small sections of GO like randomizaiont and labeled statements.

##### luckyNum

A small command line game that utilizes random values. This program also further shows the use of loops. It has been modified to meet the lucky number challenges presented at the end of the lecture 176 - Write the Game Logic.

##### Exercise1Section5

First exercise folder in the section.

###### luckyNumVerbose

Takes the Lucky NUmber game and adds a verbose flag option to allow the user to see the computer generated guess.

###### luckyNumIncrGuess

Alters the lucky number game to allow the user to plug in a higher than 10 gues, the game will then adjust the range of random numbers to meet this new max.

###### luckyNumDynDif

Alters the lucky number game for dyanmic difficulty based on the number fo the gues. the number of turns for the computer will be based on the number entered as the user's guess.

##### Exercise2Section5

Another simple set of projects to work with labeling and code reuse

###### crunchPrimes

Program that takes in a list of numbers and prints out only the prime numbers

###### pathSearch

Program that searches the path variable for keywors supplied by the user.

###### wordSearch

a simple word search porgram that is case insensitive, allowing the user to search a pre-set string.

#### ArraysandMemory

this section covers arrays, slices, and memory layout when using them.

##### moodly

a simplw program challenge to practice using arrays. It randomly selects a moood form an array of options.

##### Exercise1Section6

First exercise in section 6. More practice using arrays

###### assingArray

uses an exisitng array to populate another two and then manipulate the two new arrays.

###### compareArrays

a simple program that required me to fix the commented section of the code.

###### emptyArray

using array declaration , delcare a series of empty arrays of specific types.

###### fix

a code project that required the commented code to be repaired to be functional.

###### getSetArray

the program covers several problems in one, it uses other ways to declare and array and print the data from it.

##### Exercise2Setion6

the second exercise in the seciton working further with arrays.

###### averages

using arrays, takes in at the most 5 command line arguments, adds them to an array and produces the sum and "rough" average of the numbers entered.

###### bookSearch

a simplified search engine that searches a list of books for key words entered at the command line.

###### currencyConverter

a simplified currency converter using an array to hold the ratio values.

###### numberSort

a number sorting program that uses an array to hold 5 numbers entered at the command line by the user, and sorts them using the bubble sort algorithm.

###### wizardPrinter

an array loop to pretty print a multi-dimensional array of names.

###### wordFinder

a new modified word finder with an array of filtered words.
