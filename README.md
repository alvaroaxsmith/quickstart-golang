# quickstart-golang

This is a simple website monitoring project in Go. It allows you to monitor websites and log their status.

## Project Structure

The project has the following directory structure and files:

- `src/`: Folder containing the project source code.
- `sites.txt`: Text file that contains the list of websites to monitor.
- `log.txt`: Text file where the monitoring logs are recorded.
- `.replit`: Replit configuration file.
- `.replit.nix`: Additional configuration file for Replit.
- `go.mod`: Go package configuration file for dependency management.

## Running the Project

Make sure you have the Go development environment installed on your machine.

1. Clone this repository to your local machine.
2. Navigate to the project's root directory.
3. Run the command `go run src/main.go` to execute the program.

## Features

The program offers the following features:

1. **Start Monitoring**: Initiates the monitoring of the websites listed in the `sites.txt` file and displays their status.
2. **Display Logs**: Shows the logs recorded in the `log.txt` file.
3. **Exit Program**: Terminates the program execution.

## Customization

You can customize the program according to your needs. Some points to consider:

- **User Name**: The name displayed in the introduction can be changed by modifying the `name` variable in the `displayIntroduction()` function in the `main.go` file.
- **Version**: The program version can be changed by modifying the `version` variable in the `displayIntroduction()` function in the `main.go` file.
- **Monitoring Interval**: The time interval between website checks can be modified by changing the `delay` constant in the `main.go` file.
- **Number of Monitorings**: The number of times the monitoring will be performed can be changed by modifying the `monitorings` constant in the `main.go` file.

## Contribution

This project was developed based on the Go course by Alura. I would like to express my gratitude to Alura for their excellent course content, which can be accessed [here](https://www.alura.com.br/curso-online-golang).

Feel free to contribute by submitting pull requests with improvements, bug fixes, or new features.

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use and modify it according to your requirements.
