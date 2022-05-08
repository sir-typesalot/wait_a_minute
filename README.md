# Wait A Minute! (WAM!)
## Description

Wait a minute! Let us help you make a more infromed decision. Whether it is a finiancial or a lifestyle decision, take a step back and get a holistic view of it before you jump in. Consider many suggestions as well as weigh your options before you jump into something you might regret! Get tips from the countless others that have tried the same thing so that you can avoid some of the pitfalls.

## Local Development

This project is written in GO. You will need to have at GO v1.16 at least to be able to run this locally. To check your GO version, run 
```bash
go version
```
To install all the required packages, navigate into the directory of this project and run 
```bash
go get -u -v ./
```
This will install and update the dependencies with verbose output, to see if there were any errors throughout. 
With that set up, you're ready to run it locally. To run this you will have to use `go run *.go` to run all the file. To build the project, run `go build` to get an executable that can then be ran with `./<name_of_executable>`. 
Alternatively, you can also use `go install` to create the executable, but that will move the executable to the `$GOPATH/bin` directory.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
