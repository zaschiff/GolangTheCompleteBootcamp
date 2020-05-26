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
        - [Exercise1Section4](#exercise1section4)
        - [Exercise2Section4](#exercise2section4)
          - [Movies](#movies)
          - [OddOrEven](#oddoreven)

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

##### Exercise1Section4

Code pertaining to first exercise of section 4. Shows examples of the if statement and else if branches.

##### Exercise2Section4

Code pertaining to the second exercise in section 4. Shows examples of more if statments and error handling.

###### Movies

First problem in exercsie 2. A small program that checks a person's age and prints what movie ratings they are permitted to see.

###### OddOrEven

Second Problem in exercise 2. The program takes in an argument and checks if it is eve or odd and if it is divisble by 8.
