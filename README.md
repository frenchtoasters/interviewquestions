# Pair programming session

* Explain the code, and what the output is
	* Output:
		```
		> go run main.go
        I0815 15:09:33.054845   10639 main.go:40] Start: {bob }
        I0815 15:09:33.054997   10639 main.go:47] Data: {"name":"bob","Full":""}
		```
	* After fixing bug:
        ```
        > go run main.go
        I0815 15:10:49.501881   11706 main.go:40] Start: {bob ross}
        I0815 15:10:49.502081   11706 main.go:47] Data: {"name":"bob","last":"ross"}
		```
	* Why was `Full` part of the bad output?
* Add a new field, and have it as part of the output

## Questions

* What is the difference between strongly typed and loosely typed lanuages?
	* Loosely: Dont have to explicitly specify types for variables, Python, JS...
	* Strongly: Have to explicitly specify types for variables, Go, C, Java...
* What is an example of a project where you have used OOP?
	* Can you explain how some of the objects were related?
* I have the following files in a local directory:
```
.
├── --connect-timeout
├── -v
├── main.tf
├── terraform.tfstate.backup
├── terraform.tfstate.gpg
└── versions.tf

0 directories, 6 files
```
	* How would you remove the files starting with `-` and `--`?
