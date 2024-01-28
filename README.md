## Small Go example

Small go example showing Goroutines and channel.<br>
The program will retrieve the content from some websites in parallel and display the number of characters
for each website.
I can be run without compiling by cloning the module and running:
```
go run *.go
```
The program will log what it's doing, for example:
```
2024/01/28 22:53:44 Handling https://nu.nl
2024/01/28 22:53:44 fetchWebsite(https://nu.nl)
2024/01/28 22:53:44 Handling http://www.telegraaf.nl
2024/01/28 22:53:44 fetchWebsite(http://www.telegraaf.nl)
2024/01/28 22:53:44 Handling https://nos.nl
2024/01/28 22:53:44 fetchWebsite(https://nos.nl)
2024/01/28 22:53:44 Return result for http://www.telegraaf.nl
2024/01/28 22:53:44 Return result for https://nos.nl
2024/01/28 22:53:44 Return result for https://nu.nl
2024/01/28 22:53:44 Website: https://nos.nl, Content: 718551 
2024/01/28 22:53:44 Website: https://nu.nl, Content: 1055731 
2024/01/28 22:53:44 Website: http://www.telegraaf.nl, Content: 12337 

```