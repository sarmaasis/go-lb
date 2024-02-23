# go-lb
Demonstration of Loadbalacer using Round robin algorithm.



### Built With

* Go

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

* Go

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/sarmaasis/go-lb
   ```
2. Change Directory
   ```sh
   cd go-lb
   ```
3.
    ```go
    ** Navigate to main.go
    ** Change the <your-end -point> to the actual endpoint 
    http.Handle("/<your-end-point>", roundRobin) 
    ```
4. Build the project
   ```sh
   go build 
   ```
5. Run
   ```sh
   ./go-lb -backend="<server-ip:port>, <server-ip:port>, <server-ip:port>"
   ```


<!-- CONTACT -->
## Contact

Ashish Sharma - [@sarmaasis](https://linkedin.com/in/sarmaasis) - sarmaasis@gmail.com

<p align="right">(<a href="#readme-top">back to top</a>)</p>




