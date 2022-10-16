<!-- GETTING STARTED -->
## Getting Started

To run the project, you will need GoLang installed on your machine.

Go can be installed here: <p><a href="https://go.dev/doc/install">Install go link</a></p>


### Installation

_Below is an example of how you can instruct your audience on installing and setting up your app. This template doesn't rely on any external dependencies or services._

1. Install go language (follow steps above)

2. Clone the repo
   ```sh
   git clone https://github.com/thabo-moyo/talking_clock.git
   ```
3. Install dependecies
   ```sh
   go build && go mod vendor
   ```
4. Compile go
   ```sh
   go build main.go
   ```

<!-- USAGE EXAMPLES -->
## Usage

#### Objective 1: 
   ```sh
   ./main 16:00
   ```
   
Example output: Three o'clock

#### Objective 2: 

   ```sh
   ./main 16:00
   ```
Example output: Four o'clock

#### Objective 3: 
   A REST server will automatically start on your localhost.
   
   Accepted routes:
   
   Returns current time in human friendly format
   ```sh
   GET: /api/time/human-friendly
   ```
   Example output:
   
   
   Returns given time in human friendly format. Replace ```{slug}``` with time. E.g., ```15:00```

   ```sh
   GET: /api/time/human-friendly/{slug}
   ```