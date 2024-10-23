Time allocation: 4 hours

Instructions:
You have been asked to take a list of contact information and identify which contacts are potentially duplicates.  You need to write a function that will do the following:

	•	Identify which contacts are possible matches. 
	•	A score for each match that represents how accurate you believe the match is.  This scoring is defined by you.
	•	A contact might have multiple or no matches
	•	All processing should be done in working memory (no database).

Example Input:

Contact ID
First Name
Last Name
Email Address
Zip Code
Address
1001
C
F
mollis.lectus.pede@outlook.net

449-6990 Tellus. Rd.
1002
C
French
mollis.lectus.pede@outlook.net
39746
449-6990 Tellus. Rd.
1003
Ciara
F
non.lacinia.at@zoho.ca
39746


Example Output:
ContactID Source
ContactID Match
Accuracy
1001
1002
High
1001
1003
Low

The example input file is included in the assessment packet.

The input file is a CSV file but you can modify the input file to any format you choose; the file is intended to test your solution.  

What to focus on:
	•	The code used to identify the potential duplicates is the focus; code to take inputs from a file is not the focus and you can use libraries, etc to use the file.  Your code to identify possible duplicates should cover the all of the examples listed in the input file
	•	A set of tests that you use to verify the functions
	•	Comments in the code are welcome



