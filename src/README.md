# Keakie Microservices
You can place any usefeul information about your submission here.

To run: 
- cd into either of the shows/genre folders
- go run main.go

To test manually:
- in public directory I've added a postman collection 
which can be used to test or make request from another method, e.g. curl

To improve:
- Find a better way to loop through genre requests
- Add config file (e.g. for other service url)
- Add way better error handling, e.g. status codes etc and better formatted responses
- Add docker files etc for deployment to aws on push to main
- Potentially add air.conf/Makefile to improve development with env
/ have an auto building server which doesnt need closing
- add unit tests

TRACKING TIMELINE 
- started fri @ 8.25am for 35 mins => 9.00
- back @9.20 for 70 mins => 10.30 
Followed gokit guide as never used one. 
Struggled getting json DB working quickly.
Genre working needs debugging, e.g. the way the json is handled.
- back @ 12.15 for 55 mins => 13:10
added found bool to verify slug
shows "working" @13:10 without test case for incorrect slug
1.5 hrs left to tidy up
- back @ 13:50 until 15:00 for 40 mins
Tidied up errors and basic logging. 
Wrote Readme run, test and improvements 
Began writing some unit tests but ran out of time really
time spent: 3 hrs 50 mins