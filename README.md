# Pomodoro Timer CLI

A simple command-line Pomodoro timer built in Go. This tool allows you to choose between short and long Pomodoro sessions, run the timer in the background, check the remaining time, and receive desktop notifications when the session is complete.

## Features

- **Session Types**: Choose between a short (default) or long Pomodoro session.
- **Background Timer**: The timer runs in the background, allowing you to interact with the CLI.
- **Status Check**: Check the remaining time of your session at any point.
- **Desktop Notifications**: Receive a notification when your Pomodoro session is complete.
- **Continuous Integration**: Automated testing and building using GitHub Actions.

## Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/liquidevelop/pomodoro.git
    cd pomodoro
    ```

2. **Install dependencies**:
    Ensure Go is installed on your system, then run:
    ```bash
    go get -u github.com/gen2brain/beeep
    ```

3. **Build the project**:
    ```bash
    go build -o pomodoro
    ```

## Usage

Run the `pomodoro` binary with the following options:

1. **Start a Pomodoro Session**:
    ```bash
    ./pomodoro
    ```
    - When prompted, enter `short` for a 5-minute session or `long` for a 25-minute session. If no input is provided, a short session will start by default.

2. **Check Timer Status**:
    While the timer is running, you can type `status` to check the remaining time, or `exit` to quit the program.

3. **Example Interaction**:
    ```
    Enter Pomodoro type (short/long, default is short): long
    Starting a 25:00 Pomodoro...
    Enter 'status' to check remaining time or 'exit' to quit: status
    Time remaining: 24:35
    Enter 'status' to check remaining time or 'exit' to quit: exit
    Exiting the Pomodoro timer.
    ```

## Development

### Running Tests

To run the tests for the timer and utility functions:

```bash
go test
```

## Github Actions
THis project includes a GitHub Actions workflow for continuous integration:
- <b>Testing:</b> Automatically runs unit tests on each push and pull request to `main` to ensue code reliability
- <b>Building:</b> Verifies that the project build successfully, ensuring that the code is ready for deployment.

## License
This project is licensed under the MIT License. See the <b>License</b> tab for details

## Acknowledgement
<!-- Create a MD link with https://github.com/gen2brain/beeep and text beeep -->

- [beeep](https://github.com/gen2brain/beeep) - Used for desktop notifications