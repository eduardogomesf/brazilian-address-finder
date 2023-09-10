# Project Overview
**Brazilian Address Finder**: REST API that helps you search for Brazilian addresses by the zip code

## Key Technologies:
[Go](https://go.dev/)  

## Project Purpose:
"Brazilian Address Finder" serves as a practical illustration of how to create a simple and clean REST API using the **Go** language

## System Design
![Alt text](./docs/images/system-design.jpeg)

## Running the Application 
1\. **Install Go**:
If you don't have Go installed, download it from Go official website and follow the installation instructions.

2\. **Clone the Repository**:
```
git clone https://github.com/eduardogomesf/brazilian-address-finder.git
cd brazilian-address-finder
```

3\. **Run the application**:
```
go run src/main.go
```

4\. **Searching for an address**:
To search for a brazilian address by zip code, perform a GET Request like the example below:
```
http://localhost:8080/addresses/88706007 # You can copy it and paste it on the browser to see the result
```

x\. **Doing more tests**
If you want to test additional zip codes and aren't sure where to find them, use the link below to generate more:
```
https://www.geradordecep.com.br/ 
```

---
🛠️ Created by [eduardogomesf](https://github.com/eduardogomesf)
