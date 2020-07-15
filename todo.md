# TODO

## Input protocol
* [ ] document input protocol

## Test framework
* [ ] Perform the test of the complete app
* [ ] Improve build processing (cross compile, directory)
* [ ] Implement continuous integration



## Input processing
* [x] infer RST
* [x] Create the logic to take over from the previous line
* [x] decode and check frequency 
* [ ] New MYGRID keyword
* [ ] Support different date delimiter
* [x] Support extrapolated date
* [ ] Support date not prefixed by "date" (non header mode) DATE keyword is now optional
* [ ] Support date increment
* [ ] Support WWFF keyword
* [ ] Validate that we have the necessary data for the output 

## Output processing
* [x] WWFF ADIF output
* [x] Standard ADIF output
* [x] SOTA ADIF
* [x] SOTA CSV

## Later 
* [ ] Process contest reports
* [ ] Infer digital mode report
* [ ] add unit test for the DefaultReport validate function
* [ ] look how to make it work as a CGI

