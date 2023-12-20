# Groupie Trackers - API Data Visualization Project

Groupie Trackers is a web application project that aims to receive data from a given API, manipulate the data, and create a user-friendly website to display information about bands and artists, their concert locations, concert dates, and their relations. The project also emphasizes the creation of events/actions and their visualization.

## Project Overview

### API Structure

The API consists of four parts:

1. **Artists**: Contains information about bands and artists, including their name(s), image, the year they began their activity, the date of their first album, and the members.
2. **Locations**: Contains information about their last and/or upcoming concert locations.
3. **Dates**: Contains information about their last and/or upcoming concert dates.
4. **Relation**: Establishes the links between the other parts, connecting artists, dates, and locations.

### Project Objectives

- Develop a backend server using Go.
- Build a user-friendly website to visualize data from the API.
- Ensure the website and server operate without crashing and handle errors gracefully.
- Adhere to good coding practices.
- Include unit tests for code validation.

### Project Features

- Data manipulation and storage.
- Parsing and handling JSON data.
- HTML web page creation.
- Implementation of client-server communication.
- Event-driven actions.

## Getting Started

To get started with the Groupie Trackers project, follow these steps:

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/yourusername/groupie-trackers.git
   ```

2. Install the required Go packages:

   ```bash
   go get
   ```

3. Run the backend server:

   ```bash
   go run server.go
   ```

4. Access the web application by opening a web browser and navigating to:

   ```
   http://localhost:8040
   ```

## Project Structure

The project is organized into the following directories:

- **api**: Contains the Go backend code to handle API data and serve web requests.
- **html**: Contains the HTML templates for the web pages.
- **static**: Contains static assets (e.g., CSS, JavaScript, images).
- **tests**: Includes unit test files for code validation.

## Data Visualization

The website provides various ways to visualize data, including:

- **Blocks**: Displaying artists and their information in a block format.
- **Cards**: Presenting data in a card-style layout.
- **Tables**: Organizing data in tables.
- **Lists**: Showing artists and their information in a list format.
- **Graphics**: Visualizing data using charts and graphs.

## Events and Actions

The project includes event-driven actions, such as client-server communication, to retrieve and display information on the website.

## Maintainers

- [Hussain Alboori](https://github.com/hussainalboori)
- [Hussain Nabeel](https://github.com/7ussainnabeel)
- Ali Alsendi

## Contributing

If you'd like to contribute to the Groupie Trackers project, please follow these steps:

1. Fork the repository on GitHub.
2. Create a new branch for your feature or bug fix.
3. Implement your changes.
4. Write unit tests if necessary.
5. Commit your changes and push them to your fork.
6. Submit a pull request for review.

## Conclusion

Groupie Trackers is a Go-based web application project that focuses on manipulating API data and visualizing it in a user-friendly manner. It also demonstrates the client-server interaction with event-driven actions. We hope you find this project educational and insightful. Feel free to contribute and make it even better!

Happy coding!