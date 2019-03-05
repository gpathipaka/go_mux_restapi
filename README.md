GO Lang REST API Demo

Simple  API to perform REST operaitons on Employee Resource

#Install Gorilla/MUX router
go get -u github.com/gorilla/mux

bo bild
./go_mux_restapi

#Create a new Employee. HTTP POST
POST api/employee

#Get All Employees. HTTP GET
GET api/employees

#Get Employee by emp id. HTTP GET
GET api/employee{empId}

#Update the employee. HTTP PUT
PUT api/employee{empId}


#Delete the employee. HTTP DELETE
DELETE api/employee